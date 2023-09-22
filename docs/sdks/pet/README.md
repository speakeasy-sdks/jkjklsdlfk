# Pet

## Overview

Everything about your Pets

Find out more
<http://swagger.io>
### Available Operations

* [AddPetForm](#addpetform) - Add a new pet to the store
* [AddPetJSON](#addpetjson) - Add a new pet to the store
* [AddPetRaw](#addpetraw) - Add a new pet to the store
* [DeletePet](#deletepet) - Deletes a pet
* [FindPetsByStatus](#findpetsbystatus) - Finds Pets by status
* [FindPetsByTags](#findpetsbytags) - Finds Pets by tags
* [GetPetByID](#getpetbyid) - Find pet by ID
* [UpdatePetWithForm](#updatepetwithform) - Updates a pet in the store with form data
* [UpdatePetForm](#updatepetform) - Update an existing pet
* [UpdatePetJSON](#updatepetjson) - Update an existing pet
* [UpdatePetRaw](#updatepetraw) - Update an existing pet
* [UploadFile](#uploadfile) - uploads an image

## AddPetForm

Add a new pet to the store

### Example Usage

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
            "tempora",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: petstore.Int64(477665),
                Name: petstore.String("Irving Lehner"),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../models/shared/pet.md)              | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddPetFormResponse](../../models/operations/addpetformresponse.md), error**


## AddPetJSON

Add a new pet to the store

### Example Usage

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
    res, err := s.Pet.AddPetJSON(ctx, shared.Pet{
        Category: &shared.Category{
            ID: petstore.Int64(1),
            Name: petstore.String("Dogs"),
        },
        ID: petstore.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "nisi",
        },
        Status: shared.PetStatusSold.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: petstore.Int64(836079),
                Name: petstore.String("Erica Bogisich III"),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../models/shared/pet.md)              | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddPetJSONResponse](../../models/operations/addpetjsonresponse.md), error**


## AddPetRaw

Add a new pet to the store

### Example Usage

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
    res, err := s.Pet.AddPetRaw(ctx, []byte("repellendus"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddPetRawResponse](../../models/operations/addpetrawresponse.md), error**


## DeletePet

Deletes a pet

### Example Usage

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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            PetstoreAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pet.DeletePet(ctx, operations.DeletePetRequest{
        APIKey: petstore.String("sapiente"),
        PetID: 778157,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.DeletePetRequest](../../models/operations/deletepetrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.DeletePetResponse](../../models/operations/deletepetresponse.md), error**


## FindPetsByStatus

Multiple status values can be provided with comma separated strings

### Example Usage

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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            PetstoreAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pet.FindPetsByStatus(ctx, operations.FindPetsByStatusRequest{
        Status: operations.FindPetsByStatusStatusAvailable.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.FindPetsByStatusRequest](../../models/operations/findpetsbystatusrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.FindPetsByStatusResponse](../../models/operations/findpetsbystatusresponse.md), error**


## FindPetsByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example Usage

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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            PetstoreAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pet.FindPetsByTags(ctx, operations.FindPetsByTagsRequest{
        Tags: []string{
            "at",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.FindPetsByTagsRequest](../../models/operations/findpetsbytagsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.FindPetsByTagsResponse](../../models/operations/findpetsbytagsresponse.md), error**


## GetPetByID

Returns a single pet

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/operations"
)

func main() {
    s := petstore.New()
    operationSecurity := operations.GetPetByIDSecurity{
            APIKey: petstore.String(""),
        }

    ctx := context.Background()
    res, err := s.Pet.GetPetByID(ctx, operations.GetPetByIDRequest{
        PetID: 870088,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetPetByIDRequest](../../models/operations/getpetbyidrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.GetPetByIDSecurity](../../models/operations/getpetbyidsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.GetPetByIDResponse](../../models/operations/getpetbyidresponse.md), error**


## UpdatePetWithForm

Updates a pet in the store with form data

### Example Usage

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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            PetstoreAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pet.UpdatePetWithForm(ctx, operations.UpdatePetWithFormRequest{
        Name: petstore.String("Javier Schmidt"),
        PetID: 520478,
        Status: petstore.String("porro"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdatePetWithFormRequest](../../models/operations/updatepetwithformrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpdatePetWithFormResponse](../../models/operations/updatepetwithformresponse.md), error**


## UpdatePetForm

Update an existing pet by Id

### Example Usage

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
    res, err := s.Pet.UpdatePetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: petstore.Int64(1),
            Name: petstore.String("Dogs"),
        },
        ID: petstore.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "dolorum",
        },
        Status: shared.PetStatusAvailable.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: petstore.Int64(720633),
                Name: petstore.String("Seth Conroy"),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../models/shared/pet.md)              | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdatePetFormResponse](../../models/operations/updatepetformresponse.md), error**


## UpdatePetJSON

Update an existing pet by Id

### Example Usage

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
    res, err := s.Pet.UpdatePetJSON(ctx, shared.Pet{
        Category: &shared.Category{
            ID: petstore.Int64(1),
            Name: petstore.String("Dogs"),
        },
        ID: petstore.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "optio",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: petstore.Int64(105907),
                Name: petstore.String("Tanya Gleason"),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../models/shared/pet.md)              | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdatePetJSONResponse](../../models/operations/updatepetjsonresponse.md), error**


## UpdatePetRaw

Update an existing pet by Id

### Example Usage

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
    res, err := s.Pet.UpdatePetRaw(ctx, []byte("cum"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdatePetRawResponse](../../models/operations/updatepetrawresponse.md), error**


## UploadFile

uploads an image

### Example Usage

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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            PetstoreAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Pet.UploadFile(ctx, operations.UploadFileRequest{
        RequestBody: []byte("esse"),
        AdditionalMetadata: petstore.String("ipsum"),
        PetID: 568434,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UploadFileRequest](../../models/operations/uploadfilerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.UploadFileResponse](../../models/operations/uploadfileresponse.md), error**

