# Activity

## PIC
...

## Background:
User can manage their activities

## Contract:

### POST /v1/activity

Request Header:
| key | value |
|-----|--------|
| Authorization | bearer ... |

Request Body:
```json
{
  "activityType": "",           // required | enum based on Activity Types table
  "doneAt": "",                // string | required | ISO Date
  "durationInMinutes": 1       // number | required | min: 1
}
```

Response:
- `201` Created
```json
{
  "activityId": "",           // string | Use any id you want
  "activityType": "",
  "doneAt": "",
  "durationInMinutes": 1,
  "caloriesBurned": 100,      // calculated based on Activity Types table
  "createdAt": "",           // string | ISO Date
  "updatedAt": ""            // string | ISO Date
}
```
- `400` Bad Request case:
  - Validation error
- `401` Unauthorized for:
  - expired / invalid / missing request token
- `500` Server Error

### GET /v1/activity

Request parameters (all optional)
- `limit` & `offset` limit the output of the data
  - default `limit=5&offset=0`
  - value should be a number
  - invalid `limit` / `offset` value will use the default value
- `activityType` filter the result based on activity type
  - search should be a enum based on activity table
  - value should be a string
  - invalid value will be ignored
- `doneAtFrom` filter the result based on doneAt
  - value should be a ISO Date
  - searches activities that are done after or equal doneAt
  - invalid value will be ignored
- `doneAtTo` filter the result based on doneAt
  - value should be a ISO Date
  - searches activities that are done before or equal doneAt
  - invalid value will be ignored
- `caloriesBurnedMin` filter the result based on caloriesBurned
  - value should be a number
  - searches activities that have more or equal calories burned than caloriesBurnedMin
  - invalid value will be ignored
- `caloriesBurnedMax` filter the result based on caloriesBurned
  - value should be a number
  - searches activities that have less or equal caloriesBurned than caloriesBurnedMax
  - invalid value will be ignored

Response:
- `200` Ok
```json
[
  {
    "activityId": "",
    "activityType": "",
    "doneAt": "",
    "durationInMinutes": 1,
    "caloriesBurned": 100,
    "createdAt": "",          // string | ISO Date
    "updatedAt": ""           // string | ISO Date
  }
]
```
- `401` Unauthorized for:
  - expired / invalid / missing request token
- `500` Server Error

### PATCH /v1/activity/:activityId

Request Header:
| key | value |
|-----|--------|
| Authorization | bearer ... |

Request Body:
```json
{
  "activityType": "",          // not required | enum based on Activity Types table
  "doneAt": "",               // string | not required | ISO Date
  "durationInMinutes": 1      // number | not required | min: 1
}
```

Response:
- `200` Ok
```json
{
  "activityId": "",
  "activityType": "",
  "doneAt": "",
  "durationInMinutes": 1,
  "caloriesBurned": 100,      // calculate based on Activity Types table
  "createdAt": "",           // string | ISO Date
  "updatedAt": ""            // string | ISO Date | should be equal to now
}
```
- `400` Bad Request case:
  - Validation error
- `404` Not Found case:
  - activityId is not found
- `401` Unauthorized for:
  - expired / invalid / missing request token
- `500` Server Error

### DELETE /v1/activity/:activityId

Request Header:
| key | value |
|-----|--------|
| Authorization | bearer ... |

Response:
- `200` Ok deleted
- `404` Not Found case:
  - activityId is not found
- `401` Unauthorized for:
  - expired / invalid / missing request token
- `500` Server Error