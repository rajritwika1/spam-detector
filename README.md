Spam Detector API

Project Overview
This is a Spam Detection API built using Golang (Gin Framework) and PostgreSQL. It provides:
- User Authentication (JWT)
- Spam Number Detection
- API Logging and Error Handling

1. Tech Stack
- Backend: Golang (Gin Framework)
- Database: PostgreSQL
- Auth: JWT-based authentication
- Environment Variables: `.env`

2. Setup Instructions

 Prerequisites
- Install Go (1.23.2)
- Install PostgreSQL(17)
- Install Golang dependencies: `go mod tidy`

Environment Variables
Create a `.env` file:

3. Run Database Migrations

go run main.go

4. Testing API with Postman

- Login API

URL: http://localhost:8080/api/auth/login

Method: POST

Body:

json

{
  "email": "test@example.com",
  "password": "password123"
}

- Spam Detection API

URL: http://localhost:8080/api/spam/check

Method: POST

Body:

json

{
  "phone": "+911234567890"
}
Headers:

json

{
  "Authorization": "Bearer your_generated_jwt_token"
}

Notes

- Ensure PostgreSQL is running before starting the server.

- Error logs are written using Logrus.

- Run Postman tests to verify API functionality.
