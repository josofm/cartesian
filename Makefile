goversion=1.16
img=josofm/cartesian
run=docker run -p 80:80 $(img)
vols=-v `pwd`:/app -w /app
run_test=docker run --rm $(vols) golang:$(goversion)
cov=coverage.out
covhtml=coverage.html

all: check build

image:
	docker build . -t $(img)

build: image
	$(run_test) go build -o ./cmd/cartesian/cartesian ./cmd/cartesian

check: image
	$(run_test) go test -timeout 60s -race -coverprofile=$(cov) ./...

check-integration: image
	$(run_test) go test -timeout 120s -race -coverprofile=$(cov) -tags=integration ./...

coverage: check
	$(run_test) go tool cover -html=$(cov) -o=$(covhtml)
	xdg-open coverage.html

run: image
	$(run)

