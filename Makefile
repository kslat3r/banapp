VERSION         :=      $(shell cat ./VERSION)
IMAGE_NAME      :=      ban-app
GC_PROJECT_ID	:=		ban-app
GCR_HOSTNAME	:= 		eu.gcr.io
GCR_CONTEXT		:=		gke_ban-app_europe-west2-a_ban-app-prod

all: install

install:
	go install -v

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
	kubectl config use-context ${GCR_CONTEXT}
	kubectl apply -f env/app.yml