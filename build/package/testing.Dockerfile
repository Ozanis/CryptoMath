FROM alpine:latest
ENV GOLANG_VERSION="1.14.2"
RUN apk add --update --no-cache wget tar xz
RUN wget -O go.tar.gz https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go.tar.gz && rm go.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin
RUN apk del --purge tar xz wget
WORKDIR /pkg
ADD .pkg/ .
RUN ls -la
RUN go -version