set -e

source deploy_constants.sh

./kill_remote.sh
echo ""

cd ..
echo ">>> SCP'ing to $USERNAME@$HOSTNAME:$WORK_DIR"
scp onhub-web $USERNAME@$HOSTNAME:$WORK_DIR
cd -
echo ""

ssh $USERNAME@$HOSTNAME 'bash -s' <run_remote.sh
