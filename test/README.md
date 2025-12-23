# Healthcare Test Server

This is a test backend for the shapeist middleware library. It provides an in-memory fake healthcare server with CRUD operations for patient entities.

## Features

- **In-memory storage**: All data is stored in memory (no database required)
- **Random data generation**: Server starts with 5 randomly generated patients
- **Full CRUD operations**: Create, Read, Update, and Delete patients
- **RESTful API**: Standard REST endpoints with JSON responses

## Patient Entity

Each patient has the following fields:
- `id` (int): Unique identifier
- `first_name` (string): Patient's first name
- `last_name` (string): Patient's last name
- `birth_date` (time): Patient's date of birth
- `gender` (string): Patient's gender (Male, Female, Other)
- `email` (string): Patient's email address
- `phone` (string): Patient's phone number

## Running the Server

```bash
cd test
go build
./test
```

The server will start on `http://localhost:8080`.

## API Endpoints

### List All Patients
```bash
GET /patients
```

**Example:**
```bash
curl http://localhost:8080/patients
```

### Get Patient by ID
```bash
GET /patients/{id}
```

**Example:**
```bash
curl http://localhost:8080/patients/1
```

### Create Patient
```bash
POST /patients
Content-Type: application/json
```

**Example:**
```bash
curl -X POST http://localhost:8080/patients \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "birth_date": "1990-01-01T00:00:00Z",
    "gender": "Male",
    "email": "john.doe@example.com",
    "phone": "555-1234"
  }'
```

### Update Patient
```bash
PUT /patients/{id}
Content-Type: application/json
```

**Example:**
```bash
curl -X PUT http://localhost:8080/patients/1 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Jane",
    "last_name": "Doe",
    "birth_date": "1990-01-01T00:00:00Z",
    "gender": "Female",
    "email": "jane.doe@example.com",
    "phone": "555-5678"
  }'
```

### Delete Patient
```bash
DELETE /patients/{id}
```

**Example:**
```bash
curl -X DELETE http://localhost:8080/patients/1
```

## Module Information

This test backend has its own `go.mod` file (separate from the main shapeist module) to isolate test dependencies.
