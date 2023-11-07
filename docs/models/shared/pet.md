# Pet


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Category`                                            | [*shared.Category](../../models/shared/category.md)   | :heavy_minus_sign:                                    | N/A                                                   |                                                       |
| `ID`                                                  | **int64*                                              | :heavy_minus_sign:                                    | N/A                                                   | 10                                                    |
| `Name`                                                | *string*                                              | :heavy_check_mark:                                    | N/A                                                   | doggie                                                |
| `PhotoUrls`                                           | []*string*                                            | :heavy_check_mark:                                    | N/A                                                   |                                                       |
| `Status`                                              | [*shared.PetStatus](../../models/shared/petstatus.md) | :heavy_minus_sign:                                    | pet status in the store                               |                                                       |
| `Tags`                                                | [][shared.Tag](../../models/shared/tag.md)            | :heavy_minus_sign:                                    | N/A                                                   |                                                       |