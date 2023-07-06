# petstore

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/jkjklsdlfk
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Pet](docs/sdks/pet/README.md)

* [AddPetForm](docs/sdks/pet/README.md#addpetform) - Add a new pet to the store
* [AddPetJSON](docs/sdks/pet/README.md#addpetjson) - Add a new pet to the store
* [AddPetRaw](docs/sdks/pet/README.md#addpetraw) - Add a new pet to the store
* [DeletePet](docs/sdks/pet/README.md#deletepet) - Deletes a pet
* [FindPetsByStatus](docs/sdks/pet/README.md#findpetsbystatus) - Finds Pets by status
* [FindPetsByTags](docs/sdks/pet/README.md#findpetsbytags) - Finds Pets by tags
* [GetPetByID](docs/sdks/pet/README.md#getpetbyid) - Find pet by ID
* [UpdatePetWithForm](docs/sdks/pet/README.md#updatepetwithform) - Updates a pet in the store with form data
* [UpdatePetForm](docs/sdks/pet/README.md#updatepetform) - Update an existing pet
* [UpdatePetJSON](docs/sdks/pet/README.md#updatepetjson) - Update an existing pet
* [UpdatePetRaw](docs/sdks/pet/README.md#updatepetraw) - Update an existing pet
* [UploadFile](docs/sdks/pet/README.md#uploadfile) - uploads an image

### [Store](docs/sdks/store/README.md)

* [DeleteOrder](docs/sdks/store/README.md#deleteorder) - Delete purchase order by ID
* [GetInventory](docs/sdks/store/README.md#getinventory) - Returns pet inventories by status
* [GetOrderByID](docs/sdks/store/README.md#getorderbyid) - Find purchase order by ID
* [PlaceOrderForm](docs/sdks/store/README.md#placeorderform) - Place an order for a pet
* [PlaceOrderJSON](docs/sdks/store/README.md#placeorderjson) - Place an order for a pet
* [PlaceOrderRaw](docs/sdks/store/README.md#placeorderraw) - Place an order for a pet

### [User](docs/sdks/user/README.md)

* [CreateUserForm](docs/sdks/user/README.md#createuserform) - Create user
* [CreateUserJSON](docs/sdks/user/README.md#createuserjson) - Create user
* [CreateUserRaw](docs/sdks/user/README.md#createuserraw) - Create user
* [CreateUsersWithListInput](docs/sdks/user/README.md#createuserswithlistinput) - Creates list of users with given input array
* [DeleteUser](docs/sdks/user/README.md#deleteuser) - Delete user
* [GetUserByName](docs/sdks/user/README.md#getuserbyname) - Get user by user name
* [LoginUser](docs/sdks/user/README.md#loginuser) - Logs user into the system
* [LogoutUser](docs/sdks/user/README.md#logoutuser) - Logs out current logged in user session
* [UpdateUserForm](docs/sdks/user/README.md#updateuserform) - Update user
* [UpdateUserJSON](docs/sdks/user/README.md#updateuserjson) - Update user
* [UpdateUserRaw](docs/sdks/user/README.md#updateuserraw) - Update user
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
