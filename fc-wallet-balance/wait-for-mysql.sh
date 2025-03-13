#!/bin/sh
set -e

host="$1"
shift
cmd="$@"

until mysqladmin ping -h "$host" --silent; do
  echo "Aguardando MySQL..."
  sleep 3
done

exec $cmd
