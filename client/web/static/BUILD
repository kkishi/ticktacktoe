load("@io_bazel_rules_closure//closure:defs.bzl", "closure_js_binary", "closure_js_library")

closure_js_binary(
    name = "ttt",
    deps = [
        ":client",
    ],
)

closure_js_library(
    name = "client",
    srcs = ["client.js"],
    deps = [
        "//proto:ticktacktoe_proto",
        ":phaser_extern",
        "@io_bazel_rules_closure//closure/library",
    ],
)

closure_js_library(
    name = "phaser_extern",
    srcs = ["phaser_extern.js"],
)
