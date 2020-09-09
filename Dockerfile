# syntax = docker/dockerfile:1-experimental

FROM --platform=${BUILDPLATFORM} golang:1.13.15-alpine3.12 AS build
ENV GO_ENABLED=0
ARG TARGETOS
ARG TARGETARCH
ARG LD_FLAGS

RUN mkdir -p $GOPATH/src/code.cloudfoundry.org/cli
WORKDIR $GOPATH/src/code.cloudfoundry.org/cli
COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/cf .
# RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags=${LD_FLAGS} -o /out/cf .

FROM scratch AS bin-unix
COPY --from=build /out/cf /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/example /example.exe

FROM bin-${TARGETOS} AS bin