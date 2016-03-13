GOLANG_IMAGE = golang:1.6-alpine

TEST_CONTAINER_IDS_FILE = setup-test.container-ids

setup-test:
	test -f $(TEST_CONTAINER_IDS_FILE) || ./setup-test.sh > $(TEST_CONTAINER_IDS_FILE)

test: setup-test
	docker run --rm -v $(GOPATH):/go -v $(PWD):/app -w /app $(GOLANG_IMAGE) go test

build:
	docker run --rm -v $(GOPATH):/go -v $(PWD):/app -w /app -e CGO_ENABLED=0 $(GOLANG_IMAGE) go build -ldflags "-s" -a -installsuffix cgo -o tadis

clean:
	cat $(TEST_CONTAINER_IDS_FILE) | xargs docker stop | xargs docker rm
	rm -f $(TEST_CONTAINER_IDS_FILE)
