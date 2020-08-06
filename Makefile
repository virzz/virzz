TARGET=./build
ARCHS=amd64 386
LDFLAGS="-s -w"
GCFLAGS="all=-trimpath=$(shell pwd)"
ASMFLAGS="all=-trimpath=$(shell pwd)"
SOURCE="./cli/"
APPNAME=virzz

current:
	@mkdir -p ${TARGET}/
	@rm -f ./${TARGET}/${APPNAME}
	@go build -o ${TARGET}/${APPNAME} ${SOURCE}; \
	echo "Done."

fmt:
	@go fmt ./...; \
	echo "Done."

update:
	@go get -u; \
	go mod tidy -v; \
	echo "Done."

linux:
	@for GOARCH in ${ARCHS}; do \
		echo "Building for linux $${GOARCH} ..." ; \
		mkdir -p ${TARGET}/${APPNAME}-linux-$${GOARCH} ; \
		GOOS=linux GOARCH=$${GOARCH} GO111MODULE=on CGO_ENABLED=0 go build -ldflags=${LDFLAGS} -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} -o ${TARGET}/${APPNAME}-linux-$${GOARCH}/${APPNAME} ${SOURCE}; \
	done; \
	echo "Done."

darwin:
	@for GOARCH in ${ARCHS}; do \
		echo "Building for darwin $${GOARCH} ..." ; \
		mkdir -p ${TARGET}/${APPNAME}-darwin-$${GOARCH} ; \
		GOOS=darwin GOARCH=$${GOARCH} GO111MODULE=on CGO_ENABLED=0 go build -ldflags=${LDFLAGS} -gcflags=${GCFLAGS} -asmflags=${ASMFLAGS} -o ${TARGET}/${APPNAME}-darwin-$${GOARCH}/${APPNAME} ${SOURCE} ; \
	done; \
	echo "Done."

release: darwin linux

all: clean fmt update lint test darwin linux

test:
	@go test -v -race ./... ; \
	echo "Done."

lint:
	@go get -u github.com/golangci/golangci-lint@master ; \
	golangci-lint run ./... ; \
	go mod tidy ; \
	echo Done

install: current
	cp -f ${TARGET}/${APPNAME} /usr/local/bin/${APPNAME}

clean:
	@rm -rf ${TARGET}/* ; \
	go clean ./... ; \
	echo "Done."
