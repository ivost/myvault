#!/bin/bash
set -e

/server

echo CMD $1

while true;
  do sleep 5;
  echo working...;
  done

exec "$@"
