#!/bin/sh
set -eu

if [ "${FORCE_SAST_FAILURE:-0}" = "1" ]; then
    printf '%s\n' 'SAST stub: FAIL (forced by FORCE_SAST_FAILURE=1)' >&2
    exit 1
fi

printf '%s\n' 'SAST stub: PASS (no simulated source-code finding)'
