set -e

echo "==> Building Go Binary for $GOARCH"
cd ..
    go mod download

    COMMIT_ID=$(git rev-parse --verify HEAD)

    packr2 build -ldflags "-X main.version=$COMMIT_ID" -v .
cd -
