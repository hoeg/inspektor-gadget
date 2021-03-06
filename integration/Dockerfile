FROM golang:1.14-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /build
RUN go get -d -v github.com/kr/pretty
RUN go install -v github.com/kr/pretty
COPY . .
RUN go test -c -o gadget-integration.test ./...

FROM alpine:3.11
RUN apk add ca-certificates curl
ENV KUBECTL_GADGET /bin/kubectl-gadget

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
RUN chmod +x ./kubectl
RUN mv ./kubectl /usr/local/bin

COPY bin/kubectl-gadget /bin/kubectl-gadget
COPY --from=build_base /build/gadget-integration.test /bin/gadget-integration.test
CMD ["/bin/sh", "-c", "gadget-integration.test -test.v -integration"]
