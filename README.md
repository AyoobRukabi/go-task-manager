# Go Task Manager API

A RESTful API for managing tasks. Built with **Go (Gin) + PostgreSQL + GORM**.

---

## Features

* **User-friendly API endpoints** for tasks
* Full **CRUD operations** (Week 3)
* **PostgreSQL database** with GORM ORM
* Clean, scalable project structure

---

## Tech Stack

* **Go** (Gin framework)
* **GORM** (ORM for Go)
* **PostgreSQL**
* **Docker** (planned for Week 4)

---

## Project Structure

```
go-task-manager/
├─ cmd/
│  └─ main.go         # server entry point
├─ db/
│  └─ db.go           # database connection
├─ handlers/
│  ├─ general.go      # /ping and /hello handlers
│  └─ task.go         # CRUD task handlers
├─ models/
│  └─ task.go         # Task struct
├─ routes/
│  └─ task.go         # task routes registration
├─ go.mod
└─ README.md
```

---

## Getting Started

1. **Clone the repo**

```bash
git clone https://github.com/AyoobRukabi/go-task-manager.git
cd go-task-manager
```

2. **Run the server**

```bash
go run ./cmd/main.go
```

Server runs at `http://localhost:8080`.

---

## Week 1 – Basic Endpoints

* **GET /ping** → returns:

```json
{"message":"pong"}
```

* **GET /hello?name=YourName** → returns:

```json
{"message":"Hello YourName"}
```

---

## Week 2 – Database & Models

* PostgreSQL connected with GORM
* `Task` model created:

```go
type Task struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Title     string    `json:"title" binding:"required"`
    Completed bool      `json:"completed"`
    DueDate   *time.Time `json:"due_date,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

* Database auto-migrates the `tasks` table on startup

---

## Week 3 – CRUD Endpoints

* Full REST API for tasks:

| Method | Endpoint    | Description           |
| ------ | ----------- | --------------------- |
| POST   | /tasks      | Create a task         |
| GET    | /tasks      | List all tasks        |
| GET    | /tasks/\:id | Get single task by ID |
| PUT    | /tasks/\:id | Update a task         |
| DELETE | /tasks/\:id | Delete a task         |

* **Request validation:**

  * `title` is required
  * `completed` defaults to false
  * `due_date` is optional

---

### Example curl commands

**Create a task:**

```bash
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title":"Learn Go","completed":false}'
```

**Get all tasks:**

```bash
curl http://localhost:8080/tasks
```

**Get task by ID:**

```bash
curl http://localhost:8080/tasks/1
```

**Update a task:**

```bash
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{"title":"Master Go","completed":true}'
```

**Delete a task:**

```bash
curl -X DELETE http://localhost:8080/tasks/1
```

---

## Next Steps (Week 4)

* Add **JWT authentication**
* Write **unit tests**
* Add **Swagger API documentation**
* Dockerize the app and deploy




---

### Week 4 (JWT Authentication)

## Features (Week 4)

* JWT-based **user authentication**
* Register and login users
* Protect task routes with **Authorization: Bearer token**
* Ready for Docker, Swagger, and deployment

---

## User Routes

| Method | Endpoint  | Description             |
| ------ | --------- | ----------------------- |
| POST   | /register | Register a new user     |
| POST   | /login    | Login and get JWT token |

---

### Example curl commands

**Register a user:**

```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"username":"Ayoob","password":"mypassword"}'
```

**Login and get token:**

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{"username":"Ayoob","password":"mypassword"}'
```

Response:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

## Task Routes (JWT Protected)

| Method | Endpoint    | Description           |
| ------ | ----------- | --------------------- |
| POST   | /tasks      | Create a task         |
| GET    | /tasks      | List all tasks        |
| GET    | /tasks/\:id | Get single task by ID |
| PUT    | /tasks/\:id | Update a task         |
| DELETE | /tasks/\:id | Delete a task         |

> **All task endpoints require the JWT token in the header**:

```
Authorization: Bearer <token>
```

**Create a task (example):**

```bash
curl -X POST http://localhost:8080/tasks/ \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_token_here>" \
-d '{"title":"Learn Go","completed":false}'
```

**Get all tasks:**

```bash
curl -X GET http://localhost:8080/tasks/ \
-H "Authorization: Bearer <your_token_here>"
```

**Update a task:**

```bash
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_token_here>" \
-d '{"title":"Master Go","completed":true}'
```

**Delete a task:**

```bash
curl -X DELETE http://localhost:8080/tasks/1 \
-H "Authorization: Bearer <your_token_here>"
```

---

## Environment Variable

Set the JWT secret **before running the app**:

```bash
export JWT_SECRET="your_strong_secret_here"
go run ./cmd/main.go
```

---

✅ This README now fully documents **Week 4**.

Next step: **commit Week 4 changes + updated README** and push to GitHub.

Here’s the **git command block** you can use:

```bash
git add .
git commit -m "Week 4: Add JWT authentication and update README"
git push origin main
```

---

If you want, I can also **write a final checklist for Week 4** to make sure your repo is completely CV-ready and professional.

Do you want me to do that?

