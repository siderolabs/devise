SHELL := /bin/bash
BUILDDEPS := \
	github.com/golang/dep/cmd/dep \
	github.com/autonomy/drydock

NAMESPACE := autonomy
NAME := devise
TAG ?= develop
IMAGE := ${NAMESPACE}/${NAME}:${TAG}

GITCOMMIT := $(shell if [ -z "$$(git status --porcelain)" ]; then git rev-parse --short HEAD; else echo "dirty"; fi)
BUILT := $(shell date)


all: clean vendor
	@drydock build --template test -- \
		--tag ${IMAGE} \
		--build-arg TAG="${TAG}" \
		--build-arg GITCOMMIT="${GITCOMMIT}" \
		--build-arg BUILT="${BUILT}" \
		.
	@docker run --rm -it --volume $(shell pwd):/out ${IMAGE} cp coverage.txt /out
	@drydock build --template $@ -- \
		--tag ${IMAGE} \
		--build-arg TAG="${TAG}" \
		--build-arg GITCOMMIT="${GITCOMMIT}" \
		--build-arg BUILT="${BUILT}" \
		.

.PHONY: init
init:
	@for b in  $(BUILDDEPS); do \
		echo "Installing $$b"; \
		go get -u $$b; \
	done

vendor:
	@echo [Fetching dependencies]
	@dep ensure

.PHONY: build
build: vendor
	@echo [Building Devise ${TAG}-${GITCOMMIT}]
	@drydock build --template $@ -- \
		--tag ${IMAGE}-$@ \
		--build-arg TAG="${TAG}" \
		--build-arg GITCOMMIT="${GITCOMMIT}" \
		--build-arg BUILT="" \
		.

.PHONY: test
test: vendor
	@echo [Running tests]
	@drydock build --template $@ -- \
		--tag ${IMAGE}-$@ \
		--build-arg TAG="${TAG}" \
		--build-arg GITCOMMIT="${GITCOMMIT}" \
		--build-arg BUILT="" \
		.

.PHONY: image
image: vendor
	@echo [Building ${IMAGE}]
	@drydock build --template $@ -- \
		--tag ${IMAGE} \
		--build-arg TAG="${TAG}" \
		--build-arg GITCOMMIT="${GITCOMMIT}" \
		--build-arg BUILT="" \
		.

.PHONY: api
api: build
	@echo [Generating API]
	@docker run \
		--rm \
		-it \
		--volume $(shell pwd):/out \
		${IMAGE}-build cp -R ./api /out

# TODO: docs

.PHONY: push
push:
	@echo [Pushing ${IMAGE}]
	@docker login -u "${DOCKER_USERNAME}" -p "${DOCKER_PASSWORD}";
	@docker push ${IMAGE};

.PHONY: run
run: image stop
	@echo [Running ${IMAGE}]
	@docker run --rm -d -p 8080:8080 -p 50000:50000 --name ${NAME} ${IMAGE} serve

.PHONY: stop
stop:
	@echo [Stopping ${IMAGE}]
	-@docker stop ${NAME}

.PHONY: example-client
example-client: run
	@echo [Running example client]
	@cd examples/client && go run main.go

.PHONY: example-wrapper
example-wrapper: run
	@echo [Running example wrapper]
	@docker run \
		--rm \
		-it \
		-v $(shell pwd)/examples/wrapper:/app/plan \
		--network=host \
		${IMAGE} implement --plan=/app/plan/plan.yaml

.PHONY: clean
clean:
	@echo [Cleaning]
	@dep ensure
	@dep prune
	@cat .gitignore | while read line; do rm -rf "$$line" ; done
