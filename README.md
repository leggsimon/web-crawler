## Webpage crawler

This is a small web page crawler which crawls and analyses a web page, returning some information about it.

---

### Prerequisites

- **Go** 1.21+ (only needed if you want to run the backend outside Docker)
- **Docker** and **Docker Compose** (Docker Desktop on macOS/Windows, or Docker Engine + compose plugin on Linux)

---

### Configuration

Most configuration is provided via environment variables.

**MySQL (docker-compose)**

- `MYSQL_ROOT_PASSWORD` — root password for the MySQL container.
- `MYSQL_DATABASE` — database name (default: `app`).
- `MYSQL_USER` — application DB user (default: `app`).
- `MYSQL_PASSWORD` — password for `MYSQL_USER`.

These are read by the `mysql` service in `docker-compose.yml` and are also used by the backend to construct its DSN.

**Backend**

You can configure the backend by setting the following environment varriables in a `.env` file

- `MYSQL_DSN` — full MySQL DSN (optional). It takes priority over individual variables below.
- `MYSQL_USER` — DB user (default: `root` when running outside Docker, `app` inside Docker-compose).
- `MYSQL_PASSWORD` — DB password.
- `MYSQL_HOST` — DB host (default: `localhost` outside Docker, `mysql` in Compose).
- `MYSQL_PORT` — DB port (default: `3306`).
- `MYSQL_DATABASE` — DB name (default: `app`).
- `PORT` — HTTP port the API listens on (default: `8080`).

---

### Running with Docker

From the project root:

```bash
docker compose up --build
```

To stop everything:

```bash
docker compose down
```

This keeps the MySQL data volume (`mysql_data`) by default. To remove volumes as well:

```bash
docker compose down -v
```

---

### Running the backend without Docker

1. **Start MySQL with Docker Compose:**

   ```bash
   docker compose up -d mysql
   ```

2. **Export environment variables so the backend points at the container:**

   ```bash
   export MYSQL_HOST=localhost
   export MYSQL_PORT=3306
   export MYSQL_DATABASE=app
   export MYSQL_USER=app
   export MYSQL_PASSWORD=app
   export PORT=8080
   ```

3. **Run the backend from the `backend/` directory:**

   ```bash
   cd backend
   go run ./...
   ```

### Running the backend with Docker:

```bash
docker compose up --build backend
```

You have to rebuild and restart the backend to pick up code changes and migrations
