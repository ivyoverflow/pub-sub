module github.com/ivyoverflow/pub-sub/book

go 1.15

require (
	github.com/aws/aws-sdk-go v1.36.27 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/mock v1.4.4
	github.com/golang/snappy v0.0.2 // indirect
	github.com/google/uuid v1.1.5
	github.com/gorilla/mux v1.8.0
	github.com/ivyoverflow/pub-sub/platform v0.0.0-00010101000000-000000000000
	github.com/jmoiron/sqlx v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.8.0
	github.com/pkg/errors v0.9.1
	github.com/shopspring/decimal v1.2.0
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	golang.org/x/sys v0.0.0-20210113181707-4bcb84eeeb78 // indirect
	golang.org/x/text v0.3.5 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/ivyoverflow/pub-sub/platform => ../platform
