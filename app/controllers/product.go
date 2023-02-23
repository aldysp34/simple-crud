package controllers

import (
	"errors"
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
	c.BindJSON(&product)
	err := models.CreateProduct(repository.Db, &product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, product)
}

// get products
func (repository *ProductRepo) GetProducts(c *gin.Context) {
	var product []models.Product
	err := models.GetProducts(repository.Db, &product)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, product)
}

// get product by id
func (repository *ProductRepo) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi("id")
	var product models.Product
	err := models.GetProduct(repository.Db, &product, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, product)
}

// update product
func (repository *ProductRepo) UpdateProduct(c *gin.Context) {
	var product models.Product
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetProduct(repository.Db, &product, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&product)
	err = models.UpdateProduct(repository.Db, &product)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, product)
}

// delete Product
func (repository *ProductRepo) DeleteProduct(c *gin.Context) {
	var product models.Product
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteProduct(repository.Db, &product, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
