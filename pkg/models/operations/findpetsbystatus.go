// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"petstore/pkg/models/shared"
	"petstore/pkg/utils"
)

// FindPetsByStatusStatus - Status values that need to be considered for filter
type FindPetsByStatusStatus string

const (
	FindPetsByStatusStatusAvailable FindPetsByStatusStatus = "available"
	FindPetsByStatusStatusPending   FindPetsByStatusStatus = "pending"
	FindPetsByStatusStatusSold      FindPetsByStatusStatus = "sold"
)

func (e FindPetsByStatusStatus) ToPointer() *FindPetsByStatusStatus {
	return &e
}

func (e *FindPetsByStatusStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "available":
		fallthrough
	case "pending":
		fallthrough
	case "sold":
		*e = FindPetsByStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FindPetsByStatusStatus: %v", v)
	}
}

type FindPetsByStatusRequest struct {
	// Status values that need to be considered for filter
	Status *FindPetsByStatusStatus `default:"available" queryParam:"style=form,explode=true,name=status"`
}

func (f FindPetsByStatusRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FindPetsByStatusRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FindPetsByStatusRequest) GetStatus() *FindPetsByStatusStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type FindPetsByStatusResponse struct {
	Body        []byte
	ContentType string
	// successful operation
	Pets        []shared.Pet
	StatusCode  int
	RawResponse *http.Response
}

func (o *FindPetsByStatusResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *FindPetsByStatusResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *FindPetsByStatusResponse) GetPets() []shared.Pet {
	if o == nil {
		return nil
	}
	return o.Pets
}

func (o *FindPetsByStatusResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *FindPetsByStatusResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
