#!/usr/bin/env make
TARGET=./build
APPNAMES=virzz platform
OSS=linux windows darwin
ARCHS=amd64 arm64
GCFLAGS="all=-trimpath=$(shell pwd)"
ASMFLAGS="all=-trimpath=$(shell pwd)"
GOPATH=$(shell go env GOPATH)
SOURCE="./cli/"
PUBLICS := $(shell cd ./cli/public/ && ls)
# BUILD_ID := $(shell head .buildid)
VERSION := $(shell git tag | tail -1)
LDFLAGS := -s -w \
	-X github.com/mozhu1024/virzz/common.Mode=prod

%:
	@function failx(){ \
		echo "[-] $$1 Fail."; \
		export B=1; \
	}; \
	if [[ -d ./cli/public/$@ ]];then \
		echo "Building [$@]"; \
		echo "[*] Building [$@] ... "; \
		rm -f ./${TARGET}/$@; \
		BUILD_ID=`head .buildid/$@ 2>/dev/null || echo 0` ; \
		LDFLAGS="${LDFLAGS} -X main.BuildID=$${BUILD_ID} -X main.Version=dev" ; \
		GOOS=$${OS} GOARCH=$${GOARCH} GO111MODULE=on CGO_ENABLED=0 \
		go build -ldflags "$${LDFLAGS}" -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} \
			-o ${TARGET}/$@ ./cli/public/$@ && \
			echo "[+] $@ Built." || failx $@ ;\
		if [ -z "$${B}" ]; then \
			echo "[+] BuildID = $${BUILD_ID}"; \
			expr $${BUILD_ID} + 1 > .buildid/$@; \
		fi \
	elif [[ "$@" = "public" ]]; then \
		echo Build ${PUBLICS}; \
		for APP in ${PUBLICS}; do \
			echo "[*] Building $${APP} ... "; \
			rm -f ./${TARGET}/$${APP}; \
			BUILD_ID=`head .buildid/$${APP} 2>/dev/null || echo 0` ; \
			LDFLAGS="${LDFLAGS} -X main.BuildID=$${BUILD_ID} -X main.Version=dev" ; \
			GOOS=$${OS} GOARCH=$${GOARCH} GO111MODULE=on CGO_ENABLED=0 \
			go build -ldflags "$${LDFLAGS}" -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} \
				-o ${TARGET}/$${APP} ./cli/public/$${APP}  && \
				echo "[+] $${APP} Built." || failx $${APP} ; \
			if [ -z "$${B}" ]; then \
				echo "[+] BuildID = $${BUILD_ID}"; \
				expr $${BUILD_ID} + 1 > .buildid/$${APP}; \
			fi \
		done; \
		echo "[+] Finish."; \
	fi

virzz:
	@mkdir -p ${TARGET}/
	@rm -f ./${TARGET}/$@
	@echo "[*] Building [$@] ..." ;
	@go build -o ${TARGET}/$@ ${SOURCE}/$@  && \
			echo "[+] $@ Built." || \
			echo "[-] Build [$@] Faild";

release: clean
	@function fail(){ \
		echo "[-] $$1 Fail."; \
		export B=1; \
	}; \
	for APPNAME in ${APPNAMES}; do \
		export B=0; \
		BUILD_ID=`head .buildid/$${APPNAME} 2>/dev/null || echo 0` ; \
		LDFLAGS="${LDFLAGS} -X main.BuildID=$${BUILD_ID} -X main.Version=${VERSION}" ; \
		for OS in ${OSS}; do \
			for GOARCH in ${ARCHS}; do \
				echo "[*] Building for $${APPNAME} $${OS} $${GOARCH} ..." ; \
				GOOS=$${OS} GOARCH=$${GOARCH} GO111MODULE=on CGO_ENABLED=0 \
				go build -ldflags "$${LDFLAGS}" -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} \
				-o ${TARGET}/$${APPNAME}-$${OS}-$${GOARCH} ${SOURCE}/$${APPNAME} && \
				echo "[+] $${APPNAME}-$${OS}-$${GOARCH} Built." || \
				fail $${APPNAME}-$${OS}-$${GOARCH} ; \
			done; \
		done; \
		if [ $${B}="0" ]; then \
			echo "[+] BuildID = $${BUILD_ID}"; \
			expr $${BUILD_ID} + 1 > .buildid/$${APPNAME}; \
		fi \
	done;

archive: release
	@rm -rf release; \
	mkdir release; \
	echo "[+] Archive ..." ; \
	shasum -a 256 ./${TARGET}/* > ./${TARGET}/SHA256.txt; \
	zip ./release/virzz.zip -9 ./${TARGET}/* ;

install: virzz
	@for APPNAME in ${APPNAMES}; do \
		echo "[*] Install $${APPNAME} ..." ; \
		cp -f ${TARGET}/$${APPNAME} ${GOPATH}/bin/$${APPNAME}; \
		test -f ${GOPATH}/bin/$${APPNAME} && echo "[+] $${APPNAME} Installed"; \
	done;

uninstall: remove

remove:
	@for APPNAME in ${APPNAMES}; do \
		echo "[*] Remove $${APPNAME} ..." ; \
		rm -f ${GOPATH}/bin/$${APPNAME}; \
		rm /usr/local/bin/$${APPNAME}; \
		test -f ${GOPATH}/bin/$${APPNAME} || \
		test -f ${GOPATH}/bin/$${APPNAME}  || \
		echo "[+] $${APPNAME} Removed"; \
	done;

clean:
	@rm -rf ${TARGET}/* ; \
	rm -rf release; \
	go clean ./... ; \
	echo "[+] Cleaned."

readme:
	@echo "# Virzz" > README.md ; \
	echo '![Build](https://github.com/mozhu1024/virzz/workflows/Build/badge.svg)' >> README.md; \
	echo '' >> README.md;
	@if test -f ./build/virzz ; then \
		echo '## Virzz - CLI 命令行小工具' >> README.md; \
		echo '' >> README.md; \
		echo '```' >> README.md; \
		./build/virzz >> README.md; \
		echo '```' >> README.md; \
		echo '' >> README.md; \
	fi
	@if test -f ./build/platform ; then \
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