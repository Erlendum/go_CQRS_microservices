/*
 * Title
 *
 * Title
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"net/http"
)

// ProductsAPIService is a service that implements the logic for the ProductsAPIServicer
// This service should implement the business logic for every endpoint for the ProductsAPI API.
// Include any external packages or services that will be required by this service.
type ProductsAPIService struct {
}

// NewProductsAPIService creates a default api service
func NewProductsAPIService() ProductsAPIServicer {
	return &ProductsAPIService{}
}

// ProductsIdDelete - Delete product
func (s *ProductsAPIService) ProductsIdDelete(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update ProductsIdDelete with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProductsIdDelete method not implemented")
}

// ProductsIdGet - Get product
func (s *ProductsAPIService) ProductsIdGet(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update ProductsIdGet with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DtoProductResponse{}) or use other options such as http.Ok ...
	// return Response(200, DtoProductResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProductsIdGet method not implemented")
}

// ProductsIdPut - Update product
func (s *ProductsAPIService) ProductsIdPut(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update ProductsIdPut with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DtoUpdateProductDto{}) or use other options such as http.Ok ...
	// return Response(200, DtoUpdateProductDto{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProductsIdPut method not implemented")
}

// ProductsPost - Create product
func (s *ProductsAPIService) ProductsPost(ctx context.Context) (ImplResponse, error) {
	// TODO - update ProductsPost with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, DtoCreateProductResponseDto{}) or use other options such as http.Ok ...
	// return Response(201, DtoCreateProductResponseDto{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProductsPost method not implemented")
}

// ProductsSearchGet - Search product
func (s *ProductsAPIService) ProductsSearchGet(ctx context.Context, search string, page int32, size int32) (ImplResponse, error) {
	// TODO - update ProductsSearchGet with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DtoProductsListResponse{}) or use other options such as http.Ok ...
	// return Response(200, DtoProductsListResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ProductsSearchGet method not implemented")
}
