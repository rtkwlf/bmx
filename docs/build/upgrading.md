# Upgrading of Tools & Dependencies

The following document describes the process for upgrading dependencies in the project.

## Upgrading bazel

The version of bazel used in the codebase is managed by the `.bazelversion` file located at the root of the project. Changing the version specified here will impact which version is selected by [bazelisk](https://github.com/bazelbuild/bazelisk) for continuous integration and local development environments.

## Upgrading golang

The golang version is defined in `WORKSPACE` when registering the version of the golang toolchain. Changing this version will upgrade the go version of the project. When doing this, you should also consider upgrading the dependencies to ensure everything is working as expected.

## Upgrading golang dependencies

You can upgrade the go dependencies by running the following commands:

```bash
go get -u ./...
go get -t -u ./...
go mod vendor
go mod tidy
```

You can then update the dependencies tracked in bazel by running the following command:

```bash
bazel run //:gazelle -- update-repos -build_external=vendored -from_file=go.mod -to_macro=bazel/go/deps.bzl%go_dependencies
```

This will re-generate the `deps.bzl` file that defines all of the golang dependencies used in the project.