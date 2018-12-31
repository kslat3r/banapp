VERSION         :=      $(shell cat ./VERSION)
IMAGE_NAME      :=      ban-api
GC_PROJECT_ID	:=		ban-online
GCR_HOSTNAME	:= 		eu.gcr.io

all: install

install:
	go install -v

test:
	go test ./src/... -v

fmt:
	go fmt ./src/... -v

build:
	go build -v

image:
	docker build -t ${IMAGE_NAME}:$(VERSION) .

release:
	docker tag ${IMAGE_NAME}:$(VERSION) ${GCR_HOSTNAME}/${GC_PROJECT_ID}/${IMAGE_NAME}:${VERSION}
	docker push ${GCR_HOSTNAME}/${GC_PROJECT_ID}/${IMAGE_NAME}:${VERSION}

tag:
	git tag -a $(VERSION) -m "Release" || true
	git push origin $(VERSION)