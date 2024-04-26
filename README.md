# Go User CRUD API

This is a simple Go REST API for performing CRUD operations on users. It uses the Gin framework for routing and PostgreSQL as the database.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Endpoints](#endpoints)
  - [Request Examples](#request-examples)
- [Testing](#testing)
- [License](#license)

## Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/GDS2005/go-rest.git
    cd go-rest
    ```

2. **Set up your PostgreSQL database:**

    - Create a PostgreSQL database.

3. **Build and run the application:**

    ```bash
    go mod tidy
    go build -o app
    ./app
    ```

    The application should now be running on `http://localhost:8080`.

## Usage

### Endpoints

- `GET /users`: Get all users
- `GET /users/{id}`: Get a user by ID
- `POST /users`: Create a new user
- `PATCH /users/{id}`: Update a user by ID
- `DELETE /users/{id}`: Delete a user by ID

### Request Examples

```bash
# Get all users
curl http://localhost:8080/users

# Get user by ID
curl http://localhost:8080/users/1

# Create a new user
curl -X POST http://localhost:8080/users -d '{"name": "John Doe", "email": "john@example.com"}'

# Update user by ID
curl -X PUT http://localhost:8080/users/1 -d '{"name": "Updated Name", "email": "updated@example.com"}'

# Delete user by ID
curl -X DELETE http://localhost:8080/users/1