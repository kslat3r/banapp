FROM golang:1.11.4 

WORKDIR /usr/local/go/src/ban-app
ADD . .

RUN make install
RUN make build

CMD ban-app
EXPOSE 8080