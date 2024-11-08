package controllers

import (
	"company-hierarchy/models"
	"company-hierarchy/services"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller struct to manage requests
type DepartmentController struct {
    Service *services.DepartmentService
}

func (c *DepartmentController) AddDepartment(ctx *gin.Context) {
    var dept models.Department
    if err := ctx.ShouldBindJSON(&dept); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := c.Service.AddDepartment(dept); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add department"})
        return
    }
    ctx.JSON(http.StatusCreated, dept)
}

func (c *DepartmentController) UpdateDepartment(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var input struct {
        Name     *string `json:"name"`
        ParentID *int    `json:"parent_id"`
        Flags    *int    `json:"flags"`
    }
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Fetch the current department details
    dept, err := c.Service.GetDepartmentByID(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch department"})
        return
    }

    // Update only the provided fields
    if input.Name != nil {
        dept.Name = *input.Name
    }
    if input.ParentID != nil {
        dept.ParentID = input.ParentID
    }
    if input.Flags != nil {
        dept.Flags = *input.Flags
    }

    if err := c.Service.UpdateDepartment(dept); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update department"})
        return
    }
    ctx.JSON(http.StatusOK, dept)
}

func (c *DepartmentController) ActivateDepartment(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := c.Service.ActivateDepartment(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not activate department"})
        return
    }
    ctx.JSON(http.StatusNoContent, nil)
}

func (c *DepartmentController) DeactivateDepartment(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := c.Service.DeactivateDepartment(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not deactivate department"})
        return
    }
    ctx.JSON(http.StatusNoContent, nil)
}

func (c *DepartmentController) DeleteDepartment(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr) // Convert string to int
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := c.Service.DeleteDepartment(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete department"})
        return
    }
    ctx.JSON(http.StatusNoContent, nil)
}

func (c *DepartmentController) RestoreDepartment(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := c.Service.RestoreDepartment(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not restore department"})
        return
    }
    ctx.JSON(http.StatusNoContent, nil)
}

func (c *DepartmentController) ApproveDepartment(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := c.Service.ApproveDepartment(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not approve department"})
        return
    }
    ctx.JSON(http.StatusNoContent, nil)
}

func (c *DepartmentController) UnapproveDepartment(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := c.Service.UnapproveDepartment(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not unapprove department"})
        return
    }
    ctx.JSON(http.StatusNoContent, nil)
}

func (c *DepartmentController) GetHierarchy(ctx *gin.Context) {
    parentIDStr := ctx.Param("parent_id")
    parentID, err := strconv.Atoi(parentIDStr) // Convert string to int
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parent ID"})
        return
    }
    departments, err := c.Service.GetHierarchy(parentID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve hierarchy"})
        return
    }
    ctx.JSON(http.StatusOK, departments)
}

func SetupRoutes(router *gin.RouterGroup, db *sql.DB) {
    departmentService := services.NewDepartmentService(db)
    departmentController := &DepartmentController{Service: departmentService}

    departmentGroup := router.Group("/departments")
    {
        departmentGroup.POST("/", departmentController.AddDepartment)
        departmentGroup.PUT("/:id", departmentController.UpdateDepartment)
        departmentGroup.PUT("/:id/activate", departmentController.ActivateDepartment)
        departmentGroup.PUT("/:id/deactivate", departmentController.DeactivateDepartment)
        departmentGroup.DELETE("/:id", departmentController.DeleteDepartment)
        departmentGroup.PUT("/:id/restore", departmentController.RestoreDepartment)
        departmentGroup.PUT("/:id/approve", departmentController.ApproveDepartment)
        departmentGroup.PUT("/:id/unapprove", departmentController.UnapproveDepartment)
        departmentGroup.GET("/hierarchy/:parent_id", departmentController.GetHierarchy)
    }
}
