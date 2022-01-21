FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV http_proxy=https://27.115.15.12:18123/
ENV export https_proxy=https://27.115.15.12:18123/

WORKDIR /go/src/gin-vue-admin
COPY . .

RUN go mod tidy && go build -o server .

FROM alpine:3.10
LABEL MAINTAINER="jun.huang3@madhouse-inc.com"
WORKDIR /go/src/gin-vue-admin
COPY --from=0 /go/src/gin-vue-admin/server ./
COPY --from=0 /go/src/gin-vue-admin/config.yaml ./
COPY --from=0 /go/src/gin-vue-admin/resource ./resource
COPY --from=0 /go/src/gin-vue-admin/kubeconfig ./kubeconfig

EXPOSE 8888

ENTRYPOINT ./server
