# /bin/bash
# make_pi.sh
# Builds a complete binary for Raspberry Pi 4 running Ubuntu Server (arm64)

set -e

./build_web.sh

export GOOS=linux
export GOARCH=arm64

# For Raspberry Pi 4 running Raspbian
# export GOARCH=arm
# export GOARM=5

# export CGO_ENABLED=1
./build_bin.sh

exit
