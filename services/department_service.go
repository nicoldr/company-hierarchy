// Service to interact with database
package services

import (
	"company-hierarchy/models"
	"database/sql"
	"log"
)

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

func (s *DepartmentService) ActivateDepartment(id int) error {
    _, err := s.DB.Exec("CALL ActivateDepartment(?)", id)
    if err != nil {
        log.Printf("Error while marking department as active: %v", err)
        return err
    }
    return nil
}

func (s *DepartmentService) DeactivateDepartment(id int) error {
    _, err := s.DB.Exec("CALL DeactivateDepartment(?)", id)
    if err != nil {
        log.Printf("Error while marking department as deactivated: %v", err)
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

func (s *DepartmentService) RestoreDepartment(id int) error {
    _, err := s.DB.Exec("CALL RestoreDepartment(?)", id)
    if err != nil {
        log.Printf("Error while marking department as not deleted: %v", err)
        return err
    }
    return nil
}

func (s *DepartmentService) ApproveDepartment(id int) error {
    _, err := s.DB.Exec("CALL ApproveDepartment(?)", id)
    if err != nil {
        log.Printf("Error while marking department as approved: %v", err)
        return err
    }
    return nil
}

func (s *DepartmentService) UnapproveDepartment(id int) error {
    _, err := s.DB.Exec("CALL UnapproveDepartment(?)", id)
    if err != nil {
        log.Printf("Error while marking department as not approved: %v", err)
        return err
    }
    return nil
}

func (s *DepartmentService) GetHierarchy(parentID int) ([]models.Department, error) {
    log.Printf("Calling GetHierarchy with parentID: %d", parentID)
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
            log.Printf("Error scanning department: %v", err)
            return nil, err
        }
        departments = append(departments, dept)
    }

    if err := rows.Err(); err != nil {
        log.Printf("Error with rows: %v", err)
        return nil, err
    }

    log.Printf("Fetched departments: %v", departments)
    return departments, nil
}

func (s *DepartmentService) GetDepartmentByID(id int) (models.Department, error) {
    var dept models.Department
    row := s.DB.QueryRow("SELECT id, name, parent_id, flags FROM departments WHERE id = ?", id)
    err := row.Scan(&dept.ID, &dept.Name, &dept.ParentID, &dept.Flags)
    return dept, err
}