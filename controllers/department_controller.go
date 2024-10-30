package controllers

import (
	"company-hierarchy/models"
	"company-hierarchy/services"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Controller struct
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
    var dept models.Department
    if err := ctx.ShouldBindJSON(&dept); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := c.Service.UpdateDepartment(dept); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update department"})
        return
    }
    ctx.JSON(http.StatusOK, dept)
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
        departmentGroup.DELETE("/:id", departmentController.DeleteDepartment)
        departmentGroup.GET("/hierarchy/:parent_id", departmentController.GetHierarchy)
    }
}
