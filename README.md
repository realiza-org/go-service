# go-service

### Project Structure
````
go-service/
├── cmd/
│   └── main.go          // Entry point for the application
├── internal/
│   ├── task/
│   │   ├── handler.go   // HTTP handlers for tasks
│   │   ├── service.go   // Business logic for tasks
│   │   ├── repository.go// Data access layer for tasks
│   │   └── model.go     // Task model definition
│   └── server/
│       └── server.go    // Server initialization
└── go.mod               // Go module file
````


### Running the Project

Run the application:
````
go run cmd/main.go
````

The server will start on http://localhost:8080.

### Testing the API

- Use tools like curl or Postman for testing:

```
curl -X POST -H "Content-Type: application/json" -d '{"title": "Learn Go", "description": "Study Go Echo framework"}' http://localhost:8080/tasks
```

- Get all tasks:

```
curl http://localhost:8080/tasks
```
