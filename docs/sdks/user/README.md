# User

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
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
)

func main() {
    s := petstore.New()

    ctx := context.Background()
    res, err := s.User.CreateUserForm(ctx, shared.User{
        Email: petstore.String("john@email.com"),
        FirstName: petstore.String("John"),
        ID: petstore.Int64(10),
        LastName: petstore.String("James"),
        Password: petstore.String("12345"),
        Phone: petstore.String("12345"),
        UserStatus: petstore.Int(1),
        Username: petstore.String("theUser"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../models/shared/user.md)            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserFormResponse](../../models/operations/createuserformresponse.md), error**


## CreateUserJSON

This can only be done by the logged in user.

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
    s := petstore.New()

    ctx := context.Background()
    res, err := s.User.CreateUserJSON(ctx, shared.User{
        Email: petstore.String("john@email.com"),
        FirstName: petstore.String("John"),
        ID: petstore.Int64(10),
        LastName: petstore.String("James"),
        Password: petstore.String("12345"),
        Phone: petstore.String("12345"),
        UserStatus: petstore.Int(1),
        Username: petstore.String("theUser"),
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.User](../../models/shared/user.md)            | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUserJSONResponse](../../models/operations/createuserjsonresponse.md), error**


## CreateUserRaw

This can only be done by the logged in user.

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
    s := petstore.New()

    ctx := context.Background()
    res, err := s.User.CreateUserRaw(ctx, []byte("quasi"))
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
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

**[*operations.CreateUserRawResponse](../../models/operations/createuserrawresponse.md), error**


## CreateUsersWithListInput

Creates list of users with given input array

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
    s := petstore.New()

    ctx := context.Background()
    res, err := s.User.CreateUsersWithListInput(ctx, []shared.User{
        shared.User{
            Email: petstore.String("john@email.com"),
            FirstName: petstore.String("John"),
            ID: petstore.Int64(10),
            LastName: petstore.String("James"),
            Password: petstore.String("12345"),
            Phone: petstore.String("12345"),
            UserStatus: petstore.Int(1),
            Username: petstore.String("theUser"),
        },
        shared.User{
            Email: petstore.String("john@email.com"),
            FirstName: petstore.String("John"),
            ID: petstore.Int64(10),
            LastName: petstore.String("James"),
            Password: petstore.String("12345"),
            Phone: petstore.String("12345"),
            UserStatus: petstore.Int(1),
            Username: petstore.String("theUser"),
        },
        shared.User{
            Email: petstore.String("john@email.com"),
            FirstName: petstore.String("John"),
            ID: petstore.Int64(10),
            LastName: petstore.String("James"),
            Password: petstore.String("12345"),
            Phone: petstore.String("12345"),
            UserStatus: petstore.Int(1),
            Username: petstore.String("theUser"),
        },
        shared.User{
            Email: petstore.String("john@email.com"),
            FirstName: petstore.String("John"),
            ID: petstore.Int64(10),
            LastName: petstore.String("James"),
            Password: petstore.String("12345"),
            Phone: petstore.String("12345"),
            UserStatus: petstore.Int(1),
            Username: petstore.String("theUser"),
        },
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
| `request`                                             | [[]shared.User](../../models//.md)                    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.CreateUsersWithListInputResponse](../../models/operations/createuserswithlistinputresponse.md), error**


## DeleteUser

This can only be done by the logged in user.

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

    ctx := context.Background()
    res, err := s.User.DeleteUser(ctx, operations.DeleteUserRequest{
        Username: "Weston_Thiel",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.DeleteUserRequest](../../models/operations/deleteuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**


## GetUserByName

Get user by user name

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

    ctx := context.Background()
    res, err := s.User.GetUserByName(ctx, operations.GetUserByNameRequest{
        Username: "Whitney.Bednar",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetUserByNameRequest](../../models/operations/getuserbynamerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetUserByNameResponse](../../models/operations/getuserbynameresponse.md), error**


## LoginUser

Logs user into the system

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

    ctx := context.Background()
    res, err := s.User.LoginUser(ctx, operations.LoginUserRequest{
        Password: petstore.String("cum"),
        Username: petstore.String("Aiyana.Batz"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoginUser200ApplicationJSONString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.LoginUserRequest](../../models/operations/loginuserrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.LoginUserResponse](../../models/operations/loginuserresponse.md), error**


## LogoutUser

Logs out current logged in user session

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
)

func main() {
    s := petstore.New()

    ctx := context.Background()
    res, err := s.User.LogoutUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.LogoutUserResponse](../../models/operations/logoutuserresponse.md), error**


## UpdateUserForm

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/operations"
	"petstore/pkg/models/shared"
)

func main() {
    s := petstore.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserForm(ctx, operations.UpdateUserFormRequest{
        User: &shared.User{
            Email: petstore.String("john@email.com"),
            FirstName: petstore.String("John"),
            ID: petstore.Int64(10),
            LastName: petstore.String("James"),
            Password: petstore.String("12345"),
            Phone: petstore.String("12345"),
            UserStatus: petstore.Int(1),
            Username: petstore.String("theUser"),
        },
        Username: "Wilfrid.Carter",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateUserFormRequest](../../models/operations/updateuserformrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateUserFormResponse](../../models/operations/updateuserformresponse.md), error**


## UpdateUserJSON

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/operations"
	"petstore/pkg/models/shared"
)

func main() {
    s := petstore.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserJSON(ctx, operations.UpdateUserJSONRequest{
        User: &shared.User{
            Email: petstore.String("john@email.com"),
            FirstName: petstore.String("John"),
            ID: petstore.Int64(10),
            LastName: petstore.String("James"),
            Password: petstore.String("12345"),
            Phone: petstore.String("12345"),
            UserStatus: petstore.Int(1),
            Username: petstore.String("theUser"),
        },
        Username: "Jayden.Carter88",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateUserJSONRequest](../../models/operations/updateuserjsonrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateUserJSONResponse](../../models/operations/updateuserjsonresponse.md), error**


## UpdateUserRaw

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/operations"
	"petstore/pkg/models/shared"
)

func main() {
    s := petstore.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserRaw(ctx, operations.UpdateUserRawRequest{
        RequestBody: []byte("commodi"),
        Username: "Terrill69",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateUserRawRequest](../../models/operations/updateuserrawrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateUserRawResponse](../../models/operations/updateuserrawresponse.md), error**

