<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
)

func main() {
	s := petstore.New(
		petstore.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		Category: &shared.Category{
			ID:   petstore.Int64(1),
			Name: petstore.String("Dogs"),
		},
		ID:   petstore.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
		Tags: []shared.Tag{
			shared.Tag{},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->