load("@org_pubref_rules_protobuf//protobuf:rules.bzl", "proto_language")

# Add pb_imports and pb_inputs to the closure language definition to comiple
# proto files that import google/api/annotations.proto for gRPC Gateway. This
# depends on external repositories imported by grpc_gateway_proto_repositories.
# This rule is based on the following files:
# * https://github.com/pubref/rules_protobuf/blob/50a641aa3203aebadc5d8d3cd0d93f580a334d0b/closure/BUILD
# * https://github.com/pubref/rules_protobuf/blob/f20e112465643daa58f23199dbf163b7193640f9/grpc_gateway/BUILD
proto_language(
    name = "closure_grpc",
    output_to_library = True,
    pb_file_extensions = [".js"],
    pb_imports = [
        "external/com_google_protobuf/src/",
        "external/com_github_grpc_ecosystem_grpc_gateway_googleapis/",
    ],
    pb_inputs = [
        "@com_google_protobuf//:well_known_protos",
        "@com_github_grpc_ecosystem_grpc_gateway_googleapis//:annotations_protos",
    ],
    pb_options = [
        "import_style=closure",
        "error_on_name_conflict",
        "binary",
    ],
    pb_plugin_name = "js",
)

load("@org_pubref_rules_protobuf//closure:rules.bzl", "closure_proto_library")

closure_proto_library(
    name = "ticktacktoe_proto",
    langs = [":closure_grpc"],
    protos = ["//proto/ticktacktoe_proto:ticktacktoe.proto"],
    visibility = ["//visibility:public"],
)
