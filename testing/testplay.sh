#!/bin/bash

set -e

bazel build //...

./bazel-bin/server/main/main -v=2 &
SERVER="$!"

sleep 1

echo "0 0 0 1 0 2" | ./bazel-bin/client/main/main --name='Tick' &
CLIENT="$!"

echo "1 0 1 1" | ./bazel-bin/client/main/main --name='Tack' &
CLIENT="$CLIENT $!"

wait $CLIENT
kill $SERVER
