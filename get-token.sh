#!/bin/sh

cd "${GITHUB_WORKSPACE}" || exit 2

echo ""

# If container is running inside a proxy, reroute proxy to container host IP
if [ ! -z $HTTP_PROXY ]; then
	host=$(ip route show | awk '/default/ {print $3}')
	HTTP_PROXY="${HTTP_PROXY/127.0.0.1/${host}}"
	HTTPS_PROXY="${HTTPS_PROXY/127.0.0.1/${host}}"
fi

export GO111MODULE=on

cd /issuer
token=$(go run . -appid=${INPUT_APPID} -pem="${INPUT_PRIVATEKEY}" -repository="${GITHUB_REPOSITORY}")
status=${?}

if [ ${status} -ne 0 ]; then
	echo ${token}
	exit ${status}
fi

echo "::set-output name=token::${token}"
