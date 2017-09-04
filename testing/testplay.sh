#!/bin/bash

set -e

bazel build //...

./bazel-bin/server/main/main -v=2 &
SERVER="$!"

sleep 1

echo "0 0 0 1 0 2\n" | ./bazel-bin/client/main/main --name='Tick' &
CLIENT="$!"

echo "1 0 1 1\n" | ./bazel-bin/client/main/main --name='Tack' &
CLIENT="$CLIENT $!"

wait $CLIENT

echo "0 2 1 1 1 2 2 1 0 0\n" | ./bazel-bin/client/main/main --name='Tick' &
CLIENT="$!"

echo "0 1 1 0 2 0 2 2\n" | ./bazel-bin/client/main/main --name='Tack' &
CLIENT="$CLIENT $!"

wait $CLIENT

kill $SERVER
