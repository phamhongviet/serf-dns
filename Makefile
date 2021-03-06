GOLANG_IMAGE = golang:1.6-alpine

TEST_CONTAINER_IDS_FILE = setup-test.container-ids

get-deps:
	go get github.com/golang/lint/golint
	go get github.com/hashicorp/serf/client
	go get github.com/miekg/dns
	go get github.com/paked/configure

update-deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/hashicorp/serf/client
	go get -u github.com/miekg/dns
	go get -u github.com/paked/configure

setup-test:
	test -f $(TEST_CONTAINER_IDS_FILE) || ./setup-test.sh > $(TEST_CONTAINER_IDS_FILE)

test: setup-test get-deps
	docker run --rm -v $(GOPATH):/go -v $(PWD):/app -w /app --link `head -n 1 $(TEST_CONTAINER_IDS_FILE)`:serf --link `sed '2q;d' $(TEST_CONTAINER_IDS_FILE)`:serf-auth $(GOLANG_IMAGE) sh -c "gofmt -d . && golint ./... && go test"

build: get-deps
	docker run --rm -v $(GOPATH):/go -v $(PWD):/app -w /app -e CGO_ENABLED=0 $(GOLANG_IMAGE) go build -ldflags "-s" -a -installsuffix cgo -o serf-dns

clean:
	cat $(TEST_CONTAINER_IDS_FILE) | xargs docker stop | xargs docker rm
	rm -f $(TEST_CONTAINER_IDS_FILE)
