SHELL := /bin/bash
BUILDDEPS := \
	github.com/golang/dep/cmd/dep \
	github.com/autonomy/drydock

SHA := $(shell if [ -z "$$(git status --porcelain)" ]; then git rev-parse --short HEAD; else echo "dirty"; fi)
BUILT := $(shell date)

NAMESPACE := autonomy
NAME := devise
RELEASE ?= edge
IMAGE := ${NAMESPACE}/${NAME}:${SHA}
IMAGE_RELEASE := ${NAMESPACE}/${NAME}:${RELEASE}
IMAGE_LATEST := ${NAMESPACE}/${NAME}:latest


all: clean vendor
	@drydock build --template test -- \
		--tag ${IMAGE} \
		--build-arg RELEASE="${RELEASE}" \
		--build-arg SHA="${SHA}" \
		--build-arg BUILT="${BUILT}" \
		.
	@docker run --rm -it --volume $(shell pwd):/out ${IMAGE} cp coverage.txt /out
	@drydock build --template $@ -- \
		--tag ${IMAGE} \
		--build-arg RELEASE="${RELEASE}" \
		--build-arg SHA="${SHA}" \
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
	@echo [Building Devise ${RELEASE}-${SHA}]
	@drydock build --template $@ -- \
		--tag ${IMAGE}-$@ \
		--build-arg RELEASE="${RELEASE}" \
		--build-arg SHA="${SHA}" \
		--build-arg BUILT="" \
		.

.PHONY: test
test: vendor
	@echo [Running tests]
	@drydock build --template $@ -- \
		--tag ${IMAGE}-$@ \
		--build-arg RELEASE="${RELEASE}" \
		--build-arg SHA="${SHA}" \
		--build-arg BUILT="" \
		.

.PHONY: image
image: vendor
	@echo [Building ${IMAGE}]
	@drydock build --template $@ -- \
		--tag ${IMAGE} \
		--build-arg RELEASE="${RELEASE}" \
		--build-arg SHA="${SHA}" \
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

# TODO: Verify that $RELEASE is of the format vMAJOR.MINOR.PATCH*
.PHONY: push
push:
	@echo [Tagging ${IMAGE} as ${IMAGE_LATEST}]
	@docker tag ${IMAGE} ${IMAGE_LATEST}
	@echo [Tagging ${IMAGE} as ${IMAGE_RELEASE}]
	@docker tag ${IMAGE} ${IMAGE_RELEASE}
	@docker login -u "${DOCKER_USERNAME}" -p "${DOCKER_PASSWORD}"
ifneq (${RELEASE},develop)
	@echo [Pushing ${NAMESPACE}/${NAME}:$(shell echo $${RELEASE:1})]
	@docker push ${NAMESPACE}/${NAME}:$(shell echo $${RELEASE:1})
else
	@echo Tag is 'develop', not pushing ...
endif
ifneq (${SHA},dirty)
	@echo [Pushing ${IMAGE_LATEST}]
	@docker push ${IMAGE_LATEST}
	@echo [Pushing ${IMAGE_SHA}]
	@docker push ${IMAGE_SHA}
else
	@echo Git commit is 'dirty', not pushing ...
endif

.PHONY: run
run: image stop
	@echo [Running ${IMAGE}]
	@docker run \
		--rm \
		-d \
		-p 8080:8080 \
		-p 50000:50000 \
		--name ${NAME} \
		${IMAGE} serve --vault-address=${VAULT_ADDR}

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
		${IMAGE} implement --plan=/app/plan/plan.yaml --vault-token=${VAULT_TOKEN}

.PHONY: clean
clean:
	@echo [Cleaning]
	@dep ensure
	@dep prune
	@cat .gitignore | while read line; do rm -rf "$$line" ; done
