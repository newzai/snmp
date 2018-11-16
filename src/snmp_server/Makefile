BINARY_NAME=snmp_server

VERSION=`cat VERSION`
BUILD_TIME=`date +%FT%T%z`
GIT_HASH=`git rev-parse HEAD`

LDFLAGS="-X snmp_server/globalvars.AppVersion=${VERSION} -X snmp_server/globalvars.AppBuildTime=${BUILD_TIME} -X snmp_server/globalvars.AppGitHash=${GIT_HASH}"

build: ## Build the binary
	go build -ldflags ${LDFLAGS} -o ${BINARY_NAME} snmp_server.go


clean: ## Clean this build
	rm -rf ${BINARY_NAME}
	rm -rf build


help:  ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: all
all: build

