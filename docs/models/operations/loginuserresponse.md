# LoginUserResponse


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ContentType`                                          | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `Headers`                                              | map[string][]*string*                                  | :heavy_minus_sign:                                     | N/A                                                    |
| `StatusCode`                                           | *int*                                                  | :heavy_check_mark:                                     | N/A                                                    |
| `RawResponse`                                          | [*http.Response](https://pkg.go.dev/net/http#Response) | :heavy_minus_sign:                                     | N/A                                                    |
| `LoginUser200ApplicationJSONString`                    | **string*                                              | :heavy_minus_sign:                                     | successful operation                                   |
| `LoginUser200ApplicationXMLString`                     | **string*                                              | :heavy_minus_sign:                                     | successful operation                                   |