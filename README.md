An attempt to build a tick tack toe game supporting game plays over the internet.

To run:

1.   Clone the repository to a location according to your $GOPATH. E.g.,

     /home/ec2-user/go/src/github.com/kkishi/ticktacktoe

1.   In ticktacktoe/, run gazelle to generate BUILD files.

     `$ bazel run //:gazelle`

1.   Build & run the binary.

     `$ bazel run server/main:main`
