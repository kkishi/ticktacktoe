load("@io_bazel_rules_closure//closure:defs.bzl", "closure_js_binary", "closure_js_library")

closure_js_binary(
    name = "ttt",
    deps = [
        ":client",
        ":debug_console",
    ],
)

closure_js_library(
    name = "client",
    srcs = ["client.js"],
    deps = [
        ":board",
        ":phaser_extern",
        "//proto:ticktacktoe_proto",
        "@io_bazel_rules_closure//closure/library",
    ],
)

closure_js_library(
    name = "board",
    srcs = ["board.js"],
    deps = [
        ":phaser_extern",
        "//proto:ticktacktoe_proto",
    ],
)

closure_js_library(
    name = "debug_console",
    srcs = ["debug_console.js"],
)

closure_js_library(
    name = "phaser_extern",
    srcs = ["phaser_extern.js"],
)
