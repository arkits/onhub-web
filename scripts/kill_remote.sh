set -e

source deploy_constants.sh

echo ">>> Killing Running Process on $USERNAME@$HOSTNAME"
ssh $USERNAME@$HOSTNAME 'bash -s' <kill_service.sh
