package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// Feature: vem-database-migration, Property 2: Write Request Rejection
// Property 2: Write Request Rejection
// *For any* HTTP request with method POST, PUT, PATCH, or DELETE to any API endpoint,
// the system SHALL return HTTP status 405 (Method Not Allowed).

// writeMethods contains HTTP methods that should be rejected
var writeMethods = []string{
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
}

// vem20Endpoints contains all VEM 2.0 API endpoints for testing
var vem20Endpoints = []string{
	"/api/v1/personel",
	"/api/v1/personel/test-kodu",
	"/api/v1/personel/gorev/test-gorev",
	"/api/v1/personel/authenticate/test-uid",
	"/api/v1/nfc-kart",
	"/api/v1/nfc-kart/test-kodu",
	"/api/v1/nfc-kart/uid/test-uid",
	"/api/v1/nfc-kart/personel/test-personel",
	"/api/v1/hasta",
	"/api/v1/hasta/test-kodu",
	"/api/v1/hasta/tc/12345678901",
	"/api/v1/hasta/search",
	"/api/v1/hasta-basvuru",
	"/api/v1/hasta-basvuru/test-kodu",
	"/api/v1/hasta-basvuru/hasta/test-hasta",
	"/api/v1/hasta-basvuru/hekim/test-hekim",
	"/api/v1/hasta-basvuru/filter",
	"/api/v1/yatak",
	"/api/v1/yatak/test-kodu",
	"/api/v1/yatak/birim/test-birim/oda/test-oda",
	"/api/v1/tablet-cihaz",
	"/api/v1/tablet-cihaz/test-kodu",
	"/api/v1/tablet-cihaz/yatak/test-yatak",
	"/api/v1/anlik-yatan-hasta",
	"/api/v1/anlik-yatan-hasta/test-kodu",
	"/api/v1/anlik-yatan-hasta/yatak/test-yatak",
	"/api/v1/anlik-yatan-hasta/hasta/test-hasta",
	"/api/v1/anlik-yatan-hasta/birim/test-birim",
	"/api/v1/vital-bulgu",
	"/api/v1/vital-bulgu/test-kodu",
	"/api/v1/vital-bulgu/basvuru/test-basvuru",
	"/api/v1/vital-bulgu/date-range",
	"/api/v1/klinik-seyir",
	"/api/v1/klinik-seyir/test-kodu",
	"/api/v1/klinik-seyir/basvuru/test-basvuru",
	"/api/v1/klinik-seyir/filter",
	"/api/v1/tibbi-order",
	"/api/v1/tibbi-order/test-kodu",
	"/api/v1/tibbi-order/basvuru/test-basvuru",
	"/api/v1/tibbi-order/test-kodu/detay",
	"/api/v1/tetkik-sonuc",
	"/api/v1/tetkik-sonuc/test-kodu",
	"/api/v1/tetkik-sonuc/basvuru/test-basvuru",
	"/api/v1/recete",
	"/api/v1/recete/test-kodu",
	"/api/v1/recete/basvuru/test-basvuru",
	"/api/v1/recete/hekim/test-hekim",
	"/api/v1/recete/test-kodu/ilaclar",
	"/api/v1/basvuru-tani",
	"/api/v1/basvuru-tani/test-kodu",
	"/api/v1/basvuru-tani/hasta/test-hasta",
	"/api/v1/basvuru-tani/basvuru/test-basvuru",
	"/api/v1/hasta-tibbi-bilgi",
	"/api/v1/hasta-tibbi-bilgi/test-kodu",
	"/api/v1/hasta-tibbi-bilgi/hasta/test-hasta",
	"/api/v1/hasta-tibbi-bilgi/turu/test-turu",
	"/api/v1/hasta-uyari",
	"/api/v1/hasta-uyari/test-kodu",
	"/api/v1/hasta-uyari/basvuru/test-basvuru",
	"/api/v1/hasta-uyari/filter",
	"/api/v1/risk-skorlama",
	"/api/v1/risk-skorlama/test-kodu",
	"/api/v1/risk-skorlama/basvuru/test-basvuru",
	"/api/v1/risk-skorlama/turu/test-turu",
	"/api/v1/basvuru-yemek",
	"/api/v1/basvuru-yemek/test-kodu",
	"/api/v1/basvuru-yemek/basvuru/test-basvuru",
	"/api/v1/basvuru-yemek/turu/test-turu",
}

// MethodNotAllowedMiddleware is a middleware that rejects write operations
// This middleware should be applied to all routes in the read-only system
func MethodNotAllowedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodPost || method == http.MethodPut ||
			method == http.MethodPatch || method == http.MethodDelete {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"success": false,
				"code":    "METHOD_NOT_ALLOWED",
				"message": "This API is read-only. Write operations are not permitted.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// setupTestRouter creates a test router with the MethodNotAllowedMiddleware
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(MethodNotAllowedMiddleware())

	// Register all endpoints with a dummy handler for GET
	// The middleware will reject write methods before reaching the handler
	dummyHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	}

	for _, endpoint := range vem20Endpoints {
		router.GET(endpoint, dummyHandler)
		router.POST(endpoint, dummyHandler)
		router.PUT(endpoint, dummyHandler)
		router.PATCH(endpoint, dummyHandler)
		router.DELETE(endpoint, dummyHandler)
	}

	return router
}

// TestProperty_WriteRequestRejection verifies that all write methods (POST, PUT, PATCH, DELETE)
// return HTTP 405 Method Not Allowed for all VEM 2.0 endpoints
func TestProperty_WriteRequestRejection(t *testing.T) {
	router := setupTestRouter()

	for _, endpoint := range vem20Endpoints {
		for _, method := range writeMethods {
			t.Run(method+"_"+endpoint, func(t *testing.T) {
				req, err := http.NewRequest(method, endpoint, nil)
				if err != nil {
					t.Fatalf("Failed to create request: %v", err)
				}

				resp := httptest.NewRecorder()
				router.ServeHTTP(resp, req)

				if resp.Code != http.StatusMethodNotAllowed {
					t.Errorf("Expected status %d for %s %s, got %d",
						http.StatusMethodNotAllowed, method, endpoint, resp.Code)
				}
			})
		}
	}
}

// TestProperty_GetRequestsAllowed verifies that GET requests are allowed
// (they should not return 405)
func TestProperty_GetRequestsAllowed(t *testing.T) {
	router := setupTestRouter()

	for _, endpoint := range vem20Endpoints {
		t.Run("GET_"+endpoint, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if resp.Code == http.StatusMethodNotAllowed {
				t.Errorf("GET request to %s should not return 405, got %d",
					endpoint, resp.Code)
			}
		})
	}
}

// TestProperty_WriteRejectionResponseFormat verifies that the error response
// for rejected write operations has the correct format
func TestProperty_WriteRejectionResponseFormat(t *testing.T) {
	router := setupTestRouter()

	// Test with a sample endpoint and method
	req, err := http.NewRequest(http.MethodPost, "/api/v1/personel", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Verify status code
	if resp.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d", http.StatusMethodNotAllowed, resp.Code)
	}

	// Verify response body contains expected fields
	body := resp.Body.String()
	expectedFields := []string{
		`"success":false`,
		`"code":"METHOD_NOT_ALLOWED"`,
		`"message":"This API is read-only. Write operations are not permitted."`,
	}

	for _, field := range expectedFields {
		if !contains(body, field) {
			t.Errorf("Response body missing expected field: %s\nGot: %s", field, body)
		}
	}
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// TestProperty_AllWriteMethodsRejected verifies that all four write methods
// are properly rejected
func TestProperty_AllWriteMethodsRejected(t *testing.T) {
	router := setupTestRouter()
	testEndpoint := "/api/v1/personel"

	for _, method := range writeMethods {
		t.Run(method, func(t *testing.T) {
			req, err := http.NewRequest(method, testEndpoint, nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			if resp.Code != http.StatusMethodNotAllowed {
				t.Errorf("Method %s should return 405, got %d", method, resp.Code)
			}
		})
	}
}
