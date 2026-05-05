# Task Tracker CLI

A simple command-line interface (CLI) application to track and manage your tasks. This project helps you keep track of what you need to do, what you are currently working on, and what you have already completed. Built with Go, following clean architecture principles with separation of concerns across models, repository, usecase, and handler layers.

**Project Reference:** https://roadmap.sh/projects/task-tracker

## Features

- ✅ Add new tasks
- ✅ Update existing tasks
- ✅ Delete tasks
- ✅ Mark tasks as in-progress or done
- ✅ List all tasks
- ✅ List tasks by status (done, todo, in-progress)
- ✅ Task status management (todo, in-progress, done)
- ✅ Persistent storage using JSON
- ✅ Error handling and graceful edge case management
- ✅ No external dependencies (native Go only)

## Project Structure

```
Task-Tracker/
├── cmd/
│   └── main.go              # Application entry point
├── handler/
│   └── task_handler.go      # CLI request handlers
├── models/
│   └── task.go              # Data models and constants
├── repository/
│   └── task_repository.go   # Data persistence layer
├── usecase/
│   └── task_usecase.go      # Business logic layer
├── go.mod                   # Go module definition
├── tasks.json               # Task storage file (auto-created)
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

2. Build the application:

   ```bash
   go build -o task-cli cmd/main.go
   ```

## Running the Application

Run commands using the CLI:

```bash
./task-cli [command] [arguments]
```

## Available Commands

### Add a Task

Add a new task to your list:

```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Update a Task

Update the description of an existing task:

```bash
task-cli update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully
```

### Delete a Task

Remove a task from your list:

```bash
task-cli delete 1
# Output: Task deleted successfully
```

### Mark Task as In-Progress

Mark a task as currently being worked on:

```bash
task-cli mark-in-progress 1
# Output: Task marked as in-progress
```

### Mark Task as Done

Mark a task as completed:

```bash
task-cli mark-done 1
# Output: Task marked as done
```

### List All Tasks

Display all tasks in your list:

```bash
task-cli list
# Output: Lists all tasks with their details
```

### List Tasks by Status

Filter and display tasks by their status:

```bash
# List all completed tasks
task-cli list done

# List all tasks not yet started
task-cli list todo

# List all tasks currently in progress
task-cli list in-progress
```

## Task Status

Tasks can have the following statuses:

- **todo**: New task that hasn't been started
- **in-progress**: Task that is currently being worked on
- **done**: Completed task

## Data Storage

Tasks are stored in a `tasks.json` file in the current directory. The file is automatically created when you add your first task.

### Task Structure

Each task has the following properties:

```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo",
  "createdAt": "2026-05-05T10:30:00Z",
  "updatedAt": "2026-05-05T10:30:00Z"
}
```

**Properties:**

- **id**: A unique identifier for the task
- **description**: A short description of the task
- **status**: The current status (todo, in-progress, or done)
- **createdAt**: The date and time when the task was created
- **updatedAt**: The date and time when the task was last updated

## Architecture

This project follows **Clean Architecture** principles:

- **Models**: Define data structures and constants
- **Repository**: Handle data persistence (reading/writing to tasks.json)
- **Usecase**: Contain business logic and application rules
- **Handler**: Process CLI commands and route them to the appropriate usecase

## Development

To build the project into an executable:

```bash
go build -o task-cli cmd/main.go
```

To run directly without building:

```bash
go run cmd/main.go [command] [arguments]
```

### Example Workflow

```bash
# Add some tasks
./task-cli add "Buy groceries"
./task-cli add "Write documentation"
./task-cli add "Review code"

# Mark a task as in progress
./task-cli mark-in-progress 1

# View all tasks
./task-cli list

# Mark a task as done
./task-cli mark-done 1

# Update a task description
./task-cli update 2 "Write comprehensive documentation"

# List only completed tasks
./task-cli list done

# Delete a task
./task-cli delete 3
```

## Error Handling

The application gracefully handles errors including:

- Invalid task IDs
- Missing or invalid JSON file
- Malformed command arguments
- File system errors
- Task not found scenarios

## License

This project is provided as-is for educational and personal use.

## Getting Started Tips

1. Start by adding a few tasks
2. Experiment with listing tasks in different ways
3. Practice updating and marking tasks as complete
4. Check the `tasks.json` file to see how data is being stored

Happy task tracking!
