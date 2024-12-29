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
```plaintext
monchef-be/
├── controllers/       # API handlers to manage HTTP requests and responses
├── db/                # Database connection and models
│   ├── db.go          # Database connection setup using GORM
│   └── models/        # Database schema definitions
│       ├── recipe.go  # Recipe model definition with fields mapped to the database
├── .env               # Environment variables (excluded from Git)
├── main.go            # Entry point of the application
├── go.mod             # Dependency management file
├── README.md          # Project documentation
└── .gitignore         # Git ignore rules
```
---

## Database Models
The db/models directory contains the schema definitions for the database tables used in the MonChef application. These models are mapped to the database tables using GORM.

### Recipe Model (db/models/recipe.go)
The Recipe model represents the recipes table in the MySQL database. Below is the structure of the model:
```go
type Recipe struct {
    ID          string    `gorm:"primaryKey" json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Country     string    `json:"country"`
    Categories  string    `json:"categories"` // Comma-separated values
    CreatedAt   time.Time `json:"created_at"`
    Duration    string    `json:"duration"`
    Difficulty  string    `json:"difficulty"`
    Images      string    `json:"images"`     // Comma-separated image URLs
    RecipeSteps string    `json:"recipe"`     // Recipe steps as a string
}
```

### Field Descriptions
- ID: The primary key for the table (gorm:"primaryKey")
- Name: The name of the recipe
- Description: A short description of the recipe
- Country: The country of origin for the recipe
- Categories: A comma-separated string for categorizing recipes (e.g., "Main Dish,Poultry")
- CreatedAt: The timestamp of when the recipe was created
- Duration: Estimated cooking time
- Difficulty: The level of difficulty (e.g., "Easy", "Intermediate")
- Images: Comma-separated URLs for images of the recipe
- RecipeSteps: A string containing the step-by-step instructions for the recipe.

### Example SQL Table Schema
The Recipe model maps to the following MySQL table schema:
```sql
CREATE TABLE recipes (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    country VARCHAR(255),
    categories TEXT,
    created_at DATETIME,
    duration VARCHAR(50),
    difficulty VARCHAR(50),
    images TEXT,
    recipe_steps TEXT
);
```


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