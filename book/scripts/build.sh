# ENV for PostgreSQL
export PGHOST="localhost"
export PGPORT="5432"
export PGUSER="postgres"
export PGNAME="postgres"
export PGPASSWORD="qwerty"
export PGMIGRATIONSPATH="file://internal/store/postgres/migrations"
export PGSSLMODE="disable"

# ENV for API
export ADDR="localhost"
export PORT="8081"

cd book/ && \
go build -o build/book ./cmd/book/main.go && \
./build/book
