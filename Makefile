#!/usr/bin/env make

TARGET=./build

default:
	@if [[ -z "${DEBUG}" ]]; then \
		go run -tags debug ./internal/_compile god ; \
	else \
		go run ./internal/_compile god ; \
	fi;

%:
	@rm -f ${TARGET}/$@ ; 
	@if [[ -d ./cmd/public/$@ ]] || [[ -d ./cmd/$@ ]]; then \
		if [[ -z "${DEBUG}" ]]; then \
			go run -tags debug ./internal/_compile $@ ; \
		else \
			go run ./internal/_compile $@ ; \
		fi; \
	fi;

rc-%:
	@echo "[*] Compiling Release $(subst rc-,,$@) ..." ; \
	go run ./internal/_compile -R $(subst rc-,,$@)

i-%: rc-%
	@export NAME='$(subst i-,,$@)'; \
	echo "[*] Install Release $${NAME} ..." ; \
	cp -f ${TARGET}/$${NAME} ${GOPATH}/bin/$${NAME}; \
	test -f ${GOPATH}/bin/$${NAME} && echo "[+] $${NAME} Installed";

clean:
	@go run ./internal/_compile -C

cleanr:
	@rm -rf release; \
	go clean ./... ; \
	echo "[+] Cleaned."

swagger:
	@swag i -g services/web/swagger.go -o services/docs

readme:
	@echo "Add Title"; \
	echo "# Virzz" > README.md ; \
	echo '' >> README.md; \
	echo "Add Build Badge"; \
	echo '[![Build](https://github.com/virzz/virzz/actions/workflows/virzz.yml/badge.svg)](https://github.com/virzz/virzz/actions/workflows/virzz.yml) [![Build Release](https://github.com/virzz/virzz/actions/workflows/virzz_release.yml/badge.svg)](https://github.com/virzz/virzz/actions/workflows/virzz_release.yml)' >> README.md; \
	echo '' >> README.md;

	@go run ./cmd/_compile god;
	@echo "Add God"; \
	echo '## God - CLI 命令行小工具' >> README.md; \
	echo '' >> README.md; \
	echo '```' >> README.md; \
	./build/god >> README.md; \
	echo '```' >> README.md; \
	echo '' >> README.md; \

	@echo "Compile Public Projects"
	@go run ./cmd/_compile public;

	@echo "Add Public Project List"; \
	echo "## Public Projects" >> README.md; \
	echo '' >> README.md; \
	for i in `ls ./cmd/public/`; do \
		echo "- $$(basename $$i)" >> README.md; \
	done; \
	echo '' >> README.md;

	@echo "Add Public Projects"; \
	for i in `ls ./cmd/public/`; do \
		echo "## $$(basename $$i)" >> README.md; \
		echo '' >> README.md; \
		echo '```' >> README.md; \
		echo "$$(./build/$$i -h )" >> README.md; \
		echo '```' >> README.md; \
		echo '' >> README.md; \
	done;
