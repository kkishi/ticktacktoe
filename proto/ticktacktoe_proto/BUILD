load("@io_bazel_rules_go//proto:go_proto_library.bzl", "go_proto_library")

# Workaround for https://github.com/bazelbuild/rules_go/issues/280
# Ideally we want ticktacktoe.proto at //proto/ticktacktoe.proto, but gazelle
# doesn't handle it very well.
go_proto_library(
    name = "go_default_library",
    srcs = ["ticktacktoe.proto"],
    has_services = 1,
    visibility = ["//visibility:public"],
)
