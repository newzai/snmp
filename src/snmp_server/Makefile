BINARY_NAME=snmp_server

VERSION=`cat VERSION`
BUILD_TIME=`date +%FT%T%z`
GIT_HASH=`git rev-parse HEAD`
GoVersion=`go version`
LDFLAGS="-X snmp_server/globalvars.AppVersion=${VERSION} -X 'snmp_server/globalvars.GoVersion=${GoVersion}' -X snmp_server/globalvars.AppBuildTime=${BUILD_TIME} -X snmp_server/globalvars.AppGitHash=${GIT_HASH}"

release: ## Build the binary
	go build -v -a  -ldflags ${LDFLAGS} -o ${BINARY_NAME} snmp_server.go
	-git push origin --delete v${VERSION}
	-git tag -d v${VERSION}
	git tag -a v${VERSION}  -m "version v${VERSION} build data ${BUILD_TIME}, go version ${GoVersion}  by ${USER}"
	git push origin v${VERSION}

debug:
	go build -v -a  -ldflags ${LDFLAGS} -o ${BINARY_NAME} snmp_server.go
 
clean: ## Clean this build
	rm -rf ${BINARY_NAME}
	rm -rf build


help:  ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: all
all: debug

