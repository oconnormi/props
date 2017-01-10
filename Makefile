NAME=props
REPO=github.com/oconnormi/props
# Get the current full sha from git
GITSHA:=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GOX_OUTPUT=-output "build/${NAME}_{{.OS}}_{{.Arch}}"
GOX=CGO_ENABLED=0 gox

default: deps dev

clean:
	rm -rf ./build

deps:
	go get github.com/mitchellh/gox
	go get github.com/kardianos/govendor
	govendor sync

dev:
	echo "Building dev version"
	${GOX} -osarch="linux/amd64 darwin/amd64" \
		-ldflags "-X github.com/oconnormi/props/version.GitCommit=${GITSHA}${GIT_DIRTY}" \
		${GOX_OUTPUT} \
		${REPO}

build:
	echo "==> Building..."
		${GOX} ${GOX_OUTPUT} ${REPO}

.PHONY: default deps dev build
