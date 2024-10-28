package services

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDepartment(t *testing.T) {
    // Setup mock database connection
    db, err := sql.Open("mysql", "mock:mock@/test")
    if err != nil {
        t.Fatal(err)
    }
    service := NewDepartmentService(db)

    // Test adding a department
    dept := Department{Name: "HR"}
    err = service.AddDepartment(dept)
    assert.NoError(t, err)
}
