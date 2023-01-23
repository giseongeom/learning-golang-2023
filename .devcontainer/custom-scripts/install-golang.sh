#!/usr/bin/env bash
set -e

if [ "$(id -u)" -ne 0 ]; then
    echo -e 'Script must be run as root. Use sudo, su, or add "USER root" to your Dockerfile before running this script.'
    exit 1
fi

# go1.19.linux-amd64.tar.gz
export GO_DOWNLOAD_FILE=$(curl -L -Ss https://go.dev/dl/\?mode\=json | jq -r '.[0].files[] | select(.arch == "amd64") | select(.os == "linux") | .filename')
export GO_DOWNLOAD_URL="https://go.dev/dl/${GO_DOWNLOAD_FILE}"

curl -Ss -L -o /tmp/golang.tgz ${GO_DOWNLOAD_URL} \
    && tar zxf /tmp/golang.tgz -C /usr/local \
    && /usr/local/go/bin/go version \
    && echo "golang Installation is Done!" && echo ''

