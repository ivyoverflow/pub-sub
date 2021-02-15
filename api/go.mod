module github.com/ivyoverflow/pub-sub/api

go 1.15

require (
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/golang/mock v1.4.4
	github.com/google/uuid v1.1.5
	github.com/gorilla/mux v1.8.0
	github.com/ivyoverflow/pub-sub/book v0.0.0-20210215112123-ce11ad458e09
	github.com/ivyoverflow/pub-sub/platform v0.0.0-00010101000000-000000000000
	github.com/jmoiron/sqlx v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.8.0
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.4.4
)

replace github.com/ivyoverflow/pub-sub/platform => ../platform
