project_dir := $(realpath $(dir $(firstword $(MAKEFILE_LIST))))

VERSION ?= $(shell git rev-parse --short HEAD)
REPO ?= "docker.io"
NAMESPACE ?= "my-test"

build:
	VERSION=$(VERSION) REPO=$(REPO) docker-compose build my-app
	VERSION=$(VERSION) REPO=$(REPO) docker-compose push my-app

deploy:
	VERSION=$(VERSION) REPO=$(REPO) envsubst < $(project_dir)/pod.yml | kubectl -n $(NAMESPACE) apply -f -
