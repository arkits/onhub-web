set -e

USERNAME="pi"
HOSTNAME="192.168.86.112"
WORK_DIR="/home/pi/software/onhub-web"
BIN_NAME="onhub-web"

cd ..
    echo "SCP'ing to $USERNAME@$HOSTNAME:$WORK_DIR"
    scp onhub-web $USERNAME@$HOSTNAME:$WORK_DIR
cd -
