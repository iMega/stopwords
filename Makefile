CWD=/go/src/github.com/imega/stopwords

build: dep test
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-e CGO_ENABLED=0 \
		golang:1.8-alpine \
		sh -c 'go build -v -o stopword'

test:
	@docker run --rm -v $(CURDIR):$(CWD) -w $(CWD) \
		golang:1.8-alpine sh -c "go list ./... | grep -v 'vendor' | xargs go test -v"

dep:
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		golang:1.8-alpine sh -c 'apk add --update git && go get -u github.com/golang/dep/cmd/dep && dep ensure -v'

test2:
	go list ./... | grep -v 'vendor' | xargs go test