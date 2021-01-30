module github.com/ivyoverflow/pub-sub/listener

go 1.15

require (
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
)

replace github.com/ivyoverflow/pub-sub/platform => ../platform
