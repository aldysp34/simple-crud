package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `form:"name" json:"name" validate:"required"`
	Description string  `form:"description" json:"description" validate:"required"`
	Price       float64 `form:"price" json:"price" validate:"required"`
	Quantity    int     `form:"quantity" json:"quantity" validate:"required"`
}

// Create Product
func CreateProduct(db *gorm.DB, Product *Product) (err error) {
	err = db.Create(Product).Error
	if err != nil {
		return err
	}

	return nil
}

// Get Products
func GetProducts(db *gorm.DB, Product *[]Product, limit int, offset int, filter string) (err error) {

	err = db.Where("name LIKE ?", "%"+filter+"%").Limit(limit).Offset(offset).Find(Product).Error
	if err != nil {
		return err
	}

	return nil
}

// Get Product by ID
func GetProduct(db *gorm.DB, Product *Product, id int) (err error) {
	err = db.Where("id = ?", id).First(Product).Error
	if err != nil {
		return err
	}

	return nil
}

// Update Product
func UpdateProduct(db *gorm.DB, Product *Product) (err error) {
	db.Save(&Product)
	return nil
}

// Delete Product
func DeleteProduct(db *gorm.DB, Product *Product, id int) (err error) {
	db.Where("id = ?", id).Delete(&Product)

	return nil
}
