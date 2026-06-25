# What does this app do?

Log meals and exercises\
View them by date

# Tech Stack

Backend: Go (Golang)\
Frontend (Optional): React or simple CLI\
Storage: In-memory (maps/slices), then possibly SQLite or Postgres\
Bonus: Docker for containerization

# Features / What is learned
Log food/exercise	API design, structs, POST/GET
View by date	Maps, date handling
Edit/delete logs	Linked lists, arrays
Summary dashboard	Trees or hash maps (for aggregation)
Search logs	Search algorithms, time complexity
Backend storage	Practice structs, maps, pointers
Optionally: DB integration	ORM pain points, N+1 queries, indexing



# Data Structures

| Type              | Used For                          |
| ----------------- | --------------------------------- |
| Arrays            | Recent logs, pagination           |
| Linked Lists      | Undo/redo logs, history tracking  |
| Hash Maps         | Quick access to logs by date/user |
| Trees             | Summarize calories by day/week    |
| Queues            | Track upcoming workout plans      |
| Graphs (optional) | Exercise dependencies or routines |

# API Endpoints
POST   /api/food         -> Log a food entry\
POST   /api/exercise     -> Log an exercise\
GET    /api/logs?date=2025-10-07\
PUT    /api/food/:id     -> Update food\
DELETE /api/exercise/:id
