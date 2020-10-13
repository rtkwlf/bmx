load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test", "nogo")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/jrbeverly/bmx
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = [
        "console.go",
        "credential-process.go",
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

nogo(
    name = "nogo",
    config = "nogo_config.json",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_x_tools//go/analysis/passes/asmdecl:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/assign:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/atomic:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/bools:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/buildtag:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/composite:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/copylock:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/httpresponse:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/loopclosure:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/lostcancel:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/nilfunc:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/printf:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/shift:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/stdmethods:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/structtag:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/tests:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/unreachable:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/unsafeptr:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/unusedresult:go_tool_library",
    ],
)
