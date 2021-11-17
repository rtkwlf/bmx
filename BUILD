load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test", "nogo")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@rules_pkg//:pkg.bzl", "pkg_tar")

pkg_tar(
    name = "package",
    srcs = ["//cmd/bmx"],
    mode = "0755",
    package_dir = "/",
)

# gazelle:prefix github.com/rtkwlf/bmx
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=bazel/go/deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_library(
    name = "go_default_library",
    srcs = [
        "console.go",
        "credential-process.go",
        "login.go",
        "print.go",
        "print_unix.go",
        "print_windows.go",
        "write.go",
    ],
    importpath = "github.com/rtkwlf/bmx",
    visibility = ["//visibility:public"],
    deps = [
        "//console:go_default_library",
        "//saml/identityProviders:go_default_library",
        "//saml/identityProviders/okta:go_default_library",
        "//saml/serviceProviders:go_default_library",
        "//saml/serviceProviders/aws:go_default_library",
        "@com_github_aws_aws_sdk_go//service/sts:go_default_library",
        "@in_gopkg_ini_v1//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "console_test.go",
        "print_test.go",
    ],
    deps = [
        ":go_default_library",
        "//mocks:go_default_library",
        "//saml/identityProviders/okta:go_default_library",
        "//saml/serviceProviders/aws/mocks:go_default_library",
    ],
)

nogo(
    name = "nogo",
    config = "nogo_config.json",
    vet = True,
    visibility = ["//visibility:public"],
)
