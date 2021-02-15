.PHONY:
.SILENT:
.DEFAULT_GOAL := build

build:
	sh api/scripts/build.sh && \
		sh notifier/scripts/build.sh

lint:
	sh api/scripts/lint.sh && \
		sh notifier/scripts/lint.sh

tidy:
	sh api/scripts/tidy.sh && \
		sh notifier/scripts/tidy.sh

docker-up:
	docker-compose up -d

docker-down:
	docker-compose stop && \
		docker-compose rm

test:
	sh api/scripts/test.sh && \
		sh notifier/scripts/test.sh

clear:
	sh api/scripts/clear.sh && \
		sh notifier/scripts/clear.sh
