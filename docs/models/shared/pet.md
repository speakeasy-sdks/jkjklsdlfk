# Pet

Create a new pet in the store


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    | Example                                        |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `Category`                                     | [*Category](../../models/shared/category.md)   | :heavy_minus_sign:                             | N/A                                            |                                                |
| `ID`                                           | **int64*                                       | :heavy_minus_sign:                             | N/A                                            | 10                                             |
| `Name`                                         | *string*                                       | :heavy_check_mark:                             | N/A                                            | doggie                                         |
| `PhotoUrls`                                    | []*string*                                     | :heavy_check_mark:                             | N/A                                            |                                                |
| `Status`                                       | [*PetStatus](../../models/shared/petstatus.md) | :heavy_minus_sign:                             | pet status in the store                        |                                                |
| `Tags`                                         | [][Tag](../../models/shared/tag.md)            | :heavy_minus_sign:                             | N/A                                            |                                                |