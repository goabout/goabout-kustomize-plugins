FROM golang:stretch AS build
ENV KUSTOMIZE_VERSION=v0.0.0-20190430182652-0045d7b71604 \
    GO111MODULE=on
COPY . /src
RUN go get sigs.k8s.io/kustomize@$KUSTOMIZE_VERSION && \
    cd /go/pkg/mod/sigs.k8s.io/kustomize@$KUSTOMIZE_VERSION && \
    go build -buildmode plugin -o /build/sopsdotenv.so /src/sopsdotenv/main.go && \
    go build -buildmode plugin -o /build/sopsfiles.so /src/sopsfiles/main.go && \
    go install

FROM debian:stretch
ENV KUBECTL_VERSION=v1.14.0
MAINTAINER tech@goabout.com
COPY --from=build /go/bin/kustomize /usr/local/bin/
COPY --from=build /build /root/.config/kustomize/plugin/kvSources
RUN apt-get update && \
    apt-get install -y gnupg wget && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    wget https://storage.googleapis.com/kubernetes-release/release/$KUBECTL_VERSION/bin/linux/amd64/kubectl -O /usr/local/bin/kubectl && \
    chmod 755 /usr/local/bin/*
