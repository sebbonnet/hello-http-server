project_dir := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))

VERSION ?= $(shell git rev-parse --short HEAD)
REPO ?= "docker.io"
NAMESPACE ?= "seb-test"

build:
	VERSION=$(VERSION) REPO=$(REPO) docker-compose build test-app
	VERSION=$(VERSION) REPO=$(REPO) docker-compose push test-app

deploy:
	VERSION=$(VERSION) REPO=$(REPO) envsubst < $(project_dir)/pod.yml | kubectl -n $(NAMESPACE) apply -f -
