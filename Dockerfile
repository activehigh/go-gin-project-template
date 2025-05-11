FROM public.ecr.aws/docker/library/golang:1.24.0 AS builder

WORKDIR /app

# temp satisfy trunk
HEALTHCHECK NONE

# Install test tools
RUN go install go.uber.org/mock/mockgen@latest &&\
    go install github.com/onsi/ginkgo/v2/ginkgo@latest && \
    go install github.com/jstemmer/go-junit-report@latest && \
    go install github.com/axw/gocov/gocov@latest && \
    go install github.com/AlekSi/gocov-xml@latest


# Github fingerprint
# https://docs.github.com/en/github/authenticating-to-github/keeping-your-account-and-data-secure/githubs-ssh-key-fingerprints
ENV GITHUB_RSA_SHA256_FINGERPRINT=uNiVztksCsDhcc0u9e8BujQXVUpKZIDTMczCvj3tD2s\
    GOPRIVATE=github.com

RUN apt-get install --no-install-recommends openssh-client git -y

# add github ssh key fingerprint
# this is published by github and should be pinned and verified on client side
RUN mkdir -p /root/.ssh &&\
    ssh-keyscan -t rsa github.com               > /root/.ssh/known_hosts    &&\
    ssh-keygen -l -f /root/.ssh/known_hosts     > /tmp/githubRSAFingerPrint &&\
    if ! grep -q ${GITHUB_RSA_SHA256_FINGERPRINT} /tmp/githubRSAFingerPrint; \
    then\
    echo "github signature does not match";\
    exit 1;\
    fi

# Copy the package management manifest
COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY configs/ configs/
COPY pkg/ pkg/
COPY internal/ internal/

# Copy test payloads
COPY tests/data/ tests/data/

# generate mocks (run before copying tests as they can change without mocks changing)
RUN go generate -v ./...

# Build your app binary
FROM builder as intermittent-builder
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o runtime ./cmd/main.go

# ===========================================================

FROM public.ecr.aws/docker/library/alpine:3.21 AS runtime
WORKDIR /

RUN apk add --no-cache bash

COPY --from=intermittent-builder /app/runtime ./
USER 65532:65532

ENTRYPOINT ["/runtime"]
