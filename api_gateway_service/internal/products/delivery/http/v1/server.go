package v1

import (
	"context"
	"errors"
	"github.com/Erlendum/go_CQRS_microservices/api_gateway_service/internal/products/service"
	openapi "github.com/Erlendum/go_CQRS_microservices/api_gateway_service/internal/server"
	"net/http"
)

type Server struct {
	openapi.ProductsAPIService
	productService service.ProductService
}

func New(productService service.ProductService) openapi.ProductsAPIServicer {
	return &Server{
		productService: productService,
	}
}

func (s *Server) ProductsIdDelete(ctx context.Context, id string) (openapi.ImplResponse, error) {
	// TODO - update ProductsIdDelete with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ProductsIdDelete method not implemented")
}

func (s *Server) ProductsIdGet(ctx context.Context, id string) (openapi.ImplResponse, error) {
	// TODO - update ProductsIdGet with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DtoProductResponse{}) or use other options such as http.Ok ...
	// return Response(200, DtoProductResponse{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ProductsIdGet method not implemented")
}

func (s *Server) ProductsIdPut(ctx context.Context, id string) (openapi.ImplResponse, error) {
	// TODO - update ProductsIdPut with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DtoUpdateProductDto{}) or use other options such as http.Ok ...
	// return Response(200, DtoUpdateProductDto{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ProductsIdPut method not implemented")
}

func (s *Server) ProductsPost(ctx context.Context) (openapi.ImplResponse, error) {
	// TODO - update ProductsPost with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, DtoCreateProductResponseDto{}) or use other options such as http.Ok ...
	// return Response(201, DtoCreateProductResponseDto{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ProductsPost method not implemented")
}

func (s *Server) ProductsSearchGet(ctx context.Context, search string, page int32, size int32) (openapi.ImplResponse, error) {
	// TODO - update ProductsSearchGet with the required logic for this service method.
	// Add api_products_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, DtoProductsListResponse{}) or use other options such as http.Ok ...
	// return Response(200, DtoProductsListResponse{}), nil

	return openapi.Response(http.StatusNotImplemented, nil), errors.New("ProductsSearchGet method not implemented")
}
