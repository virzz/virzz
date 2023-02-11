#!/usr/bin/env make

TARGET=./build

default:
	go run ./cli/_compile god

public:
	go run ./cli/_compile public

%:
	@rm -f ${TARGET}/$@ ; 
	@if [[ -d ./cli/public/$@ ]] || [[ -d ./cli/$@ ]]; then \
		if [[ -z "${DEBUG}" ]]; then \
			go run -tags debug ./cli/_compile $@ ; \
		else \
			go run ./cli/_compile $@ ; \
		fi; \
	fi;

rc-%:
	@echo "[*] Compiling Release $(subst rc-,,$@) ..." ; \
	go run ./cli/_compile -R $(subst rc-,,$@)

i-%: rc-%
	@export NAME='$(subst i-,,$@)'; \
	echo "[*] Install Release $${NAME} ..." ; \
	cp -f ${TARGET}/$${NAME} ${GOPATH}/bin/$${NAME}; \
	test -f ${GOPATH}/bin/$${NAME} && echo "[+] $${NAME} Installed";

clean:
	@go run ./cli/_compile -C

cleanr:
	@rm -rf release; \
	go clean ./... ; \
	echo "[+] Cleaned."

swagger:
	@swag i -g services/web/swagger.go -o services/docs

readme:
	@echo "# Virzz" > README.md ; \
	echo '' >> README.md;
	@if test -f ${TARGET}/virzz ; then \
		echo '## Virzz - CLI 命令行小工具' >> README.md; \
		echo '' >> README.md; \
		echo '```' >> README.md; \
		./build/virzz >> README.md; \
		echo '```' >> README.md; \
		echo '' >> README.md; \
	fi
	@if test -f ${TARGET}/platform ; then \
		echo '## Virzz - Platform 服务端工具' >> README.md; \
		echo '' >> README.md; \
		echo '```' >> README.md; \
		./build/platform >> README.md; \
		echo '```' >> README.md; \
	fi
