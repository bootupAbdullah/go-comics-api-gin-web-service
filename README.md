# go-comics-api-tutorial-web-service-gin

This is a beginner-level Go project that builds a simple web service using the Gin web framework.

The purpose of this project is to get hands-on practice with:

- Structs and JSON responses in Go
- Basic HTTP routing using Gin
- Creating a `GET /comics` endpoint
- Understanding how web services work under the hood
- Learning by applying examples from the Gin and Go documentation

### Features

- `GET /comics` route returns a list of hardcoded comic book data
- Uses Gin’s `IndentedJSON` to return clean, readable output
- Project was built entirely for learning purposes

### Technologies

- Go (Golang)
- Gin Web Framework

### Example JSON Response

```json
[
  {
    "id": "1",
    "title": "The Thing",
    "artist": "John Byrne",
    "issue": "#1",
    "price": 8.0
  },
  {
    "id": "2",
    "title": "Fantastic Four",
    "artist": "John Byrne",
    "issue": "#220",
    "price": 20.0
  }
]
