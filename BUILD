load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/jrbeverly/bmx
gazelle(name = "gazelle")

go_binary(
    name = "bmx",
    embed = ["//cmd/bmx:go_default_library"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "console.go",
        "print.go",
        "print_unix.go",
        "print_windows.go",
        "write.go",
    ],
    importpath = "github.com/jrbeverly/bmx",
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
    srcs = ["print_test.go"],
    embed = [":go_default_library"],
    deps = ["//mocks:go_default_library"],
)
