#!/bin/bash

set -e

if [ -z "${PINGAIC_TF_TENANT_ENV_FQDN}" ]; then
    echo "PINGAIC_TF_TENANT_ENV_FQDN is not set. Cannot get access token."
    exit 1
fi

if [ -z "${SERVICE_ACCOUNT_ID}" ]; then
    echo "SERVICE_ACCOUNT_ID is not set. Cannot get access token."
    exit 1
fi

if [ -z "${SERVICE_ACCOUNT_PRIVATE_KEY}" ]; then
    echo "SERVICE_ACCOUNT_PRIVATE_KEY is not set. Cannot get access token."
    exit 1
fi

AUD="https://${PINGAIC_TF_TENANT_ENV_FQDN}:443/am/oauth2/access_token"

EXP=$(($(date -u +%s) + 899))

JTI=$(openssl rand -base64 16)

echo -n "{
    \"iss\":\"${SERVICE_ACCOUNT_ID}\",
    \"sub\":\"${SERVICE_ACCOUNT_ID}\",
    \"aud\":\"${AUD}\",
    \"exp\":${EXP},
    \"jti\":\"${JTI}\"
}" > payload.json

echo "${SERVICE_ACCOUNT_PRIVATE_KEY}" > service_account_private_key.jwk

jose jws sig -I payload.json -k service_account_private_key.jwk -s '{"alg":"RS256"}' -c -o jwt.txt

rm service_account_private_key.jwk

ACCESS_TOKEN=$(curl \
--request POST ${AUD} \
--data "client_id=service-account" \
--data "grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer" \
--data "assertion=$(< jwt.txt)" \
--data "scope=fr:idc:certificate:* fr:idc:content-security-policy:* fr:idc:cookie-domain:* fr:idc:custom-domain:* fr:idc:esv:* fr:idc:promotion:* fr:idc:sso-cookie:*" | jq ".access_token")

echo "${ACCESS_TOKEN}" > access_token.txt

FILESIZE=$(stat -c%s access_token.txt)
echo "Size of token = $FILESIZE bytes."

echo "Access token created"
