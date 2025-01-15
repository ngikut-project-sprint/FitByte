# File Upload

## PIC
Dimas

## Background:
User can upload files

## Contract:

### POST /v1/file

Request Header:
| key | value |
|-----|--------|
| Authorization | bearer ... |

Request Multipart Form-Data:
| key | value |
|-----|--------|
| file | file \| should be jpeg/jpg \| max 100KB |

Response:
- `200` Ok
```json
{
  "uri": "name@name.com/file.jpg"  // should be the URI from storage
}
```
- `400` Bad Request case:
  - Validation error
- `401` Unauthorized for:
  - expired / invalid / missing request token
- `500` Server Error