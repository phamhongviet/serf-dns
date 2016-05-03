#!/bin/bash
set -ev

publish() {
	make build
	docker build -t "$DOCKER_REPO":"$DOCKER_TAG" .
	docker login -e="$DOCKER_EMAIL" -u="$DOCKER_USERNAME" -p="$DOCKER_PASSWORD" "$DOCKER_REGISTRY"
	docker push "$DOCKER_REPO":"$DOCKER_TAG"
}


if [ "$(git rev-parse --abbrev-ref HEAD)" = "master" ]; then
	DOCKER_TAG="latest"
	publish
fi
