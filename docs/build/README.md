# Developing with BMX

BMX is built with bazel. 

You can built the tool using the following:

## Setting up developer environment

Project tooling is managed by bazel, and once you have installed bazelisk you should be able

The project supports [devcontainers]() for quick configuration of an environment assuming you have a working environment.

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
