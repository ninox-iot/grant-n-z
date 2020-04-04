package driver

import (
	"testing"

	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var permissionRepository PermissionRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	stubConnection, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = stubConnection
	permissionRepository = GetPermissionRepositoryInstance()
}

// FindAll InternalServerError test
func TestPermissionFindAll_InternalServerError(t *testing.T) {
	_, err := permissionRepository.FindAll()
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPermissionFindAll_InternalServerError test")
		t.FailNow()
	}
}

// FindOffSetAndLimit InternalServerError test
func TestPermissionFindOffSetAndLimit_InternalServerError(t *testing.T) {
	_, err := permissionRepository.FindOffSetAndLimit(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPermissionFindOffSetAndLimit_InternalServerError test")
		t.FailNow()
	}
}

// FindById InternalServerError test
func TestPermissionFindById_InternalServerError(t *testing.T) {
	_, err := permissionRepository.FindById(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPermissionFindById_InternalServerError test")
		t.FailNow()
	}
}

// FindByName InternalServerError test
func TestPermissionFindByName_InternalServerError(t *testing.T) {
	_, err := permissionRepository.FindByName("test")
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPermissionFindByName_InternalServerError test")
		t.FailNow()
	}
}

// FindByNames InternalServerError test
func TestPermissionFindByNames_InternalServerError(t *testing.T) {
	_, err := permissionRepository.FindByNames([]string{"test"})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPermissionFindByNames_InternalServerError test")
		t.FailNow()
	}
}

// FindByGroupId InternalServerError test
func TestPermissionFindByGroupId_InternalServerError(t *testing.T) {
	_, err := permissionRepository.FindByGroupId(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPermissionFindByGroupId_InternalServerError test")
		t.FailNow()
	}
}

// FindNameById name is nil test
func TestPermissionFindNameById_Nil(t *testing.T) {
	name := permissionRepository.FindNameById(1)
	if name != nil {
		t.Errorf("Incorrect TestPermissionFindNameById_Nil test")
		t.FailNow()
	}
}

// Save InternalServerError test
func TestPermissionSave_InternalServerError(t *testing.T) {
	_, err := permissionRepository.Save(entity.Permission{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPermissionSave_InternalServerError test")
		t.FailNow()
	}
}

// SaveWithRelationalData InternalServerError test
func TestPermissionSaveWithRelationalData_InternalServerError(t *testing.T) {
	_, err := permissionRepository.SaveWithRelationalData(1, entity.Permission{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestPermissionSaveWithRelationalData_InternalServerError test")
		t.FailNow()
	}
}
