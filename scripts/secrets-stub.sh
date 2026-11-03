#!/bin/sh
set -eu

if [ "${FORCE_SECRETS_FAILURE:-0}" = "1" ]; then
    printf '%s\n' 'Secrets stub: FAIL (forced by FORCE_SECRETS_FAILURE=1)' >&2
    exit 1
fi

printf '%s\n' 'Secrets stub: PASS (no simulated secret finding)'
