# Scoreboard Server

## Setup

### Test
`go test ./...`

### Compile
`go build -o server.exe`

### Run
`./server.exe -p <port>`

## HTTP Endpoints

### /scores

#### GET
Return a list of all scoreboards

##### Request
###### Query Parameters
| Parameter | Type    | Description                                       |
| --------- | ------- | ------------------------------------------------- |
| score-id  | string  | Filter response to only list specific scoreboards |
| featured  | boolean | Filter out non-featured scoreboards               |

###### Example
```GET /scores```

##### Responses
```json
200 OK
{
    "scoreboards": {
        "<score-id>": {
            "completed": <boolean>,
            "teams": [
                {
                    "name": <string>,
                    "score": <float>
                }
            ]
        }
    }
}
```

### /scores/{score-id}

#### PUT
Create or update a scoreboard

##### Request

###### Path Parameters
| Parameter | Type   | Description                       |
| --------- | ------ | --------------------------------- |
| score-id  | string | ID of scoreboard to create/update |

###### Payload Parameters
| Parameter  | Type    | Description                                       |
| ---------- | ------- | ------------------------------------------------- |
| completed  | boolean | Whether the game for the scoreboard has completed |
| teams      | array   | Array of team objects                             |
| team.name  | string  | Name of team                                      |
| team.score | float   | Score of team

###### Example
```json
PUT /scores/score_1
{
    "completed": true,
    "teams": [
        {
            "name": "Flet Inc.",
            "score": 795
        },
        {
            "name": "Pleb Corp",
            "score": 96
        }
    ]
}
```

##### Responses
```json
200 OK
{
    "score-id": <string>,
    "scoreboard": {
        "completed": <boolean>,
        "teams": [
            {
                "name": <string>,
                "score": <float>
            }
        ]
    }
}
```

```json
400 Bad Request
{
    "message": <string>
}
```

#### GET
Retrieve a specific scoreboard

##### Request

###### Path Parameters
| Parameter | Type   | Description                       |
| --------- | ------ | --------------------------------- |
| score-id  | string | ID of scoreboard to create/update |

###### Example
```GET /scores/score_1```

##### Responses
```json
200 OK
{
    "completed": <boolean>,
    "teams": [
        {
            "name": <string>,
            "score": <float>
        }
    ]
}
```

```json
404 Not Found
{
    "message": <string>
}
```

#### DELETE
Delete a specific scoreboard

##### Request

###### Path Parameters
| Parameter | Type   | Description                       |
| --------- | ------ | --------------------------------- |
| score-id  | string | ID of scoreboard to create/update |

###### Example
```DELETE /scores/score_1```

##### Responses
```json
200 OK
{
    "message": <string>
}
```

```json
404 Not Found
{
    "message": <string>
}
```

## WebSocket API
The scoreboard server publishes scoreboard updates via a WebSocket interface for real-time score monitoring

### Endpoint
`/live`

### Messages
Scoreboard update events are published as JSON in the following format:
```json
{
    "score-id": <string>,
    "scoreboard": {
        "completed": <boolean>,
        "teams": [
            {
                "name": <string>,
                "score": <float>
            }
        ]
    }
}
```