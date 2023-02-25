package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aldysp34/simple-crud/database"
	"github.com/aldysp34/simple-crud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductRepo struct {
	Db *gorm.DB
}

// Connecting to Model
func New() *ProductRepo {
	db := database.InitDB()
	db.AutoMigrate(&models.Product{})
	return &ProductRepo{Db: db}
}

// Create Products
func (repository *ProductRepo) CreateProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBind(&product)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	check := models.CreateProduct(repository.Db, &product)
	if check != nil {
		fmt.Println(check)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, product)
}

// get products
func (repository *ProductRepo) GetProducts(c *gin.Context) {
	var product []models.Product
	err := models.GetProducts(repository.Db, &product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// get product by id
func (repository *ProductRepo) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var product models.Product
	err := models.GetProduct(repository.Db, &product, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// update product
func (repository *ProductRepo) UpdateProduct(c *gin.Context) {
	var product models.Product
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := models.GetProduct(repository.Db, &product, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.ShouldBind(&product)
	err = models.UpdateProduct(repository.Db, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// delete Product
func (repository *ProductRepo) DeleteProduct(c *gin.Context) {
	var product models.Product
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	err := models.DeleteProduct(repository.Db, &product, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
