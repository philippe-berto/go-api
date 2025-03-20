package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/philippeberto/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 100)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	productDB := NewProduct(db)

	db.AutoMigrate(&entity.Product{})
	for i := range 25 {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i+1), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.NotEmpty(t, products)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 5", products[4].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.NotEmpty(t, products)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 15", products[4].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.NotEmpty(t, products)
	assert.Equal(t, 5, len(products))
	assert.Equal(t, "Product 25", products[4].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	productDB := NewProduct(db)

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 100)
	db.Create(product)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Product 1", productFound.Name)
	assert.Equal(t, 100.00, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	productDB := NewProduct(db)

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 100)
	db.Create(product)

	product.Name = "Product 2"
	product.Price = 200
	err = productDB.Update(product)
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, "Product 2", productFound.Name)
	assert.Equal(t, 200.00, productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	productDB := NewProduct(db)

	db.AutoMigrate(&entity.Product{})
	product, _ := entity.NewProduct("Product 1", 100)
	db.Create(product)

	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NotNil(t, err)
	assert.Nil(t, productFound)
}
