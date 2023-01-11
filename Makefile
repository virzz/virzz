#!/usr/bin/env make

TARGET=./build

%:
	@rm -f ${TARGET}/$@ ; \
	if [[ -d ./cli/public/$@ || -d ./cli/$@ ]];then \
		if [[ -z $DEBUG ]]; then \
			go run -tags debug ./cli/_compile $@ ; \
		else \
			go run ./cli/_compile $@ ; \
		fi; \
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
	@echo "Add Title"; \
	echo "# Virzz" > README.md ; \
	echo '' >> README.md; \
	echo "Add Build Badge"; \
	echo '![Build](https://github.com/virzz/virzz/workflows/Build/badge.svg)' >> README.md; \
	echo '' >> README.md;

	@go run ./cli/_compile god;
	@echo "Add God"; \
	echo '## God - CLI 命令行小工具' >> README.md; \
	echo '' >> README.md; \
	echo '```' >> README.md; \
	./build/god >> README.md; \
	echo '```' >> README.md; \
	echo '' >> README.md; \

	@echo "Compile Public Projects"
	@go run ./cli/_compile public;

	@echo "Add Public Project List"; \
	echo "## Public Projects" >> README.md; \
	echo '' >> README.md; \
	for i in `ls ./cli/public/`; do \
		echo "- $$(basename $$i)" >> README.md; \
	done; \
	echo '' >> README.md;

	@echo "Add Public Projects"; \
	for i in `ls ./cli/public/`; do \
		echo "## $$(basename $$i)" >> README.md; \
		echo '' >> README.md; \
		echo '```' >> README.md; \
		echo "$$(./build/$$i -h )" >> README.md; \
		echo '```' >> README.md; \
		echo '' >> README.md; \
	done;
