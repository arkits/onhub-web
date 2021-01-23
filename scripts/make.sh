set -e

echo "==> Building Web"
cd ../web
yarn build
cd -

echo "==> Building Go Binary"
cd ..
go mod download
packr2 build .
