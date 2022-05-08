#!/usr/bin/env bash

set -euo pipefail

cd config/rsa-key
openssl genrsa -out oauth-private.key 4096
openssl rsa -in oauth-private.key -pubout -out oauth-public.key