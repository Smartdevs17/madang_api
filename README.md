# Madang API

Madang API is a backend application built using Golang and the Gin framework. It provides API endpoints to manage users, food items, tables, and other restaurant-related functionalities. This README serves as a guide for understanding, running, and contributing to the project.

## Features

- **User Management**: Handle user authentication and authorization.
- **Restaurant Management**: Manage restaurant food items, tables, and services.
- **Middleware**: Secure routes with authentication middleware.
- **Utilities**: Helper functions for structured responses.

---

## Installation

### Prerequisites

- Go 1.18 or higher
- A PostgreSQL database
- Git

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/madang_api.git
   cd madang_api
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up environment variables:

   Create a `.env` file in the root directory with the following variables:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=yourusername
   DB_PASSWORD=yourpassword
   DB_NAME=madang
   JWT_SECRET=yourjwtsecret
   ```

4. Run database migrations (if applicable):

   ```bash
   go run migrations/migrate.go
   ```

5. Start the application:

   ```bash
   go run main.go
   ```

---

## Usage

### Endpoints

#### Authentication

- **POST** `/api/auth/login`: Authenticate a user and return a JWT token.

#### Users

- **GET** `/api/users/me`: Retrieve the details of the logged-in user.

#### Inits

- **GET** `/api/inits/`: Retrieve user, food, and table data.

### Example Request

#### Retrieve User, Food, and Table Data

```bash
curl -X GET \  
-H "Authorization: Bearer <your-token>" \  
http://localhost:8080/api/inits/
```

---

## Project Structure

```plaintext
madang_api/
├── controllers/   # API endpoint handlers
├── middleware/    # Middleware functions
├── models/        # Data models
├── routes/        # Route definitions
├── services/      # Business logic
├── utils/         # Helper utilities
├── migrations/    # Database migrations
├── main.go        # Application entry point
```

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a feature branch:

   ```bash
   git checkout -b feature/new-feature
   ```

3. Commit your changes:

   ```bash
   git commit -m "Add new feature"
   ```

4. Push to the branch:

   ```bash
   git push origin feature/new-feature
   ```

5. Open a pull request.

---

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---

## Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)

