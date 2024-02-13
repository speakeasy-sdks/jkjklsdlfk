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
		petstore.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Pet.AddPetForm(ctx, shared.Pet{
		ID:   petstore.Int64(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
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