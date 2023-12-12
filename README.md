# petstore

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/jkjklsdlfk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
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
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/sdkerrors"
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

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https:///api/v3` | None |

#### Example

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
		petstore.WithServerIndex(0),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		petstore.WithServerURL("https:///api/v3"),
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
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name           | Type           | Scheme         |
| -------------- | -------------- | -------------- |
| `PetstoreAuth` | oauth2         | OAuth2 token   |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/operations"
)

func main() {
	s := petstore.New()

	operationSecurity := operations.GetPetByIDSecurity{
		APIKey: petstore.String("<YOUR_API_KEY_HERE>"),
	}

	ctx := context.Background()
	res, err := s.Pet.GetPetByID(ctx, operations.GetPetByIDRequest{
		PetID: 504151,
	}, operationSecurity)
	if err != nil {
		log.Fatal(err)
	}

	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
