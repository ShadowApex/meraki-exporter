ARG GO_VERSION=1.19.2
FROM golang:${GO_VERSION}-bullseye as build

WORKDIR /usr/src
COPY ./ ./

RUN make

# Runtime image
FROM ubuntu:20.04
RUN apt-get update && \
	apt-get install -y ca-certificates
COPY --from=build /usr/src/meraki-exporter /meraki-exporter

ENTRYPOINT [ "/meraki-exporter" ]
