module github.com/ivyoverflow/pub-sub/server

go 1.15

require (
	github.com/golang/mock v1.4.4
	github.com/ivyoverflow/pub-sub/platform v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.6.1
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	golang.org/x/tools v0.0.0-20200818005847-188abfa75333 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect
)

replace github.com/ivyoverflow/pub-sub/platform => ../platform
