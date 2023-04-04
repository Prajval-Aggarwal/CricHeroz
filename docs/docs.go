// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/DecisionUpdate": {
            "put": {
                "description": "Updates the decison taken by the team after wining the toss",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Toss"
                ],
                "parameters": [
                    {
                        "description": "Descision Updated",
                        "name": "toss",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Toss"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Toss"
                        }
                    }
                }
            }
        },
        "/addCareer": {
            "post": {
                "description": "Add player career",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Player"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Player ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Adds Player career",
                        "name": "playerCareer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Career"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Career"
                        }
                    }
                }
            }
        },
        "/addPlayertoTeam": {
            "post": {
                "description": "Add player to team",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Team"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the team",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Array of players",
                        "name": "player",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/addToScoreCard": {
            "post": {
                "description": "stores players info in scorecard",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Scorecard"
                ],
                "parameters": [
                    {
                        "description": "ScoreCard details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CardData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ScoreCard"
                        }
                    }
                }
            }
        },
        "/ballUpdate": {
            "post": {
                "description": "Update the ball",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ball"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the ball",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Balls"
                        }
                    }
                }
            }
        },
        "/createMatch": {
            "post": {
                "description": "Create the match between the teams",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Match"
                ],
                "parameters": [
                    {
                        "description": "Match details",
                        "name": "match",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Match"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Match"
                        }
                    }
                }
            }
        },
        "/createPlayer": {
            "post": {
                "description": "Creates a new Player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Player"
                ],
                "parameters": [
                    {
                        "description": "Create Player",
                        "name": "player",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Player"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Player"
                        }
                    }
                }
            }
        },
        "/createTeam": {
            "post": {
                "description": "Creates a team",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Team"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the user",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Details of the team",
                        "name": "TeamDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Team"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Team"
                        }
                    }
                }
            }
        },
        "/deleteTeamByID": {
            "delete": {
                "description": "Delete the team",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Team"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the team",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "ID of the user",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/endInning": {
            "post": {
                "description": "Ends the current team innings",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Inning"
                ],
                "parameters": [
                    {
                        "description": "Id of the team to end its inning",
                        "name": "team_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/endMatch": {
            "post": {
                "description": "Ends the match and updates the scorecard of every player",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Match"
                ],
                "parameters": [
                    {
                        "description": "Id of the match to end",
                        "name": "match_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Match"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login a user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "parameters": [
                    {
                        "description": "Log in the user",
                        "name": "UserDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Credential"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Registers a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "parameters": [
                    {
                        "description": "Registers a user",
                        "name": "UserDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Credential"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Credential"
                        }
                    }
                }
            }
        },
        "/showMatch": {
            "post": {
                "description": "Show the list of matches",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Match"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Match"
                        }
                    }
                }
            }
        },
        "/showPlayer": {
            "get": {
                "description": "Shows the list of all the player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Player"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Player"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/showPlayerID": {
            "get": {
                "description": "Shows the list of all the player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Player"
                ],
                "parameters": [
                    {
                        "description": "ID of the player",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/showScoreCard": {
            "post": {
                "description": "Shows the score card for the current matcha",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Scorecard"
                ],
                "parameters": [
                    {
                        "description": "Id of the match whose scoredcard is to be viewed",
                        "name": "match_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ScoreCard"
                        }
                    }
                }
            }
        },
        "/showTeamByID": {
            "post": {
                "description": "Shows the list of teams",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Team"
                ],
                "parameters": [
                    {
                        "description": "ID of the team",
                        "name": "team_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Player"
                        }
                    }
                }
            }
        },
        "/showTeams": {
            "get": {
                "description": "Shows the list of teams",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Team"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the User",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Team"
                        }
                    }
                }
            }
        },
        "/tossResult": {
            "post": {
                "description": "Give the random result of coin toss and which team won the toss",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Toss"
                ],
                "parameters": [
                    {
                        "description": "Toss Details",
                        "name": "toss",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Toss"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Toss"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Balls": {
            "type": "object",
            "properties": {
                "ball_count": {
                    "type": "integer"
                },
                "ball_id": {
                    "type": "string"
                },
                "ball_type": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "is_valid": {
                    "type": "string"
                },
                "match_id": {
                    "type": "string"
                },
                "over": {
                    "type": "number"
                },
                "player_id": {
                    "type": "string"
                },
                "runs": {
                    "description": "runs on that particular ball",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.CardData": {
            "type": "object",
            "properties": {
                "ball_type": {
                    "type": "string"
                },
                "baller": {
                    "type": "string"
                },
                "batsmen": {
                    "type": "string"
                },
                "match_id": {
                    "type": "string"
                },
                "prev_runs": {
                    "type": "integer"
                },
                "runs": {
                    "type": "integer"
                },
                "scorecard_id": {
                    "type": "string"
                }
            }
        },
        "models.Career": {
            "type": "object",
            "properties": {
                "average_score": {
                    "type": "number"
                },
                "balls_bowled": {
                    "description": "Balls Bowled",
                    "type": "integer"
                },
                "balls_faced": {
                    "type": "integer"
                },
                "bowling_average": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "economy": {
                    "description": "BowlSR     float64 ` + "`" + `json:\"bowl_sr\"` + "`" + ` //Bowling strike rate",
                    "type": "number"
                },
                "fifties": {
                    "description": "BatSR      float64 ` + "`" + `json:\"bat_sr\"` + "`" + ` //batting strike rate",
                    "type": "integer"
                },
                "fours": {
                    "type": "integer"
                },
                "highest_score": {
                    "description": "high score",
                    "type": "integer"
                },
                "hundreds": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "matches_played": {
                    "type": "integer"
                },
                "player_id": {
                    "type": "string"
                },
                "run_scored": {
                    "type": "integer"
                },
                "runs_conced": {
                    "description": "Runs Conceded",
                    "type": "integer"
                },
                "sixes": {
                    "type": "integer"
                },
                "two_hundreds": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "wickets": {
                    "type": "integer"
                }
            }
        },
        "models.Credential": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone_no": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Match": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "match_id": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string"
                },
                "scorecard_id": {
                    "description": "scorecard related to it",
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "team1_id": {
                    "type": "string"
                },
                "team2_id": {
                    "type": "string"
                },
                "text": {
                    "description": "who won the match/",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "venue": {
                    "type": "string"
                }
            }
        },
        "models.Player": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "jersey_no": {
                    "type": "integer"
                },
                "phone_no": {
                    "type": "string"
                },
                "player_age": {
                    "type": "integer"
                },
                "player_id": {
                    "type": "string"
                },
                "player_name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "career": {
                    "$ref": "#/definitions/models.Career"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "player": {
                    "$ref": "#/definitions/models.Player"
                },
                "teams": {
                    "$ref": "#/definitions/models.TeamList"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.ScoreCard": {
            "type": "object",
            "properties": {
                "balls_played": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "economy": {
                    "type": "number"
                },
                "fours": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_out": {
                    "type": "string"
                },
                "maiden_overs": {
                    "type": "integer"
                },
                "no_balls": {
                    "type": "integer"
                },
                "overs_bowled": {
                    "type": "integer"
                },
                "player_id": {
                    "type": "string"
                },
                "player_type": {
                    "type": "string"
                },
                "runScored": {
                    "type": "integer"
                },
                "runs_given": {
                    "type": "integer"
                },
                "scorecard_id": {
                    "type": "string"
                },
                "sixes": {
                    "type": "integer"
                },
                "strike_rate": {
                    "type": "number"
                },
                "updatedAt": {
                    "type": "string"
                },
                "wickets": {
                    "type": "integer"
                },
                "wide_balls": {
                    "type": "integer"
                }
            }
        },
        "models.Team": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "player_id": {
                    "type": "string"
                },
                "team_captain": {
                    "type": "string"
                },
                "team_id": {
                    "type": "string"
                },
                "team_name": {
                    "type": "string"
                },
                "team_type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.TeamList": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "p_id": {
                    "type": "string"
                },
                "t_id": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Toss": {
            "type": "object",
            "properties": {
                "decision": {
                    "type": "string"
                },
                "head_team": {
                    "type": "string"
                },
                "match_id": {
                    "type": "string"
                },
                "tail_team": {
                    "type": "string"
                },
                "toss_id": {
                    "type": "string"
                },
                "toss_won": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Cric Heros API",
	Description:      "API Documentation for Cric Heros",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}