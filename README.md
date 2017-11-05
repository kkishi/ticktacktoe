A tick tack toe game which supports game plays over the internet.

To run:

1.  Clone the repository to a location according to your $GOPATH. E.g.,

    /home/ec2-user/go/src/github.com/kkishi/ticktacktoe

1.  cd into ticktacktoe/.

1.  (Optional) Run gazelle to generate BUILD files.

    `$ bazel run //:gazelle`

1.  Build.

    `$ bazel build ...`

1.  Run the server.

    `./bazel-bin/server/server`

1.  Run the gRPC gateway client.

    `./bazel-bin/client/web/main/main`

1.  Access http://localhost:8081 in your browser.
