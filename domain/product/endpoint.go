package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateProduct endpoint.Endpoint
	GetProduct    endpoint.Endpoint
	GetProducts   endpoint.Endpoint
	DeleteProduct endpoint.Endpoint
	UpdateProduct endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateProduct: MakeCreateProductEndpoint(s),
		GetProduct:    MakeGetProductEndpoint(s),
		GetProducts:   MakeGetProductsEndpoint(s),
		DeleteProduct: MakeDeleteProductEndpoint(s),
		UpdateProduct: MakeUpdateProductEndpoint(s),
	}
}

type CreateProductRequest struct {
	Product Product `json:"product"`
}

type CreateProductResponse struct {
	Product Product `json:"product"`
	Err     error   `json:"err,omitempty"`
}

func MakeCreateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateProductRequest)
		p, err := s.CreateProduct(ctx, req.Product)
		return CreateProductResponse{p, err}, nil
	}
}

type GetProductRequest struct {
	ID string `json:"id"`
}

type GetProductResponse struct {
	Product Product `json:"product"`
	Err     error   `json:"err,omitempty"`
}

func MakeGetProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductRequest)
		p, err := s.GetProduct(ctx, req.ID)
		return GetProductResponse{p, err}, nil
	}
}

type GetProductsRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type GetProductsResponse struct {
	Products []Product `json:"products"`
	Err      error     `json:"err,omitempty"`
}

func MakeGetProductsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetProductsRequest)
		p, err := s.GetProducts(ctx, req.Limit, req.Offset)
		return GetProductsResponse{p, err}, nil
	}
}

type DeleteProductRequest struct {
	ID string `json:"id"`
}

type DeleteProductResponse struct {
	Product Product `json:"product"`
	Err     error   `json:"err,omitempty"`
}

func MakeDeleteProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteProductRequest)
		p, err := s.DeleteProduct(ctx, req.ID)
		return DeleteProductResponse{p, err}, nil
	}
}

type UpdateProductRequest struct {
	ID      string  `json:"id"`
	Product Product `json:"product"`
}

type UpdateProductResponse struct {
	Product Product `json:"product"`
	Err     error   `json:"err,omitempty"`
}

func MakeUpdateProductEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateProductRequest)
		p, err := s.UpdateProduct(ctx, req.ID, req.Product)
		return UpdateProductResponse{p, err}, nil
	}
}
