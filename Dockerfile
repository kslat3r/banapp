FROM golang:1.11.4 

WORKDIR /usr/local/go/src/ban-api/src
ADD . .

RUN make install
RUN make test
RUN make build

CMD ban-api
EXPOSE 8080