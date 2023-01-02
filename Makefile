#!/usr/bin/env make

TARGET=./build

default:
	go run ./cli/_compile virzz


%:
	@rm -f ${TARGET}/$@ ; \
	if [[ -d ./cli/public/$@ ]];then \
		go run ./cli/_compile $@ ; \
	elif [[ -d ./cli/$@ ]]; then \
		go run ./cli/_compile $@ ; \
	else \
		echo "Not found target project: $@"; \
	fi


install: virzz
	@echo "[*] Install virzz ..." ; \
	cp -f ${TARGET}/virzz ${GOPATH}/bin/virzz; \
	test -f ${GOPATH}/bin/virzz && echo "[+] virzz Installed";

uninstall: remove

remove:
	@echo "[*] Remove virzz ..." ; \
	rm -f ${GOPATH}/bin/virzz; \
	test -f ${GOPATH}/bin/virzz || \
	echo "[+] virzz Removed";

clean:
	@go run ./cli/_compile -C

cleanr:
	@rm -rf release; \
	go clean ./... ; \
	echo "[+] Cleaned."

readme:
	@echo "# Virzz" > README.md ; \
	echo '![Build](https://github.com/virzz/virzz/workflows/Build/badge.svg)' >> README.md; \
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
	cat README.md

docker:
	@test -f ./build/platform-linux-amd64 && \
	cp ./build/platform-linux-amd64 ./deploy/platform-linux-amd64 && \
	cd ./deploy && \
	command -v docker-compose && \
	docker-compose build platform && \
	echo "[+] Success" && rm -f platform-linux-amd64 || echo "[-] Fail"