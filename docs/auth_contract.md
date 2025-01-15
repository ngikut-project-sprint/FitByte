# Authentication & Authorization

## PIC
Fardan & Deni

## Background:
User can register to the app and start tracking their fitness journey

## Contract:

### POST /v1/login

Request Body:
```json
{
  "email": "name@name.com",     // string | required | should be email format
  "password": "asdfasdf"        // string | required | minLength: 8 | maxLength: 32
}
```

Response:
- `200` Ok for existing user
```json
{
  "email": "name@name.com",
  "token": "asdfasdf"         // string | use any token you want
}
```
- `400` Bad Request case:
  - Validation error
- `404` Not Found case:
  - Email is not found
- `500` Server Error

### POST /v1/register

Request Body:
```json
{
  "email": "name@name.com",     // string | required | should be email format
  "password": "asdfasdf"        // string | required | minLength: 8 | maxLength: 32
}
```

Response:
- `201` Created for new user
```json
{
  "email": "name@name.com",
  "token": "asdfasdf"         // string | use any token you want
}
```
- `400` Bad Request case:
  - Validation error
- `409` Conflict case:
  - Email is exist
- `500` Server Error