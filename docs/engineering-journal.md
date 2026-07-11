# Engineering Journal

## Day 1

Started the Cab Share Backend project.

## Milestone: Database Setup

### Completed
- Configured PostgreSQL using Docker Compose.
- Connected the Go application to PostgreSQL using GORM.
- Learned the purpose of database drivers and ORM.
- Created the first `User` model.
- Used `AutoMigrate` to generate the `users` table.
- Verified the generated schema in PostgreSQL.
- Inserted and retrieved records using GORM.
- Explored `Create`, `First`, `Find`, and `Where`.

### Key Learnings
- Difference between `First` and `Find`.
- GORM infers table names from model types.
- Method chaining in GORM.
- Importance of checking returned errors.
- Passwords should be hashed, not encrypted.
- Docker port mapping (`HOST_PORT:CONTAINER_PORT`).