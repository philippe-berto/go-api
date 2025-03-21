package database

import (
	"github.com/philippeberto/go-api/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {

	return &Product{
		DB: db,
	}
}

func (p *Product) Create(product *entity.Product) error {

	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		order := "created_at " + sort
		err = p.DB.Offset(offset).Limit(limit).Order(order).Find(&products).Error
	} else {
		err = p.DB.Find(&products).Error
	}

	return products, err
}

func (p *Product) FindByID(id string) (*entity.Product, error) {
	product := &entity.Product{}
	err := p.DB.Where("id = ?", id).First(product).Error

	return product, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {

		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindByID(id)
	if err != nil {

		return err
	}

	return p.DB.Delete(product).Error
}
