# Refactoring Go with Database

## Instuctions
1. copy `.env.example` to `.env` and fill in the values
1. if you don't have [direnv](https://direnv.net/) installed then `source .env` to load the environment variables;
1. `make start-db` to start the database
1. `make run` to run the server
1. `make get-skill` to test get skill `go` from the server

## e2e
1. `make e2e` to run the e2e tests if returned `true` then the tests passed. this is a simple e2e to check just the get skill by key endpoint

## Endpoints
- `GET /skills/:key` - get a skill by key

response

```json
	{
  "data": {
    "key": "go",
    "name": "Go",
    "description": "Go is an open source programming...",
    "logo": "",
    "levels": [
      {
        "key": "",
        "name": "Beginner",
        "brief": "",
        "descriptions": [
          "basic knowledge ..."
        ],
        "level": 1
      },
      {
        "key": "",
        "name": "Intermediate",
        "brief": "",
        "descriptions": [
          "complex programs..."
        ],
        "level": 2
      }
    ],
    "tags": [
      "go",
      "golang"
    ]
  }
}
```

- `GET /skills` - get all skills

response

```json
	{
  "data": [{
    "key": "go",
    "name": "Go",
    "description": "Go is an open source programming...",
    "logo": "",
    "levels": [
      {
        "key": "",
        "name": "Beginner",
        "brief": "",
        "descriptions": [
          "basic knowledge ..."
        ],
        "level": 1
      },
      {
        "key": "",
        "name": "Intermediate",
        "brief": "",
        "descriptions": [
          "complex programs..."
        ],
        "level": 2
      }
    ],
    "tags": [
      "go",
      "golang"
    ]
  }]
}
```

- `POST /skills` - create a new skill

payload

```json
{
	"key": "go",
	"name": "Go",
	"description": "Go is an open source programming...",
	"logo": "",
	"levels": [
		{
			"key": "",
			"name": "Beginner",
			"brief": "",
			"descriptions": [
				"basic knowledge ..."
			],
			"level": 1
		},
		{
			"key": "",
			"name": "Intermediate",
			"brief": "",
			"descriptions": [
				"complex programs..."
			],
			"level": 2
		}
	],
	"tags": [
		"go",
		"golang"
	]
}
```