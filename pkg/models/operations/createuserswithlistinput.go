// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/pkg/models/shared"
)

type CreateUsersWithListInputResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful operation
	User *shared.User
}

func (o *CreateUsersWithListInputResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *CreateUsersWithListInputResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUsersWithListInputResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUsersWithListInputResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateUsersWithListInputResponse) GetUser() *shared.User {
	if o == nil {
		return nil
	}
	return o.User
}
