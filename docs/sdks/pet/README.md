# Pet
(*Pet*)

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
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    request := shared.Pet{
        Category: &shared.Category{
            ID: petstore.Int64(1),
            Name: petstore.String("Dogs"),
        },
        ID: petstore.Int64(10),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../pkg/models/shared/pet.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddPetFormResponse](../../pkg/models/operations/addpetformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AddPetJSON

Add a new pet to the store

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    request := shared.Pet{
        Category: &shared.Category{
            ID: petstore.Int64(1),
            Name: petstore.String("Dogs"),
        },
        ID: petstore.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "<value>",
        },
    }
    
    ctx := context.Background()
    res, err := s.Pet.AddPetJSON(ctx, request)
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
| `request`                                             | [shared.Pet](../../pkg/models/shared/pet.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddPetJSONResponse](../../pkg/models/operations/addpetjsonresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AddPetRaw

Add a new pet to the store

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    var request []byte = []byte("0xcf5E85CDde")
    
    ctx := context.Background()
    res, err := s.Pet.AddPetRaw(ctx, request)
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
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddPetRawResponse](../../pkg/models/operations/addpetrawresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeletePet

Deletes a pet

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"petstore/v3/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    request := operations.DeletePetRequest{
        PetID: 441876,
    }
    
    ctx := context.Background()
    res, err := s.Pet.DeletePet(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeletePetRequest](../../pkg/models/operations/deletepetrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.DeletePetResponse](../../pkg/models/operations/deletepetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## FindPetsByStatus

Multiple status values can be provided with comma separated strings

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"petstore/v3/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    request := operations.FindPetsByStatusRequest{}
    
    ctx := context.Background()
    res, err := s.Pet.FindPetsByStatus(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONClasses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.FindPetsByStatusRequest](../../pkg/models/operations/findpetsbystatusrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.FindPetsByStatusResponse](../../pkg/models/operations/findpetsbystatusresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## FindPetsByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"petstore/v3/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    request := operations.FindPetsByTagsRequest{}
    
    ctx := context.Background()
    res, err := s.Pet.FindPetsByTags(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONClasses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.FindPetsByTagsRequest](../../pkg/models/operations/findpetsbytagsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.FindPetsByTagsResponse](../../pkg/models/operations/findpetsbytagsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetPetByID

Returns a single pet

### Example Usage

```go
package main

import(
	petstore "petstore/v3"
	"petstore/v3/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := petstore.New()

    request := operations.GetPetByIDRequest{
        PetID: 504151,
    }

    security := operations.GetPetByIDSecurity{
            APIKey: petstore.String("<YOUR_API_KEY_HERE>"),
        }
    
    ctx := context.Background()
    res, err := s.Pet.GetPetByID(ctx, request, security)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetPetByIDRequest](../../pkg/models/operations/getpetbyidrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.GetPetByIDSecurity](../../pkg/models/operations/getpetbyidsecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.GetPetByIDResponse](../../pkg/models/operations/getpetbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdatePetWithForm

Updates a pet in the store with form data

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"petstore/v3/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    request := operations.UpdatePetWithFormRequest{
        PetID: 303241,
    }
    
    ctx := context.Background()
    res, err := s.Pet.UpdatePetWithForm(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdatePetWithFormRequest](../../pkg/models/operations/updatepetwithformrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdatePetWithFormResponse](../../pkg/models/operations/updatepetwithformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdatePetForm

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    request := shared.Pet{
        Category: &shared.Category{
            ID: petstore.Int64(1),
            Name: petstore.String("Dogs"),
        },
        ID: petstore.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "<value>",
        },
    }
    
    ctx := context.Background()
    res, err := s.Pet.UpdatePetForm(ctx, request)
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
| `request`                                             | [shared.Pet](../../pkg/models/shared/pet.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdatePetFormResponse](../../pkg/models/operations/updatepetformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdatePetJSON

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    request := shared.Pet{
        Category: &shared.Category{
            ID: petstore.Int64(1),
            Name: petstore.String("Dogs"),
        },
        ID: petstore.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "<value>",
        },
    }
    
    ctx := context.Background()
    res, err := s.Pet.UpdatePetJSON(ctx, request)
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
| `request`                                             | [shared.Pet](../../pkg/models/shared/pet.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdatePetJSONResponse](../../pkg/models/operations/updatepetjsonresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdatePetRaw

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    var request []byte = []byte("0x6bCA76De67")
    
    ctx := context.Background()
    res, err := s.Pet.UpdatePetRaw(ctx, request)
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
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdatePetRawResponse](../../pkg/models/operations/updatepetrawresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UploadFile

uploads an image

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"petstore/v3/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("<YOUR_PETSTORE_AUTH_HERE>"),
    )

    request := operations.UploadFileRequest{
        PetID: 565380,
    }
    
    ctx := context.Background()
    res, err := s.Pet.UploadFile(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UploadFileRequest](../../pkg/models/operations/uploadfilerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UploadFileResponse](../../pkg/models/operations/uploadfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
