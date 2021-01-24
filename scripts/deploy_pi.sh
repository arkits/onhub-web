set -e

source deploy_constants.sh

./kill_remote.sh

cd ..
echo ">>> SCP'ing to $USERNAME@$HOSTNAME:$WORK_DIR"
scp onhub-web $USERNAME@$HOSTNAME:$WORK_DIR
cd -

echo ">>> Starting New Process"
ssh -q $USERNAME@$HOSTNAME <<EOF
 cd $WORK_DIR
 ./onhub-web > service.log 2>&1 & 
  ls -la
EOF
