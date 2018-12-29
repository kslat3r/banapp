VERSION         :=      $(shell cat ./VERSION)
IMAGE_NAME      :=      bananagrams-api
GC_PROJECT_ID	:=		bananagrams
GCR_HOSTNAME	:= 		eu.gcr.io

all: install

install:
	go install -v

test:
	go test ./... -v

fmt:
	go fmt ./... -v

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