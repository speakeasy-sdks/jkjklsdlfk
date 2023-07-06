# GetOrderByIDResponse


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Body`                                                 | *[]byte*                                               | :heavy_minus_sign:                                     | N/A                                                    |
| `ContentType`                                          | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `Order`                                                | [*shared.Order](../../models/shared/order.md)          | :heavy_minus_sign:                                     | successful operation                                   |
| `StatusCode`                                           | *int*                                                  | :heavy_check_mark:                                     | N/A                                                    |
| `RawResponse`                                          | [*http.Response](https://pkg.go.dev/net/http#Response) | :heavy_minus_sign:                                     | N/A                                                    |