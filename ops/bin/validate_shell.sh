#!/bin/sh

# shellcheck disable=SC2044

set -euo

for filename in $(find ./ops/* -name '*.sh');do
  echo "validating ${filename}"
  shellcheck "${filename}" || exit $?
done

exit 0