// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"petstore/pkg/models/shared"
)

type FindPetsByStatusSecurity struct {
	PetstoreAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

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
	Status *FindPetsByStatusStatus `queryParam:"style=form,explode=true,name=status"`
}

type FindPetsByStatusResponse struct {
	Body        []byte
	ContentType string
	// successful operation
	Pets        []shared.Pet
	StatusCode  int
	RawResponse *http.Response
}