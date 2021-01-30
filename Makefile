.PHONY:
.SILENT:

test:
	sh book/scripts/test.sh && \
		sh server/scripts/test.sh
