#!/usr/bin/env make

TARGET=./build

default:
	@if [[ -z "${DEBUG}" ]]; then \
		go run -tags debug ./internal/_compile enyo ; \
	else \
		go run ./internal/_compile enyo ; \
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
	go run ./internal/_compile -R -G $(shell git rev-parse HEAD || echo "latest" ) $(subst rc-,,$@)

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

readme: enyo
	@echo "Add Title"; \
	echo "# Virzz" > README.md ; \
	echo '' >> README.md; \
	echo "Add Build Badge"; \
	echo '[![Build](https://github.com/virzz/virzz/actions/workflows/virzz.yml/badge.svg)](https://github.com/virzz/virzz/actions/workflows/virzz.yml) [![Build Release](https://github.com/virzz/virzz/actions/workflows/virzz_release.yml/badge.svg)](https://github.com/virzz/virzz/actions/workflows/virzz_release.yml)' >> README.md; \
	echo '' >> README.md; \
	echo "## Install" >> README.md; \
	echo '' >> README.md; \
	echo '`brew install virzz/virzz/<formula>` || `brew tap virzz/virzz; brew install <formula>`' >> README.md; \
	echo '' >> README.md; \
	echo '### Formulae' >> README.md; \
	echo '' >> README.md; \
	echo '- Enyo `brew install virzz/virzz/enyo` || `brew tap virzz/virzz; brew install enyo`' >> README.md; \
	echo '' >> README.md; 

	@go run ./internal/_compile enyo;
	@echo "Add Enyo"; \
	echo '## Enyo - CLI 命令行小工具' >> README.md; \
	echo '' >> README.md; \
	echo '```' >> README.md; \
	./build/enyo >> README.md; \
	echo '```' >> README.md; \
	echo '' >> README.md; \

	@echo "Compile Public Projects"
	@go run ./internal/_compile public;

	@echo "Add Public Project List"; \
	echo "## Public Projects" >> README.md; \
	echo '' >> README.md; \
	for i in `ls -d ./cmd/public/*/`; do \
		echo "- $$(basename $$i)" >> README.md; \
	done; \
	echo '' >> README.md;

	@echo "Add Public Projects"; \
	for i in `ls -d ./cmd/public/*/`; do \
		echo "## $$(basename $$i)" >> README.md; \
		echo '' >> README.md; \
		echo '```' >> README.md; \
		echo "$$(./build/$$(basename $$i) -h )" >> README.md; \
		echo '```' >> README.md; \
		echo '' >> README.md; \
	done;

ghext: rc-gh-mozhu
	@gh extension remove gh-mozhu > /dev/null || true
	@./build/gh-mozhu install
	gh extension list
