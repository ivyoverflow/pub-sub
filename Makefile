.PHONY:
.SILENT:
.DEFAULT_GOAL := build

build:
	sh book/scripts/build.sh && \
		sh server/scripts/build.sh

lint:
	sh book/scripts/lint.sh && \
		sh server/scripts/lint.sh

tidy:
	sh book/scripts/tidy.sh && \
		sh server/scripts/tidy.sh

docker-up:
	docker-compose up -d

docker-down:
	docker-compose stop && \
		docker-compose rm

test: docker-up
	sh book/scripts/test.sh && \
		sh server/scripts/test.sh

clear:
	sh book/scripts/clear.sh && \
		sh server/scripts/clear.sh
