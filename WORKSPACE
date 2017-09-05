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
  tag = "v0.7.2",
)
load("@org_pubref_rules_protobuf//go:rules.bzl", "go_proto_repositories")
go_proto_repositories()
load("@org_pubref_rules_protobuf//grpc_gateway:rules.bzl", "grpc_gateway_proto_repositories")
grpc_gateway_proto_repositories()

# From https://github.com/cgrushko/proto_library/blob/master/WORKSPACE
http_archive(
    name = "com_google_protobuf",
    urls = ["https://github.com/google/protobuf/archive/b4b0e304be5a68de3d0ee1af9b286f958750f5e4.zip"],
    strip_prefix = "protobuf-b4b0e304be5a68de3d0ee1af9b286f958750f5e4",
    sha256 = "ff771a662fb6bd4d3cc209bcccedef3e93980a49f71df1e987f6afa3bcdcba3a",
)

# For gomock
go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    tag = "v1.0.0",
)
