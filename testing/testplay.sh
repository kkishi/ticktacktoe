#!/bin/bash

set -e

bazel build //...

./bazel-bin/server/server -v=2 &
SERVER="$!"

sleep 1

CMD=./bazel-bin/client/cmd/cmd

echo "##################################################"

echo "0 0 0 1 0 2" | $CMD --name='Tick' &
CLIENT="$!"

sleep 1

echo "1 0 1 1" | $CMD --name='Tack' &
CLIENT="$CLIENT $!"

sleep 1

wait $CLIENT

echo "##################################################"

echo "0 2 1 1 1 2 2 1 0 0" | $CMD --name='Tick' &
CLIENT="$!"

sleep 1

echo "0 1 1 0 2 0 2 2" | $CMD --name='Tack' &
CLIENT="$CLIENT $!"

sleep 1

wait $CLIENT

kill $SERVER
