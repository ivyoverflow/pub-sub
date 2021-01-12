.PHONY: server
server:
	sh server/scripts/build.sh

.PHONY: publisher
publisher:
	sh publisher/scripts/build.sh

.PHONY: listener
listener:
	sh listener/scripts/build.sh

clear:
	rm -rf server/build/ && \
	rm -rf publisher/build/ && \
	rm -rf listener/build/
