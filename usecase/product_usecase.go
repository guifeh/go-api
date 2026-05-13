package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repository,
	}
}

func (p *ProductUseCase) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}

func (p *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := p.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (p *ProductUseCase) GetProductByID(id int) (*model.Product, error) {
	product, err := p.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductUseCase) UpdateProduct(id int, product model.Product) (*model.Product, error) {
	updatedProduct, err := p.repository.UpdateProduct(id, product)
	if err != nil {
		return nil, err
	}
	return &updatedProduct, nil
}

func (p *ProductUseCase) DeleteProduct(id int) error {
	return p.repository.DeleteProduct(id)
}
