# Go Task Manager API

A RESTful API for managing tasks. Built with Go (Gin) + PostgreSQL (GORM).

---

## 📅 Project Roadmap

### Week 1 – Setup & Basic API
- Go module initialized
- Gin framework installed
- `/ping` endpoint → returns `{ "message": "pong" }`
- `/hello?name=YourName` endpoint → returns personalized greeting

### Week 2 – Database & Models
- Connected Go API to PostgreSQL using GORM
- Created `Task` model with fields:
  - ID
  - Title
  - Completed
  - DueDate
  - CreatedAt
  - UpdatedAt
- Migrated database schema automatically with GORM
- Tested database insert & read operations

---

## 📂 Project Structure



## Getting Started

### 1. Clone repository
```bash
git clone https://github.com/AyoobRukabi/go-task-manager.git
cd go-task-manager/cmd

