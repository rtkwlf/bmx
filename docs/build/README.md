# Developing with BMX

BMX uses bazel for local development, managing the golang toolchain. GitHub Actions is responsible for building and packaging releases of the tool.

## Setting up developer environment

Project tooling is managed by bazel, and once you have installed [bazelisk](https://github.com/bazelbuild/bazelisk) you should be able to build the project by running:

```bash
bazel build //cmd/bmx:bmx
```

And BMX itself can be run with bazel by leveraging the `run` command:

```bash
bazel run //cmd/bmx:bmx -- login
```

The project supports [devcontainers](https://code.visualstudio.com/docs/remote/containers) to allow for an isolated development environment to avoid conflict with an existing bmx installation on the developer machine. This may not always be needed for development.

## Building the project

```bash
bazel build //cmd/bmx:bmx
```

You may also built and test everything in the entire repository by running the following:

```bash
bazel build //...
bazel test //...
```

### Cross Platform

You can built by specifying the platforms (`--platforms`) by running the following commands:

```bash
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //cmd/bmx:bmx
bazel build --platforms=@io_bazel_rules_go//go/toolchain:windows_amd64 //cmd/bmx:bmx
bazel build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 //cmd/bmx:bmx
```
