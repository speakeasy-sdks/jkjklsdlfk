# User
(*User*)

## Overview

Operations about user

### Available Operations

* [CreateUserForm](#createuserform) - Create user
* [CreateUserJSON](#createuserjson) - Create user
* [CreateUserRaw](#createuserraw) - Create user
* [CreateUsersWithListInput](#createuserswithlistinput) - Creates list of users with given input array
* [DeleteUser](#deleteuser) - Delete user
* [GetUserByName](#getuserbyname) - Get user by user name
* [LoginUser](#loginuser) - Logs user into the system
* [LogoutUser](#logoutuser) - Logs out current logged in user session
* [UpdateUserForm](#updateuserform) - Update user
* [UpdateUserJSON](#updateuserjson) - Update user
* [UpdateUserRaw](#updateuserraw) - Update user

## CreateUserForm

This can only be done by the logged in user.

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
    res, err := s.User.CreateUserForm(ctx, &shared.User{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../pkg/models/shared/user.md)        | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserFormResponse](../../pkg/models/operations/createuserformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateUserJSON

This can only be done by the logged in user.

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
    res, err := s.User.CreateUserJSON(ctx, &shared.User{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../pkg/models/shared/user.md)        | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserJSONResponse](../../pkg/models/operations/createuserjsonresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateUserRaw

This can only be done by the logged in user.

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
    res, err := s.User.CreateUserRaw(ctx, []byte("0xB4dDB1Eeed"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.CreateUserRawResponse](../../pkg/models/operations/createuserrawresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateUsersWithListInput

Creates list of users with given input array

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
    res, err := s.User.CreateUsersWithListInput(ctx, []shared.User{
        shared.User{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]shared.User](../../.md)                            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUsersWithListInputResponse](../../pkg/models/operations/createuserswithlistinputresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## DeleteUser

This can only be done by the logged in user.

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
    res, err := s.User.DeleteUser(ctx, operations.DeleteUserRequest{
        Username: "Demetris_Torphy",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteUserRequest](../../pkg/models/operations/deleteuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteUserResponse](../../pkg/models/operations/deleteuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetUserByName

Get user by user name

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
    res, err := s.User.GetUserByName(ctx, operations.GetUserByNameRequest{
        Username: "Zachery_Schneider",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetUserByNameRequest](../../pkg/models/operations/getuserbynamerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetUserByNameResponse](../../pkg/models/operations/getuserbynameresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## LoginUser

Logs user into the system

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
    res, err := s.User.LoginUser(ctx, operations.LoginUserRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.String != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.LoginUserRequest](../../pkg/models/operations/loginuserrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.LoginUserResponse](../../pkg/models/operations/loginuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## LogoutUser

Logs out current logged in user session

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
    res, err := s.User.LogoutUser(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.LogoutUserResponse](../../pkg/models/operations/logoutuserresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateUserForm

This can only be done by the logged in user.

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
    res, err := s.User.UpdateUserForm(ctx, operations.UpdateUserFormRequest{
        Username: "Bo_Lynch4",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateUserFormRequest](../../pkg/models/operations/updateuserformrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateUserFormResponse](../../pkg/models/operations/updateuserformresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateUserJSON

This can only be done by the logged in user.

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
    res, err := s.User.UpdateUserJSON(ctx, operations.UpdateUserJSONRequest{
        Username: "Alanna_Waters81",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateUserJSONRequest](../../pkg/models/operations/updateuserjsonrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateUserJSONResponse](../../pkg/models/operations/updateuserjsonresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateUserRaw

This can only be done by the logged in user.

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
    res, err := s.User.UpdateUserRaw(ctx, operations.UpdateUserRawRequest{
        Username: "Maximus.DuBuque29",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateUserRawRequest](../../pkg/models/operations/updateuserrawrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateUserRawResponse](../../pkg/models/operations/updateuserrawresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
