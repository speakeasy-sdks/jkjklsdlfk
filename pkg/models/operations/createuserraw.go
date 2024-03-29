// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/v3/pkg/models/shared"
)

type CreateUserRawResponse struct {
	Body []byte
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	User *shared.User
}

func (o *CreateUserRawResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *CreateUserRawResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUserRawResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUserRawResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUserRawResponse) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}
