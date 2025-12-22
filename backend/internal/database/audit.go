package database

// This file previously contained audit logging callbacks for write operations.
// Since the VEM 2.0 migration converts this application to read-only mode,
// audit logging for Create, Update, and Delete operations is no longer needed.
//
// The VEM 2.0 database is managed externally, and this application only
// reads data from it. Any audit logging for the source data is handled
// by the VEM 2.0 system itself.
//
// This file is kept as a placeholder to document the architectural decision
// and to prevent import errors from any legacy code that might reference it.

// Note: RegisterAuditCallbacks has been removed as there are no write operations
// to audit in a read-only system.
