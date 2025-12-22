package repository

import (
	"reflect"
	"strings"
	"testing"
)

// Feature: vem-database-migration, Property 1: Read-Only Enforcement
// Property 1: Read-Only Enforcement
// *For any* repository interface in the system, the interface SHALL NOT contain
// Create, Update, or Delete methods—only Find/Get methods for reading data.

// forbiddenMethodPrefixes contains method name prefixes that indicate write operations
var forbiddenMethodPrefixes = []string{
	"Create",
	"Update",
	"Delete",
	"Save",
	"Remove",
	"Insert",
	"Upsert",
	"Add",
	"Set",
	"Modify",
	"Edit",
}

// allowedMethodPrefixes contains method name prefixes that indicate read operations
var allowedMethodPrefixes = []string{
	"Find",
	"Get",
	"Search",
	"List",
	"Count",
	"Exists",
}

// repositoryInterfaces contains all VEM 2.0 repository interface types for testing
var repositoryInterfaces = []reflect.Type{
	reflect.TypeOf((*PersonelRepository)(nil)).Elem(),
	reflect.TypeOf((*NFCKartRepository)(nil)).Elem(),
	reflect.TypeOf((*HastaRepository)(nil)).Elem(),
	reflect.TypeOf((*HastaBasvuruRepository)(nil)).Elem(),
	reflect.TypeOf((*YatakRepository)(nil)).Elem(),
	reflect.TypeOf((*TabletCihazRepository)(nil)).Elem(),
	reflect.TypeOf((*AnlikYatanHastaRepository)(nil)).Elem(),
	reflect.TypeOf((*HastaVitalFizikiBulguRepository)(nil)).Elem(),
	reflect.TypeOf((*KlinikSeyirRepository)(nil)).Elem(),
	reflect.TypeOf((*TibbiOrderRepository)(nil)).Elem(),
	reflect.TypeOf((*TetkikSonucRepository)(nil)).Elem(),
	reflect.TypeOf((*ReceteRepository)(nil)).Elem(),
	reflect.TypeOf((*BasvuruTaniRepository)(nil)).Elem(),
	reflect.TypeOf((*HastaTibbiBilgiRepository)(nil)).Elem(),
	reflect.TypeOf((*HastaUyariRepository)(nil)).Elem(),
	reflect.TypeOf((*RiskSkorlamaRepository)(nil)).Elem(),
	reflect.TypeOf((*BasvuruYemekRepository)(nil)).Elem(),
}

// TestProperty_ReadOnlyEnforcement verifies that all repository interfaces
// contain only read methods (Find/Get) and no write methods (Create/Update/Delete)
func TestProperty_ReadOnlyEnforcement(t *testing.T) {
	for _, repoType := range repositoryInterfaces {
		repoName := repoType.Name()
		t.Run(repoName, func(t *testing.T) {
			// Check each method in the interface
			for i := 0; i < repoType.NumMethod(); i++ {
				method := repoType.Method(i)
				methodName := method.Name

				// Check for forbidden method prefixes (write operations)
				for _, prefix := range forbiddenMethodPrefixes {
					if strings.HasPrefix(methodName, prefix) {
						t.Errorf("Repository %s contains forbidden write method: %s (prefix: %s)",
							repoName, methodName, prefix)
					}
				}

				// Verify method has an allowed prefix (read operations)
				hasAllowedPrefix := false
				for _, prefix := range allowedMethodPrefixes {
					if strings.HasPrefix(methodName, prefix) {
						hasAllowedPrefix = true
						break
					}
				}

				if !hasAllowedPrefix {
					t.Errorf("Repository %s contains method with unknown prefix: %s (expected one of: %v)",
						repoName, methodName, allowedMethodPrefixes)
				}
			}
		})
	}
}

// TestProperty_AllRepositoriesHaveMethods verifies that all repository interfaces
// have at least one method defined
func TestProperty_AllRepositoriesHaveMethods(t *testing.T) {
	for _, repoType := range repositoryInterfaces {
		repoName := repoType.Name()
		t.Run(repoName, func(t *testing.T) {
			if repoType.NumMethod() == 0 {
				t.Errorf("Repository %s has no methods defined", repoName)
			}
		})
	}
}

// TestProperty_FindByKoduExists verifies that all repository interfaces
// have a FindByKodu method for primary key retrieval
func TestProperty_FindByKoduExists(t *testing.T) {
	for _, repoType := range repositoryInterfaces {
		repoName := repoType.Name()
		t.Run(repoName, func(t *testing.T) {
			hasFindByKodu := false
			for i := 0; i < repoType.NumMethod(); i++ {
				method := repoType.Method(i)
				if method.Name == "FindByKodu" {
					hasFindByKodu = true
					break
				}
			}
			if !hasFindByKodu {
				t.Errorf("Repository %s does not have FindByKodu method for primary key retrieval", repoName)
			}
		})
	}
}

// TestProperty_NoContextInReadMethods verifies that read-only methods
// don't require context (since they don't perform write operations that need transactions)
// This is a design choice for the read-only system
func TestProperty_MethodSignaturesConsistent(t *testing.T) {
	for _, repoType := range repositoryInterfaces {
		repoName := repoType.Name()
		t.Run(repoName, func(t *testing.T) {
			for i := 0; i < repoType.NumMethod(); i++ {
				method := repoType.Method(i)
				methodType := method.Type

				// All methods should return at least one value (the result or error)
				if methodType.NumOut() == 0 {
					t.Errorf("Repository %s method %s has no return values",
						repoName, method.Name)
				}

				// Last return value should be error
				lastOut := methodType.Out(methodType.NumOut() - 1)
				if lastOut.String() != "error" {
					t.Errorf("Repository %s method %s last return value should be error, got %s",
						repoName, method.Name, lastOut.String())
				}
			}
		})
	}
}
