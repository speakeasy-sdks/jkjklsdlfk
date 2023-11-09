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
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../pkg/models/shared/pet.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddPetFormResponse](../../pkg/models/operations/addpetformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## AddPetJSON

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../pkg/models/shared/pet.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.AddPetJSONResponse](../../pkg/models/operations/addpetjsonresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## AddPetRaw

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Pet.AddPetRaw(ctx, []byte("0xcf5E85CDde"))
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
| sdkerrors.SDKError | 400-600            | */*                |

## DeletePet

Deletes a pet

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
	"petstore/v2/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Pet.DeletePet(ctx, operations.DeletePetRequest{
        PetID: 441876,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeletePetRequest](../../pkg/models/operations/deletepetrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.DeletePetResponse](../../pkg/models/operations/deletepetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## FindPetsByStatus

Multiple status values can be provided with comma separated strings

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
	"petstore/v2/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Pet.FindPetsByStatus(ctx, operations.FindPetsByStatusRequest{})
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
| sdkerrors.SDKError | 400-600            | */*                |

## FindPetsByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
	"petstore/v2/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Pet.FindPetsByTags(ctx, operations.FindPetsByTagsRequest{
        Tags: []string{
            "string",
        },
    })
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
| sdkerrors.SDKError | 400-600            | */*                |

## GetPetByID

Returns a single pet

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/operations"
)

func main() {
    s := petstore.New()


    operationSecurity := operations.GetPetByIDSecurity{
            APIKey: petstore.String(""),
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
| sdkerrors.SDKError | 400-600            | */*                |

## UpdatePetWithForm

Updates a pet in the store with form data

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
	"petstore/v2/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Pet.UpdatePetWithForm(ctx, operations.UpdatePetWithFormRequest{
        PetID: 303241,
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdatePetWithFormRequest](../../pkg/models/operations/updatepetwithformrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdatePetWithFormResponse](../../pkg/models/operations/updatepetwithformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdatePetForm

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../pkg/models/shared/pet.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdatePetFormResponse](../../pkg/models/operations/updatepetformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdatePetJSON

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Pet](../../pkg/models/shared/pet.md)          | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdatePetJSONResponse](../../pkg/models/operations/updatepetjsonresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdatePetRaw

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Pet.UpdatePetRaw(ctx, []byte("0x6bCA76De67"))
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
| sdkerrors.SDKError | 400-600            | */*                |

## UploadFile

uploads an image

### Example Usage

```go
package main

import(
	"context"
	"log"
	petstore "petstore/v2"
	"petstore/v2/pkg/models/shared"
	"petstore/v2/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Pet.UploadFile(ctx, operations.UploadFileRequest{
        RequestBody: []byte("0xc7cca7F47D"),
        PetID: 621158,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UploadFileRequest](../../pkg/models/operations/uploadfilerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UploadFileResponse](../../pkg/models/operations/uploadfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
