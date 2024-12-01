#!/bin/bash
set -euo pipefail
ROOT_DIR=$(realpath "$(dirname "$0")")

if [[ $# != 1 ]]; then
  echo Please provide a day to run.
  echo usage: "$0" DAY
  exit 1
fi

if [[ ! "$1" =~ ^([1-9]|1[0-9]|2[0-5])$ ]]; then
  echo Not a valid day: "$1"
  exit 1
fi

if [[ -z "${AOC_SESSION-""}" ]]; then
  echo \$AOC_SESSION not set
  exit 1
fi

l0day=$(printf "%02d" "$1")

DAY_DIR=$ROOT_DIR/$l0day

mkdir -p $DAY_DIR
curl -H "Cookie: session=$AOC_SESSION" https://adventofcode.com/2024/day/$1/input > $DAY_DIR/input.txt
cp template.go $DAY_DIR/main.go
