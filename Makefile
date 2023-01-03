#!/usr/bin/env make

TARGET=./build

%:
	@rm -f ${TARGET}/$@ ; \
	if [[ -d ./cli/public/$@ ]];then \
		go run ./cli/_compile $@ ; \
	else \
		echo "Not found target project: $@"; \
	fi

public:
	go run ./cli/_compile public

clean:
	@go run ./cli/_compile -C

cleanr:
	@rm -rf release; \
	go clean ./... ; \
	echo "[+] Cleaned."

readme:
	@echo "# Virzz" > README.md ; \
	echo '' >> README.md; \
	echo '![Build](https://github.com/virzz/virzz/workflows/Build/badge.svg)' >> README.md; \
	echo '' >> README.md;

	@go run ./cli/_compile public;

	@echo "## Projects" >> README.md; \
	echo '' >> README.md; \
	for i in `ls ./cli/public/`; do \
		echo "- $$(basename $$i)" >> README.md; \
	done; \
	echo '' >> README.md;

	@for i in `ls ./cli/public/`; do \
		echo "## $$(basename $$i)" >> README.md; \
		echo '' >> README.md; \
		echo '```' >> README.md; \
		echo "$$(./build/$$i -h )" >> README.md; \
		echo '```' >> README.md; \
		echo '' >> README.md; \
	done;
	
	@cat README.md
