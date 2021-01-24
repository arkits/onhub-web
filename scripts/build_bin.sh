set -e

# CGO is required for go-sqlite
export CGO_ENABLED=1

echo "==> Building Go Binary for $GOARCH"
cd ..
    go mod download
    packr2 build .
cd -
