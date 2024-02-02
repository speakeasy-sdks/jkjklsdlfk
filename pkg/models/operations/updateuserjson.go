// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/v3/pkg/models/shared"
)

type UpdateUserJSONRequest struct {
	// Update an existent user in the store
	User *shared.User `request:"mediaType=application/json"`
	// name that need to be deleted
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *UpdateUserJSONRequest) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *UpdateUserJSONRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type UpdateUserJSONResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateUserJSONResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUserJSONResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUserJSONResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
