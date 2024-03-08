# Store
(*Store*)

## Overview

Access to Petstore orders

Find out more about our store
<http://swagger.io>
### Available Operations

* [DeleteOrder](#deleteorder) - Delete purchase order by ID
* [GetInventory](#getinventory) - Returns pet inventories by status
* [GetOrderByID](#getorderbyid) - Find purchase order by ID
* [PlaceOrderForm](#placeorderform) - Place an order for a pet
* [PlaceOrderJSON](#placeorderjson) - Place an order for a pet
* [PlaceOrderRaw](#placeorderraw) - Place an order for a pet

## DeleteOrder

For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"context"
	"petstore/v3/pkg/models/operations"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Store.DeleteOrder(ctx, operations.DeleteOrderRequest{
        OrderID: 127902,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteOrderRequest](../../pkg/models/operations/deleteorderrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.DeleteOrderResponse](../../pkg/models/operations/deleteorderresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetInventory

Returns a map of status codes to quantities

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


    operationSecurity := operations.GetInventorySecurity{
            APIKey: "<YOUR_API_KEY_HERE>",
        }

    ctx := context.Background()
    res, err := s.Store.GetInventory(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `security`                                                                             | [operations.GetInventorySecurity](../../pkg/models/operations/getinventorysecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetInventoryResponse](../../pkg/models/operations/getinventoryresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetOrderByID

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

### Example Usage

```go
package main

import(
	"petstore/v3/pkg/models/shared"
	petstore "petstore/v3"
	"context"
	"petstore/v3/pkg/models/operations"
	"log"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Store.GetOrderByID(ctx, operations.GetOrderByIDRequest{
        OrderID: 614993,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetOrderByIDRequest](../../pkg/models/operations/getorderbyidrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetOrderByIDResponse](../../pkg/models/operations/getorderbyidresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PlaceOrderForm

Place a new order in the store

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
        petstore.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Store.PlaceOrderForm(ctx, &shared.Order{
        ID: petstore.Int64(10),
        PetID: petstore.Int64(198772),
        Quantity: petstore.Int(7),
        Status: shared.StatusApproved.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Order](../../pkg/models/shared/order.md)      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PlaceOrderFormResponse](../../pkg/models/operations/placeorderformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PlaceOrderJSON

Place a new order in the store

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
        petstore.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Store.PlaceOrderJSON(ctx, &shared.Order{
        ID: petstore.Int64(10),
        PetID: petstore.Int64(198772),
        Quantity: petstore.Int(7),
        Status: shared.StatusApproved.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Order](../../pkg/models/shared/order.md)      | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PlaceOrderJSONResponse](../../pkg/models/operations/placeorderjsonresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PlaceOrderRaw

Place a new order in the store

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
        petstore.WithSecurity("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Store.PlaceOrderRaw(ctx, []byte("0xcB9dC14dEe"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
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

**[*operations.PlaceOrderRawResponse](../../pkg/models/operations/placeorderrawresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
