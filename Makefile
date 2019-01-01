VERSION         :=      $(shell cat ./VERSION)
IMAGE_NAME      :=      ban-app
GC_PROJECT_ID	:=		ban-app
GCR_HOSTNAME	:= 		eu.gcr.io

all: install

test:
	go test ./... -v

fmt:
	go fmt ./... -v

build:
	go build -v

change:
	make test && make image && make release && make deploy

image:
	docker build -t ${IMAGE_NAME}:$(VERSION) .

release:
	docker tag ${IMAGE_NAME}:$(VERSION) ${GCR_HOSTNAME}/${GC_PROJECT_ID}/${IMAGE_NAME}:${VERSION}
	docker push ${GCR_HOSTNAME}/${GC_PROJECT_ID}/${IMAGE_NAME}:${VERSION}

deploy:
	kubectl apply -f env/app.yml
	kubectl rollout status deployment/${GC_PROJECT_ID}