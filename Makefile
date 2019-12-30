GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run 
GOTEST=$(GOCMD) test
BUILDPATH=./release
BINARY_NAME=mathgen-web
DOCKER_USERNAME=ledongthuc
DOCKER_IMG_NAME=$(BINARY_NAME)
DOCKER_VERSION=latest

all: test clean build
clean:
	rm -rf $(BUILDPATH); mkdir $(BUILDPATH)
test: 
	$(GOTEST) ./...
build: 
	$(GOBUILD) -o $(BUILDPATH)/$(BINARY_NAME) -v ./web/main.go
run:
	$(GORUN) ./web/main.go
docker-clean:
	docker rmi --force $(DOCKER_USERNAME)/$(DOCKER_IMG_NAME):$(DOCKER_VERSION)
docker-build:
	docker build -t $(DOCKER_USERNAME)/$(DOCKER_IMG_NAME):$(DOCKER_VERSION) .
docker-push: docker-build
	docker push $(DOCKER_USERNAME)/$(DOCKER_IMG_NAME):$(DOCKER_VERSION)
docker-run:
	docker run -p 8080:8080 $(DOCKER_USERNAME)/$(BINARY_NAME):$(DOCKER_VERSION)
