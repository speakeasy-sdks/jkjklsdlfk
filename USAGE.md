<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	petstore "petstore/v3"
	"petstore/v3/pkg/models/shared"
)

func main() {
	s := petstore.New(
		petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
	)

	request := shared.Pet{
		Category: &shared.Category{
			ID:   petstore.Int64(1),
			Name: petstore.String("Dogs"),
		},
		ID:   petstore.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"<value>",
		},
	}

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->