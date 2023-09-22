<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            PetstoreAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pet.AddPetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: petstore.Int64(1),
            Name: petstore.String("Dogs"),
        },
        ID: petstore.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "corrupti",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: petstore.Int64(715190),
                Name: petstore.String("Stuart Stiedemann"),
            },
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
<!-- End SDK Example Usage -->