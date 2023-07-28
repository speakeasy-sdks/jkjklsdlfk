// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/pkg/models/shared"
)

type PlaceOrderJSONResponse struct {
	ContentType string
	// successful operation
	Order       *shared.Order
	StatusCode  int
	RawResponse *http.Response
}

func (o *PlaceOrderJSONResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PlaceOrderJSONResponse) GetOrder() *shared.Order {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *PlaceOrderJSONResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PlaceOrderJSONResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
