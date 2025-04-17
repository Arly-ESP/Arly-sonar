*# API Documentation

## Overview
This API provides a robust foundation for managing users, activities, and administrative operations. Built with Go, Fiber, and PostgreSQL, it incorporates JWT-based authentication for securing routes and role-based access control for user and admin-specific operations.

---

## Features

### Authentication
- **JWT Authentication**: Secure access to protected routes with Bearer tokens.
- **Role-Based Access Control**: Admin-specific operations are restricted to `user_id = 1`.

### User Management
- **User Registration**: Create new users with email and password validation.
- **User Profile**:
  - Fetch and update authenticated user details.
  - Passwords are securely hashed using bcrypt.

### Activity Tracking
- Log and retrieve user mood and activity data.
- Fetch historical activity logs with optional date range filtering.

### Administrative Features
- Fetch all users.
- Manage user data (fetch by ID and delete).

---

## Environment Configuration
The following environment variables must be configured:

```env
# Database Configuration
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=
DB_SSLMODE=
DB_TIMEZONE=

# Server Configuration
PORT=5050
SERVER_URL=0.0.0.0 

# JWT Configuration
JWT_SECRET=
JWT_EXPIRE=24h
JWT_REFRESH_SECRET=
JWT_REFRESH_EXPIRE=72h

# Admin User
ADMIN_FIRSTNAME=
ADMIN_LASTNAME=
ADMIN_EMAIL=
ADMIN_PASSWORD=

OPEN_AI_API_KEY= 
```

---

## Getting Started

### Prerequisites
- Go >= 1.23.4
- PostgreSQL

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Zesor/Arly
   cd Arly
   ```
2. Set up environment variables:
   - copy the correponding `.env.*` file from `backend/environments` to the root folder and rename it to `.env`.
   - Update `.env` with the appropriate configuration.
3. Run the application:
   ```bash
   make run
   ```

### Database Initialization
- Migrations are automatically executed during startup.
- If the database is reinitialized, an admin user (`user_id = 1`) is automatically created using the environment variables provided.
- The admin user's credentials are used to access admin-specific routes and operations.
- The admin user's credentials can be updated in the `.env` file.

### Launch with Docker Compose
1. Ensure Docker and Docker Compose are installed on your system.
2. Build and run the application:
   ```bash
   docker-compose up -d
   ```
3. Check the logs to ensure the application is running:
   ```bash
   docker-compose logs -f
   ```

---

## API Endpoints

### Public Routes
#### User Registration
- **POST** `/api/users`
  - Register a new user.

### Protected Routes
#### User Operations
- **GET** `/api/user`
  - Fetch authenticated user's details.
- **PUT** `/api/user`
  - Update authenticated user's details.

#### Activity Operations
- **POST** `/api/mood`
  - Log user's mood.
- **GET** `/api/user/activity`
  - Fetch today's activity.
- **GET** `/api/user/activities`
  - Fetch historical activities (optionally filtered by date range).

### Admin Routes
#### Admin Operations
- **GET** `/api/admin/users`
  - Fetch all users.
- **GET** `/api/admin/users/:id`
  - Fetch user details by ID.
- **DELETE** `/api/admin/users/:id`
  - Delete a user by ID.

---

## Middleware
### JWTMiddleware
Protects routes by:
- Validating Bearer tokens.
- Extracting `userID` from JWT claims.

### AdminJWTMiddleware
Restricts routes to admin users (`user_id = 1`) by:
- Validating Bearer tokens.
- Verifying the `userID` claim matches `1`.

---

## Project Structure
```plaintext
.
├── controllers   # Route handlers
├── commands      # CLI commands
├── docs          # API documentation
├── enums         # Enumerations
├── index         # Application entry point
├── middleware    # JWT and Admin middleware
├── services      # Custom services (e.g., JWT, database)
├── templates     # HTML templates ( e.g., email templates)
├── tests         # Unit tests
├── types         # Custom types
├── models        # Database models
├── database      # Database connection and migrations
├── routes        # API route definitions
├── utilities     # Utility functions (e.g., logging, password hashing)
└── config        # Environment and configuration management

```

---

## License
This project is licensed under the MIT License. See `LICENSE` for more information.

---

## Testing

### Technical Details
A new folder has been added in the root folder: ./tests. This folder is used as some kind of "helper" for tests.
The tests are written in the same way as the previous ones, but now we have a new folder that contains the tests for the API.
- Unit tests are written using the Go testing framework.
- The `testify` library is used for assertions.
- The `pgmock` library is used to mock PostgreSQL database operations.
- The `http` package is used to mock HTTP requests.
- The `httptest` package is used to test HTTP handlers.
- The `fiber` package is used to test API routes.

### How to Run Unit Tests
From the route project:
```bash
  go test ./... -cover
```

## Contact

*