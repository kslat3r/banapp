FROM golang:1.11.4 

WORKDIR /usr/local/go/src/banapp
ADD . .

RUN go install

CMD banapp
EXPOSE 8080