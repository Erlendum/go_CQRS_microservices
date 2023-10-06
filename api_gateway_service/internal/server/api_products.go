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
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ProductsAPIController binds http requests to an api service and writes the service results to the http response
type ProductsAPIController struct {
	service      ProductsAPIServicer
	errorHandler ErrorHandler
}

// ProductsAPIOption for how the controller is set up.
type ProductsAPIOption func(*ProductsAPIController)

// WithProductsAPIErrorHandler inject ErrorHandler into controller
func WithProductsAPIErrorHandler(h ErrorHandler) ProductsAPIOption {
	return func(c *ProductsAPIController) {
		c.errorHandler = h
	}
}

// NewProductsAPIController creates a default api controller
func NewProductsAPIController(s ProductsAPIServicer, opts ...ProductsAPIOption) Router {
	controller := &ProductsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ProductsAPIController
func (c *ProductsAPIController) Routes() Routes {
	return Routes{
		"ProductsIdDelete": Route{
			strings.ToUpper("Delete"),
			"/products/{id}",
			c.ProductsIdDelete,
		},
		"ProductsIdGet": Route{
			strings.ToUpper("Get"),
			"/products/{id}",
			c.ProductsIdGet,
		},
		"ProductsIdPut": Route{
			strings.ToUpper("Put"),
			"/products/{id}",
			c.ProductsIdPut,
		},
		"ProductsPost": Route{
			strings.ToUpper("Post"),
			"/products",
			c.ProductsPost,
		},
		"ProductsSearchGet": Route{
			strings.ToUpper("Get"),
			"/products/search",
			c.ProductsSearchGet,
		},
	}
}

// ProductsIdDelete - Delete product
func (c *ProductsAPIController) ProductsIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.ProductsIdDelete(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ProductsIdGet - Get product
func (c *ProductsAPIController) ProductsIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.ProductsIdGet(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ProductsIdPut - Update product
func (c *ProductsAPIController) ProductsIdPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	result, err := c.service.ProductsIdPut(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ProductsPost - Create product
func (c *ProductsAPIController) ProductsPost(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ProductsPost(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ProductsSearchGet - Search product
func (c *ProductsAPIController) ProductsSearchGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	searchParam := query.Get("search")
	pageParam, err := parseNumericParameter[int32](
		query.Get("page"),
		WithParse[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	sizeParam, err := parseNumericParameter[int32](
		query.Get("size"),
		WithParse[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.ProductsSearchGet(r.Context(), searchParam, pageParam, sizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
