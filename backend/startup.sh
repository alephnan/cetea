#!/usr/bin/env bash
# https://medium.com/@gchudnov/trapping-signals-in-docker-containers-7a57fdda7d86
set -x

pid=0

handle_interrupt() {
  if [ $pid -ne 0 ]; then
    kill -SIGINT "$pid"
    wait "$pid"
  fi
  exit 143; # 128 + 15 -- SIGTERM
}
handle_terminate() {
  if [ $pid -ne 0 ]; then
    kill -SIGTERM "$pid"
    wait "$pid"
  fi

  exit 143; # 128 + 15 -- SIGTERM
}

# setup handlers
# on callback, kill the last background process, which is `tail -f /dev/null` and execute the specified handler
trap 'kill ${!}; handle_interrupt' SIGINT
trap 'kill ${!}; handle_terminate' SIGTERM

# run application
./main --port=$PORT &
pid="$!"

# wait forever
while true
do
  tail -f /dev/null & wait ${!}
done