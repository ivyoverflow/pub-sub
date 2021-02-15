![CircleCI Status](https://img.shields.io/circleci/build/github/ivyoverflow/pub-sub/main?style=flat-square)
![GitHub Issues](https://img.shields.io/github/issues/ivyoverflow/pub-sub?style=flat-square)
![GitHub Pull Requests](https://img.shields.io/github/issues-pr/ivyoverflow/pub-sub?style=flat-square)
![GitHub Last Commit](https://img.shields.io/github/last-commit/ivyoverflow/pub-sub?style=flat-square)
# 🦹🏻‍♀️ pub-sub
## 🧪 What we are using?
1. 🍃 MongoDB.
2. 🐘 PostgreSQL.
3. 🐹 Golang.
## 📌 How to run tests?
### 🐳 With Docker:
🍨 Run docker-compose with the following command:
```bash
docker-compose up -d
```
🧁  Run tests with the following commands:<br>
To initialize PostgreSQL migrations:
```bash
migrate -path ./book/internal/storage/postgres/migrations -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up
```
To run tests:
```bash
make test
```
### 🪐 With CircleCI:
🍰  Run the following command:
```bash
circleci local execute --job build
```
## 📌 How to run services?
>💡 WARNING: before trying to run anything, you must have the following environment variables:
```bash
# MongoDB environment variables.
export MONGOHOST="<YOUR HOST>"
export MONGOPORT="<YOUR PORT>"
export MONGONAME="<YOUR DATABASE NAME>"
export MONGOUSER="<YOUR USERNAME>"
export MONGOPASSWORD="<YOUR PASSWORD>"
# PostgreSQL environment variables.
export PGHOST="<YOUR HOST>"
export PGPORT="<YOUR PORT>"
export PGUSER="<YOUR USERNAME>"
export PGNAME="<YOUR DATABASE NAME>"
export PGPASSWORD="<YOUR PASSWORD>"
export PGSSLMODE="<YOUR SSL MODE>"
# server environment variables.
export ADDR="<YOUR HOST>"
export PORT="<YOUR PORT>"
```
>💡 WARNING: you also need to initialize PostgreSQL migrations:
```bash
migrate -path ./book/internal/storage/postgres/migrations -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up
```
## 🚀 Contributors
[👨🏻‍🎓 ivyoverflow](https://github.com/ivyoverflow) &&  [👨🏻‍🚀 kiryalovik](https://github.com/kiryalovik)
