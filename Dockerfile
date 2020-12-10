FROM golang:1.14-stretch AS BUILD

WORKDIR /src

COPY . .

RUN set -eux && cd tools && make update-tools

RUN go mod download

RUN set -eux && chmod +x ./tools/bin/swag && make swag

RUN set -eux && make bin && chmod +x ./echo-server

FROM gcr.io/distroless/base:latest

COPY --from=BUILD /src/echo-server /usr/bin/echo-server
COPY --from=BUILD src/static/* /static/

WORKDIR /

ENTRYPOINT [ "/usr/bin/echo-server" ]