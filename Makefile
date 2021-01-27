.PHONY: server
server:
	sh server/scripts/build.sh
.PHONY: listener
listener:
	sh listener/scripts/build.sh

.PHONY: test
test:
	sh server/scripts/test.sh && sh book/scripts/test.sh

.PHONY: linter
linter:
	sh server/scripts/linter.sh && sh server/scripts/linter.sh

.PHONY: book
book:
	sh book/scripts/build.sh

clear:
	rm -rf server/build/ && \
	rm -rf listener/build/ && \
	rm -rf book/build/
