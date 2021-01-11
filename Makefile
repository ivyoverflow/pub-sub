.PHONY: server
server:
	sh server/scripts/build.sh

.PHONY: client
client:
	sh client/scripts/build.sh

clear:
	rm -rf server/build/ && \
	rm -rf client/build/
