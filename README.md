# go-student-api

A simple RESTful API for managing student records, built with Go. This project demonstrates clean architecture, modular design, and best practices for building web APIs in Go.

## Features
- Create, retrieve, and list student records
- Input validation using `go-playground/validator`
- SQLite storage backend
- Modular internal structure for easy maintenance and scalability

## Project Structure
```
student-api/
├── cmd/                # Application entrypoint
│   └── student-api/
│       └── main.go
├── internal/           # Internal application code
│   ├── config/         # Configuration management
│   ├── http/           # HTTP handlers
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go
│   ├── storage/        # Storage abstraction and SQLite implementation
│   ├── types/          # Type definitions
│   └── utils/          # Utility functions
├── go.mod
├── go.sum
└── README.md           # Project documentation
```

## Getting Started

### Prerequisites
- Go 1.20 or later

### Installation
1. Clone the repository:
```sh
git clone https://github.com/aminShahid573/go-student-api.git
cd go-student-api
```
2. Download dependencies:
```sh
go mod download
```

### Running the API
1. Build and run the application:
```sh
go run ./cmd/student-api/main.go
```
2. The API will start **http://localhost:`{{your_port in_config}}`**`

## API Endpoints

### Create Student
- **POST** `/students`
- **Body:**
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "age": 21
  }
  ```
- **Response:**
  ```json
  { "id": 1 }
  ```

### Get Student by ID
- **GET** `/students/{id}`
- **Response:**
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 21
  }
  ```

### List All Students
- **GET** `/students`
- **Response:**
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "age": 21
    },
    ...
  ]
  ```

## Contributing
Contributions are welcome! Please open issues or submit pull requests for improvements and bug fixes.

## License
This project is licensed under the MIT License. 