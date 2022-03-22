#!/usr/bin/env bash

# generate 100 3Mb mock image files

# usage:
# cd images
# bash ./create-images.sh

set -euo pipefail
thisdir=$( cd "$(dirname "$0")" || exit 1; pwd )
for idx in $(seq 1 100); do
  dd if=/dev/urandom of="${thisdir}/image_${idx}.bin" bs=1024x1024 count=3
done
