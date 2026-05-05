# Task Tracker

A simple and efficient task management web application built with Go. This project follows clean architecture principles with separation of concerns across models, repository, usecase, and handler layers.

## Features

- ✅ Create new tasks
- ✅ List all tasks with optional status filtering
- ✅ Delete tasks by ID
- ✅ Task status management (todo, in-progress, done)
- ✅ Persistent storage using JSON
- ✅ RESTful API endpoints
- ✅ Lightweight and fast

## Project Structure

```
Task-Tracker/
├── cmd/
│   └── main.go              # Application entry point
├── handler/
│   └── task_handler.go      # HTTP request handlers
├── models/
│   └── task.go              # Data models and constants
├── repository/
│   └── task_repository.go   # Data persistence layer
├── usecase/
│   └── task_usecase.go      # Business logic layer
├── go.mod                   # Go module definition
├── tasks.json               # Task storage file
└── README.md                # This file
```

## Prerequisites

- Go 1.25.0 or higher
- Basic command line knowledge

## Installation

1. Clone or download the project:

   ```bash
   cd Task-Tracker
   ```

2. Download dependencies:
   ```bash
   go mod download
   ```

## Running the Application

Start the web server:

```bash
go run cmd/main.go
```

The server will start at `http://localhost:8080`

## API Endpoints

### List Tasks

- **Endpoint:** `GET /list`
- **Description:** Retrieve all tasks or filter by status
- **Query Parameters:**
  - `status` (optional): Filter by status - `todo`, `in-progress`, or `done`
- **Example:**

  ```bash
  # Get all tasks
  curl http://localhost:8080/list

  # Get only completed tasks
  curl "http://localhost:8080/list?status=done"
  ```

### Add Task

- **Endpoint:** `POST /add`
- **Description:** Create a new task
- **Query Parameters:**
  - `desc` (required): Task description
- **Example:**
  ```bash
  curl "http://localhost:8080/add?desc=Buy%20groceries"
  ```

### Delete Task

- **Endpoint:** `DELETE /delete`
- **Description:** Delete a task by ID
- **Query Parameters:**
  - `id` (required): Task ID to delete
- **Example:**
  ```bash
  curl -X DELETE "http://localhost:8080/delete?id=1"
  ```

## Task Status

Tasks can have the following statuses:

- **todo**: New task that hasn't been started
- **in-progress**: Task that is currently being worked on
- **done**: Completed task

## Data Storage

Tasks are stored in `tasks.json` file in the project root. The file is automatically created and managed by the repository layer.

### Task Structure

```json
{
  "id": 1,
  "description": "Sample task",
  "status": "todo",
  "createdAt": "2026-05-05T10:30:00Z",
  "updatedAt": "2026-05-05T10:30:00Z"
}
```

## Architecture

This project follows **Clean Architecture** principles:

- **Models**: Define data structures and constants
- **Repository**: Handle data persistence (reading/writing to tasks.json)
- **Usecase**: Contain business logic and application rules
- **Handler**: Process HTTP requests and responses

## Development

To build the project:

```bash
go build -o task-tracker cmd/main.go
```

Then run the compiled binary:

```bash
./task-tracker
```

## License

This project is provided as-is for educational and personal use.

## Support

For issues or questions, please review the code structure and ensure all dependencies are properly installed.
