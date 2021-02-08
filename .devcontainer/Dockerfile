FROM gcr.io/cloud-marketplace-containers/google/bazel
RUN apt-get update
RUN apt-get install -y zsh && \
    sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

ARG BAZEL_BUILDTOOLS_VERSION
ARG BAZEL_BUILDTOOLS_URL="https://github.com/bazelbuild/buildtools/releases/download/"
RUN curl -sSL "${BAZEL_BUILDTOOLS_URL}/${BAZEL_BUILDTOOLS_VERSION}/buildozer" > /usr/local/bin/buildozer && \
    curl -sSL "${BAZEL_BUILDTOOLS_URL}/${BAZEL_BUILDTOOLS_VERSION}/buildifier" >  /usr/local/bin/buildifier && \
    chmod +x /usr/local/bin/buildozer && \
    chmod +x /usr/local/bin/buildifier 