FROM golang:1.11.4 

WORKDIR /usr/local/go/src/bananapp
ADD . .

RUN go install

CMD bananapp
EXPOSE 8080