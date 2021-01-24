set -e

echo "==> Building Web"
cd ../web
    # rm -rf node_modules
    yarn install
    yarn build
cd -
