FROM golang:1.14.6-alpine3.11 AS build
WORKDIR /app
COPY main.go /app/
COPY pkg /app/pkg
COPY go.* /app/
ARG AppBuildVersion="0.0.0"
ARG AppBuildHash="???"
ARG AppBuildDate="???"
ARG AppBuildMode="development"
ARG GOARCH=""
RUN CGO_ENABLED=0 GOOS=linux GOARCH="$GOARCH" go build \
  -a \
  -installsuffix cgo \
  -ldflags "-extldflags '-static' -s -w \
    -X gitlab.com/safesurfer/minio-file-server/pkg/common.AppBuildVersion=$AppBuildVersion \
    -X gitlab.com/safesurfer/minio-file-server/pkg/common.AppBuildHash=$AppBuildHash \
    -X gitlab.com/safesurfer/minio-file-server/pkg/common.AppBuildDate=$AppBuildDate \
    -X gitlab.com/safesurfer/minio-file-server/pkg/common.AppBuildMode=$AppBuildMode" \
  -o minio-file-server \
  main.go

FROM alpine:3.11 as extras
RUN apk add tzdata ca-certificates
RUN adduser -D user

FROM scratch
WORKDIR /app
ENV PATH=/app
COPY --from=build /app/minio-file-server .
COPY --from=extras /etc/passwd /etc/passwd
COPY --from=extras /etc/group /etc/group
COPY --from=extras /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=extras /etc/ssl /etc/ssl
EXPOSE 8080
USER user
ENTRYPOINT ["/app/minio-file-server"]
