# AGENTS.md

## Cursor Cloud specific instructions

This is a Go REST API for user management (CRUD) using Gorilla Mux and MySQL.

### Services

| Service | Purpose | Port |
|---------|---------|------|
| Go app (`go run main.go`) | REST API server | 3000 |
| MySQL | Database backend | 3306 |

### Starting MySQL

MySQL must be running before the app starts. Use:

```bash
sudo mysqld_safe --user=mysql &
sleep 3
```

If MySQL was freshly installed, you may need to initialize it first:

```bash
sudo mysqld --initialize-insecure --user=mysql
sudo mkdir -p /var/run/mysqld && sudo chown mysql:mysql /var/run/mysqld
```

### Database setup

The database `dev-book` must exist with the schema from `sql/sql.sql` applied. The `.env` file at the project root configures the connection (user: `admin`, password: `password`, port: `3306`).

### Gotchas

- The `.env.example` lists `DB_PORT=5432` (PostgreSQL default), but the app uses MySQL — always use port `3306`.
- The repo has no test files (`*_test.go`), so `go test ./...` reports no tests but exits cleanly.
- The repo has no linter config; use `go vet ./...` for basic static analysis.
- Build with `go build -o main_dev` to avoid overwriting the checked-in `main` binary used for Heroku deployment.

### Standard commands

- **Build:** `go build -o main_dev`
- **Run:** `go run main.go` (reads `.env` automatically in dev mode)
- **Lint:** `go vet ./...`
- **Test:** `go test ./...`
- **Deps:** `go mod download`
