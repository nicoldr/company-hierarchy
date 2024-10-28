package services

import (
	"company-hierarchy/models"
	"database/sql"
	"log"
)

// Service sto interact with database
type DepartmentService struct {
    DB *sql.DB
}

func NewDepartmentService(db *sql.DB) *DepartmentService {
    return &DepartmentService{DB: db}
}

func (s *DepartmentService) AddDepartment(dept models.Department) error {
    _, err := s.DB.Exec("CALL AddDepartment(?, ?, ?)", dept.Name, dept.ParentID, dept.Flags)
    if err != nil {
        log.Printf("Error while adding department: %v", err)
        return err
    }
    return nil
}

func (s *DepartmentService) UpdateDepartment(dept models.Department) error {
    _, err := s.DB.Exec("CALL UpdateDepartment(?, ?, ?, ?)", dept.ID, dept.Name, dept.ParentID, dept.Flags)
    if err != nil {
        log.Printf("Error while updating department: %v", err)
        return err
    }
    return nil
}

func (s *DepartmentService) DeleteDepartment(id int) error {
    _, err := s.DB.Exec("CALL DeleteDepartment(?)", id)
    if err != nil {
        log.Printf("Error while deleting department: %v", err)
        return err
    }
    return nil
}

func (s *DepartmentService) GetHierarchy(parentID int) ([]models.Department, error) {
    rows, err := s.DB.Query("CALL GetHierarchy(?)", parentID)
    if err != nil {
        log.Printf("Error while getting hierarchy: %v", err)
        return nil, err
    }
    defer rows.Close()

    var departments []models.Department
    for rows.Next() {
        var dept models.Department
        if err := rows.Scan(&dept.ID, &dept.Name, &dept.ParentID, &dept.Flags); err != nil {
            log.Printf("Error while scanning row: %v", err)
            return nil, err
        }
        departments = append(departments, dept)
    }
    return departments, nil
}
