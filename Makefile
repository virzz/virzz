#!/usr/bin/env make
TARGET=./build
ARCHS=amd64 386
LDFLAGS="-s -w"
GCFLAGS="all=-trimpath=$(shell pwd)"
ASMFLAGS="all=-trimpath=$(shell pwd)"
GOPATH=$(shell go env GOPATH)
SOURCE="./cli/"
ifeq ($(B), all)
APPNAMES := $(shell cd ./cli/ && ls)
else ifneq ($(B),)
APPNAMES := $B
else 
APPNAMES := virzz
endif

current:
	@mkdir -p ${TARGET}/
	@rm -rf ./${TARGET}/*
	@for APPNAME in ${APPNAMES}; do \
		echo "[*] Building $${APPNAME} ..." ; \
		go build -o ${TARGET}/$${APPNAME} ${SOURCE}/$${APPNAME} ; \
	done; \
	echo "[+] Current Built."

windows:
	@for GOARCH in ${ARCHS}; do \
		echo "[*] Building for windows $${GOARCH} ..." ; \
		for APPNAME in ${APPNAMES}; do \
			echo "[+] $${APPNAME} ..." ; \
			GOOS=windows GOARCH=$${GOARCH} GO111MODULE=on CGO_ENABLED=0 \
			go build -ldflags=${LDFLAGS} -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} \
			-o ${TARGET}/$${APPNAME}-windows-$${GOARCH}.exe ${SOURCE}/$${APPNAME}; \
		done; \
	done; \
	echo "[+] Windows Built."

linux:
	@for GOARCH in ${ARCHS}; do \
		echo "[*] Building for linux $${GOARCH} ..." ; \
		for APPNAME in ${APPNAMES}; do \
			echo "[+] $${APPNAME} ..." ; \
			GOOS=linux GOARCH=$${GOARCH} GO111MODULE=on CGO_ENABLED=0 \
			go build -ldflags=${LDFLAGS} -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} \
			-o ${TARGET}/$${APPNAME}-linux-$${GOARCH} ${SOURCE}/$${APPNAME}; \
		done; \
	done; \
	echo "[+] Linux Built."

darwin:
	@for GOARCH in ${ARCHS}; do \
		echo "[*] Building for darwin $${GOARCH} ..." ; \
		for APPNAME in ${APPNAMES}; do \
			echo "[+] $${APPNAME} ..." ; \
			GOOS=darwin GOARCH=$${GOARCH} GO111MODULE=on CGO_ENABLED=0 \
			go build -ldflags=${LDFLAGS} -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} \
			-o ${TARGET}/$${APPNAME}-darwin-$${GOARCH} ${SOURCE}/$${APPNAME}; \
		done; \
	done; \
	echo "[+] Darwin Built."

all: clean darwin linux windows

fmt:
	@go fmt ./...; \
	echo "[+] Fmted."

update:
	@go get -u; \
	go mod tidy -v; \
	echo "[+] Updated."

link: current
	@for APPNAME in ${APPNAMES}; do \
		echo "[*] Link $${APPNAME} ..." ; \
		test -f /usr/local/bin/$${APPNAME} && rm /usr/local/bin/$${APPNAME}; \
		ln -s `pwd`/${TARGET}/$${APPNAME} /usr/local/bin/$${APPNAME}; \
		test -f /usr/local/bin/$${APPNAME} && echo "[+] $${APPNAME} Linked" || echo "[-] Fail"; \
	done;

install: current
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
	go clean ./... ; \
	echo "[+] Cleaned."

archive: all
	@mkdir release; \
	cd ${TARGET}; \
	for APPNAME in ${APPNAMES}; do \
		echo "[+] Archive $${APPNAME} ..." ; \
		sha256sum $${APPNAME}* > SHA256.txt; \
		zip ../release/$${APPNAME}.zip -9 $${APPNAME}*; \
	done;
