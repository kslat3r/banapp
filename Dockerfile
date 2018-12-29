FROM golang:1.11.4 

WORKDIR /usr/local/go/src/bananagrams-api
ADD . .

RUN make install
RUN make test
RUN make build

CMD bananagrams-api