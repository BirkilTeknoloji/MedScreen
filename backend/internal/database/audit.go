package database

import (
	"context"
	"encoding/json"
	"medscreen/internal/models"
	"reflect"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Context keys for audit log information
type contextKey string

const (
	AuditUserIDKey    contextKey = "audit_user_id"
	AuditIPAddressKey contextKey = "audit_ip_address"
	AuditUserAgentKey contextKey = "audit_user_agent"
)

// WithAuditContext adds audit information to the context
func WithAuditContext(ctx context.Context, userID *uint, ip, userAgent string) context.Context {
	ctx = context.WithValue(ctx, AuditUserIDKey, userID)
	ctx = context.WithValue(ctx, AuditIPAddressKey, ip)
	ctx = context.WithValue(ctx, AuditUserAgentKey, userAgent)
	return ctx
}

// RegisterAuditCallbacks registers GORM callbacks for audit logging
func RegisterAuditCallbacks(db *gorm.DB) {
	// Create
	db.Callback().Create().After("gorm:create").Register("audit:after_create", afterCreateCallback)

	// Update
	db.Callback().Update().Before("gorm:update").Register("audit:before_update", beforeUpdateCallback)

	// Delete
	db.Callback().Delete().Before("gorm:delete").Register("audit:before_delete", beforeDeleteCallback)
}

func afterCreateCallback(db *gorm.DB) {
	if db.Error != nil || db.Statement.Schema == nil {
		return
	}

	// Skip audit log table itself to prevent recursion
	if db.Statement.Schema.Table == "audit_logs" {
		return
	}

	createAuditLog(db, "CREATE", nil, db.Statement.Dest)
}

func beforeUpdateCallback(db *gorm.DB) {
	if db.Error != nil || db.Statement.Schema == nil {
		return
	}

	if db.Statement.Schema.Table == "audit_logs" {
		return
	}

	// For updates, we want to capture old values.
	// However, getting old values requires a query which might be expensive.
	// For now, we will just log the new values (changes).
	// In a more advanced implementation, we could query the DB for old values here.

	// db.Statement.Dest contains the model with updated values
	createAuditLog(db, "UPDATE", nil, db.Statement.Dest)
}

func beforeDeleteCallback(db *gorm.DB) {
	if db.Error != nil || db.Statement.Schema == nil {
		return
	}

	if db.Statement.Schema.Table == "audit_logs" {
		return
	}

	// For delete, we might want to log what was deleted.
	// db.Statement.Dest usually contains the ID or condition.
	createAuditLog(db, "DELETE", db.Statement.Dest, nil)
}

func createAuditLog(db *gorm.DB, action string, oldVal, newVal interface{}) {
	ctx := db.Statement.Context

	// Extract audit info from context
	userID, _ := ctx.Value(AuditUserIDKey).(*uint)
	ip, _ := ctx.Value(AuditIPAddressKey).(string)
	userAgent, _ := ctx.Value(AuditUserAgentKey).(string)

	// Get Entity ID
	var entityID uint

	// Try to extract ID from newVal (for Create/Update)
	if newVal != nil {
		entityID = extractID(newVal, db.Statement.Schema)
	}

	// If ID is still 0 (e.g. Delete operation or failed extraction), try oldVal
	if entityID == 0 && oldVal != nil {
		entityID = extractID(oldVal, db.Statement.Schema)
	}

	// Serialize values
	var oldValJSON, newValJSON string
	if oldVal != nil {
		b, _ := json.Marshal(oldVal)
		oldValJSON = string(b)
	}
	if newVal != nil {
		b, _ := json.Marshal(newVal)
		newValJSON = string(b)
	}

	log := &models.AuditLog{
		UserID:     userID,
		Action:     action,
		EntityName: db.Statement.Schema.Table,
		EntityID:   entityID,
		OldValues:  oldValJSON,
		NewValues:  newValJSON,
		IPAddress:  ip,
		UserAgent:  userAgent,
		CreatedAt:  time.Now(),
	}

	// Use a new session without hooks to prevent recursion
	// We use a background context to ensure log is written even if original request context is cancelled
	db.Session(&gorm.Session{NewDB: true, SkipHooks: true, Context: context.Background()}).Create(log)
}

func extractID(value interface{}, schema *schema.Schema) uint {
	if value == nil {
		return 0
	}

	reflectValue := reflect.Indirect(reflect.ValueOf(value))

	// If it's a struct, try to find the primary key
	if reflectValue.Kind() == reflect.Struct {
		// Use GORM schema if available to find primary key
		if schema != nil {
			for _, field := range schema.PrimaryFields {
				// Assuming single primary key of type uint/int
				val, isZero := field.ValueOf(context.Background(), reflectValue)
				if !isZero {
					return convertToUint(val)
				}
			}
		}

		// Fallback: look for "ID" field
		idField := reflectValue.FieldByName("ID")
		if idField.IsValid() {
			return convertToUint(idField.Interface())
		}
	}

	return 0
}

func convertToUint(val interface{}) uint {
	switch v := val.(type) {
	case uint:
		return v
	case int:
		return uint(v)
	case uint64:
		return uint(v)
	case int64:
		return uint(v)
	case uint32:
		return uint(v)
	case int32:
		return uint(v)
	default:
		return 0
	}
}
