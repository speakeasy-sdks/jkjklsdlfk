<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/models/operations"
)

func main() {
    s := petstore.New()

    ctx := context.Background()
    res, err := s.Pet.AddPetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: petstore.Int64(1),
            Name: petstore.String("Dogs"),
        },
        ID: petstore.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "provident",
            "distinctio",
            "quibusdam",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: petstore.Int64(544883),
                Name: petstore.String("Ben Mueller"),
            },
            shared.Tag{
                ID: petstore.Int64(437587),
                Name: petstore.String("Raquel Bednar"),
            },
            shared.Tag{
                ID: petstore.Int64(383441),
                Name: petstore.String("Alexandra Schulist"),
            },
            shared.Tag{
                ID: petstore.Int64(568045),
                Name: petstore.String("Mrs. Sophie Smith MD"),
            },
        },
    }, operations.AddPetFormSecurity{
        PetstoreAuth: "",
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