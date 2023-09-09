
# Fat image with sources
FROM golang:1.19 as development

# Go requirements files
WORKDIR /go/project
COPY go.work go.work

WORKDIR /go/project/hw
COPY hw/go.mod .
COPY hw/go.sum .

WORKDIR /go/project/hw/cmd
COPY hw/cmd/go.mod .
COPY hw/cmd/go.sum .

# ----

RUN go mod download -x

# Source-code files
COPY hw/cmd/main.go . 

WORKDIR /go/project/hw/sources
COPY hw/sources/config.go .
COPY hw/sources/httpserver.go .

# Special compile with out additional info
WORKDIR /go/project
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /opt/service/service ./hw/cmd/

# Slim image
FROM alpine:3.9 as production

LABEL MODULE="Http server"

COPY --from=development /opt/service/service /opt/service/service

COPY hw/configs/http.yaml "/etc/service/config.yaml"

# One imege - one process
ENTRYPOINT [ "/opt/service/service", "--config", "/etc/service/config.yaml" ]
