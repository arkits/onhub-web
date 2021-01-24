set -e

echo "==> Building Go Binary for $GOARCH"
cd ..
    go mod download
    packr2 build .
cd -
