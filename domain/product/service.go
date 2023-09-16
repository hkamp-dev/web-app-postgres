package product

import "context"

type Service interface {
	CreateProduct(ctx context.Context, p Product) (Product, error)
	GetProduct(ctx context.Context, id string) (Product, error)
	GetProducts(ctx context.Context, limit, offset int) ([]Product, error)
	DeleteProduct(ctx context.Context, id string) (Product, error)
	UpdateProduct(ctx context.Context, id string, p Product) (Product, error)
}

type service struct {
	products Repository
}
