# Currently Dockerfile.arm64 is linked to this so if any arch specific
# pieces go in then we need to split these files
FROM golang:alpine as build
COPY . /source
WORKDIR /source
RUN mv /var/run /var/run.real && mkdir /var/run &&\
    apk update && apk upgrade &&\
    apk add git make &&\
    rm -rf /var/run && mv /var/run.real /var/run &&\
    make build &&\
    mkdir -p /release/etc &&\
    cp -r /etc/ssl /release/etc &&\
    grep nobody /etc/passwd >/release/etc/passwd &&\
    mv bin/hopper /release

FROM scratch
COPY --from=build /release /
USER nobody
ENTRYPOINT [ "/hopper" ]