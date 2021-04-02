img=josofm/cartesian
run=docker run -p 80:80 --rm -ti -v `pwd`:/app $(img)
cov=coverage.out
covhtml=coverage.html

all: check build

image:
	docker build . -t $(img)

shell: image
	$(run) sh

build: image
	$(run) go build -o ./cmd/cartesian/cartesian ./cmd/cartesian

check: image
	$(run) go test -timeout 60s -race -coverprofile=$(cov) ./...

check-integration: image
	$(run) go test -timeout 120s -race -coverprofile=$(cov) -tags=integration ./...

coverage: check
	$(run) go tool cover -html=$(cov) -o=$(covhtml)
	xdg-open coverage.html

run: build
	$(run) ./cmd/cartesian/cartesian

