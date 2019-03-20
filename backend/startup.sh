#!/usr/bin/env bash
# https://medium.com/@gchudnov/trapping-signals-in-docker-containers-7a57fdda7d86
set -x

pid=0

# SIGTERM-handler
term_handler() {
  if [ $pid -ne 0 ]; then
    kill -SIGTERM "$pid"
    wait "$pid"
  fi
  exit 143; # 128 + 15 -- SIGTERM
}

# setup handlers
# on callback, kill the last background process, which is `tail -f /dev/null` and execute the specified handler
trap 'kill ${!}; term_handler' SIGTERM
trap 'kill ${!}; term_handler' SIGKILL

# run application
./cetea --port=$PORT &
pid="$!"

# wait forever
while true
do
  tail -f /dev/null & wait ${!}
done