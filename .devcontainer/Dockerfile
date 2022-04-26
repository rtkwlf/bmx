ARG VARIANT="1.17"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

RUN apt-get update

ARG BAZEL_BUILDTOOLS_VERSION
ARG BAZEL_BAZELISK_VERSION
ARG BAZEL_BAZELISK_URL="https://github.com/bazelbuild/bazelisk/releases/download"
ARG BAZEL_BUILDTOOLS_URL="https://github.com/bazelbuild/buildtools/releases/download"
RUN curl -sSL "${BAZEL_BUILDTOOLS_URL}/${BAZEL_BUILDTOOLS_VERSION}/buildozer-linux-$(dpkg --print-architecture)" > /usr/local/bin/buildozer && \
    curl -sSL "${BAZEL_BUILDTOOLS_URL}/${BAZEL_BUILDTOOLS_VERSION}/buildifier-linux-$(dpkg --print-architecture)" >  /usr/local/bin/buildifier && \
    curl -sSL "${BAZEL_BAZELISK_URL}/${BAZEL_BAZELISK_VERSION}/bazelisk-linux-$(dpkg --print-architecture)" > /usr/local/bin/bazel && \
    chmod +x /usr/local/bin/bazel && \
    chmod +x /usr/local/bin/buildozer && \
    chmod +x /usr/local/bin/buildifier

# Bazel & Python compatibility
RUN ln /usr/bin/python3 /usr/bin/python