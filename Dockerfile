FROM golang:1.21.5-alpine3.19
ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/solvencyanalytics
COPY . .
ENTRYPOINT ["go", "test", "-cover", "-bench=.", "-benchtime=1s" ,"./..."]