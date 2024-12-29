# MonChef Backend

MonChef is a platform where community members can share recipes, restaurants, and food photos. This repository contains the backend server for the MonChef project.

## Key Features
- RESTful APIs to upload and retrieve recipe information
- Integration with a MySQL database
- Scalable and maintainable architecture

---

## Tech Stack
- **Programming Language**: Go (Golang)
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: MySQL
- **ORM**: GORM
- **Environment Variable Management**: `.env` file

---

## Project Structure
```
monchef-be/
├── controllers/       # API handlers
├── db/                # Database connection and migrations
├── models/            # Database models
├── .env               # Environment variables (excluded from Git)
├── main.go            # Entry point of the application
├── go.mod             # Dependency management file
├── README.md          # Project documentation
└── .gitignore         # Git ignore rules
```

---

## Installation and Setup

### Prerequisites
- [Go 1.22+](https://golang.org/dl/)
- [MySQL](https://www.mysql.com/)

### Step 1: Configure Environment Variables
Create a `.env` file in the root directory of `monchef-be` and add the following content:
```
DB_USER=monchef_user DB_PASSWORD=Tempmon!12 DB_HOST=127.0.0.1 DB_PORT=3306 DB_NAME=monchef
```


### Step 2: Install Dependencies
Run the following command to install dependencies:
```bash
go mod tidy
```

### Step 3: Prepare the MySQL Database
Set up the monchef database and a user with the necessary permissions:
```sql
CREATE DATABASE monchef;
CREATE USER 'monchef_user'@'%' IDENTIFIED BY 'Tempmon!12';
GRANT ALL PRIVILEGES ON monchef.* TO 'monchef_user'@'%';
FLUSH PRIVILEGES;
```

### Step 4: Start the Server
Run the server using the following command:
```bash
go run main.go
```