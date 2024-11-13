ARG GO_BASE_IMAGE
ARG DISTROLESS_BASE_IMAGE=gcr.io/distroless/static:nonroot
FROM --platform=${BUILDPLATFORM} $GO_BASE_IMAGE AS builder

WORKDIR /workspace
COPY . .
ARG TARGETARCH
RUN make build-scheduler GO_BUILD_ENV='CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH}'

FROM --platform=${BUILDPLATFORM} $DISTROLESS_BASE_IMAGE

WORKDIR /bin
COPY --from=builder /workspace/bin/kube-scheduler .
USER 65532:65532

ENTRYPOINT ["/bin/kube-scheduler"]