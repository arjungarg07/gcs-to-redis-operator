#!/bin/sh

set -ex

if test -f "/root/.gcp/ds-tpu-sa.json"; then
    gcloud auth activate-service-account --key-file=/root/.gcp/ds-tpu-sa.json
    export GOOGLE_APPLICATION_CREDENTIALS=/root/.gcp/ds-tpu-sa.json
fi

if test -f "/root/.gcp/ds-composer-psc-sa.json"; then
    gcloud auth activate-service-account --key-file=/root/.gcp/ds-composer-psc-sa.json
    export GOOGLE_APPLICATION_CREDENTIALS=/root/.gcp/ds-composer-psc-sa.json
fi

gcloud config list account
gcloud config set project sharechat-production

echo "[Gcloud credentials set]"

./main "$@"
