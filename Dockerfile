FROM golang:1.14-alpine as development

ENV CGO_ENABLED=1

WORKDIR /go/src/app

RUN apk update && apk add git
RUN go get github.com/cespare/reflex
COPY . .
RUN go build -o elasticsearch-tool main.go

CMD ["reflex", "-c", "./reflex.conf"]

FROM alpine AS build
COPY --from=development /go/src/app/elasticsearch-tool /opt/elasticsearch-tool
ENTRYPOINT ["/opt/elasticsearch-tool"]