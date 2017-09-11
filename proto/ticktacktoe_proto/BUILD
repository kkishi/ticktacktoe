# Workaround for https://github.com/bazelbuild/rules_go/issues/280
# Ideally we want ticktacktoe.proto at //proto/ticktacktoe.proto, but gazelle
# doesn't handle it very well.

package(default_visibility = ["//visibility:public"])

load("@org_pubref_rules_protobuf//grpc_gateway:rules.bzl", "grpc_gateway_proto_library")

grpc_gateway_proto_library(
    name = "go_default_library",
    protos = ["ticktacktoe.proto"],
    visibility = ["//visibility:public"],
)
