# UpdatePetRawResponse


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Body`                                                 | *[]byte*                                               | :heavy_minus_sign:                                     | N/A                                                    |
| `ContentType`                                          | *string*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `Pet`                                                  | [*shared.Pet](../../models/shared/pet.md)              | :heavy_minus_sign:                                     | Successful operation                                   |
| `StatusCode`                                           | *int*                                                  | :heavy_check_mark:                                     | N/A                                                    |
| `RawResponse`                                          | [*http.Response](https://pkg.go.dev/net/http#Response) | :heavy_minus_sign:                                     | N/A                                                    |