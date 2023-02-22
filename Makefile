#!/usr/bin/env make

TARGET=./build

default:
	@if [[ -z "${DEBUG}" ]]; then \
		go run -tags debug ./internal/_compile $@ ; \
	else \
		go run ./internal/_compile $@ ; \
	fi;

%:
	@rm -f ${TARGET}/$@ ; 
	@if [[ -d ./cli/public/$@ ]] || [[ -d ./cli/$@ ]]; then \
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
