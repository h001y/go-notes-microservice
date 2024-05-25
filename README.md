# go-notes-microservice

Here is will be annotations and notes about the microservice architecture and the Go programming language.

# Structure of the project


## Directory Structure 


```
go-notes-microservice/
├── cmd/
│ └── main.go
├── internal/
│ ├── handler/
│ │ └── note_handler.go
│ ├── model/
│ │ └── note.go
│ ├── repository/
│ │ └── note_repository.go
│ ├── service/
│ │ └── note_service.go
│ ├── router/
│ │ └── router.go
├── pkg/
│ ├── config/
│ │ └── config.go
│ ├── middleware/
│ │ └── auth.go
│ └── util/
│ └── util.go
├── go.mod
└── go.sum
```


## Description of Directories and Files

### `cmd/`
- **Purpose**: Contains the entry point for the application.
- **Files**: 
  - `main.go`: The main file to run the application.

### `internal/`
- **Purpose**: Contains internal packages that implement core functionality of the application.
  
  #### `handler/`
  - **Purpose**: Handles HTTP requests and responses.
  - **Files**: 
    - `note_handler.go`: Defines handlers for endpoints related to notes (create, read, update, delete).

  #### `model/`
  - **Purpose**: Defines the data structures used in the application.
  - **Files**: 
    - `note.go`: Defines the structure of a note.

  #### `repository/`
  - **Purpose**: Handles data storage and retrieval.
  - **Files**: 
    - `note_repository.go`: Implements CRUD operations for notes in the database.

  #### `service/`
  - **Purpose**: Contains business logic for the application.
  - **Files**: 
    - `note_service.go`: Implements business logic for managing notes.

  #### `router/`
  - **Purpose**: Sets up the HTTP routes for the application.
  - **Files**: 
    - `router.go`: Initializes the routes and links them with the handlers.

### `pkg/`
- **Purpose**: Contains shared packages that can be used across the application.
  
  #### `config/`
  - **Purpose**: Manages application configuration.
  - **Files**: 
    - `config.go`: Loads and parses application configuration settings.

  #### `middleware/`
  - **Purpose**: Implements middleware for the application.
  - **Files**: 
    - `auth.go`: Provides authentication and authorization middleware.

  #### `util/`
  - **Purpose**: Contains utility functions and helpers.
  - **Files**: 
    - `util.go`: Implements various helper functions used throughout the application.

### `go.mod`
- **Purpose**: Defines the module path and dependencies of the project.

### `go.sum`
- **Purpose**: Contains checksums for the dependencies to ensure integrity and reproducibility.