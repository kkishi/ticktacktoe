# From https://github.com/bazelbuild/rules_go
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.5.3",
)

load("@io_bazel_rules_go//go:def.bzl", "go_repositories", "go_repository")

go_repositories()

# From https://github.com/pubref/rules_protobuf
# Use proto dependencies provided from pubref for using grpc gateway.
git_repository(
    name = "org_pubref_rules_protobuf",
    remote = "https://github.com/pubref/rules_protobuf",
    tag = "v0.8.1",
)

load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_repositories")

go_proto_repositories()

load("@org_pubref_rules_protobuf//grpc_gateway:rules.bzl", "grpc_gateway_proto_repositories")

grpc_gateway_proto_repositories()

# For gomock
go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    tag = "v1.0.0",
)

# For gRPC Websocket Proxy
go_repository(
    name = "com_github_tmc_grpc_websocket_proxy",
    commit = "89b8d40f7ca833297db804fcb3be53a76d01c238",
    importpath = "github.com/tmc/grpc-websocket-proxy",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    tag = "v1.0.3",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    tag = "v1.2.0",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "9aade4d3a3b7e6d876cd3823ad20ec45fc035402",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "81e90905daefcd6fd217b62423c0908922eadb30",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "81e90905daefcd6fd217b62423c0908922eadb30",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "81e90905daefcd6fd217b62423c0908922eadb30",
    importpath = "golang.org/x/crypto",
)

# From https://github.com/bazelbuild/rules_closure
http_archive(
    name = "io_bazel_rules_closure",
    strip_prefix = "rules_closure-0.4.2",
    sha256 = "25f5399f18d8bf9ce435f85c6bbf671ec4820bc4396b3022cc5dc4bc66303609",
    urls = [
        "http://mirror.bazel.build/github.com/bazelbuild/rules_closure/archive/0.4.2.tar.gz",
        "https://github.com/bazelbuild/rules_closure/archive/0.4.2.tar.gz",
    ],
)

load("@io_bazel_rules_closure//closure:defs.bzl", "closure_repositories")

closure_repositories()
