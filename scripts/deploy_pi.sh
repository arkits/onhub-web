set -e

USERNAME="pi"
HOSTNAME="192.168.86.112"
WORK_DIR="/home/pi/software/onhub-web"
BIN_NAME="onhub-web"

echo ">>> Killing Running Process"
ssh -q $USERNAME@$HOSTNAME <<EOF
  PID=`ps -eaf | grep onhub-web | grep -v grep | awk '{print $2}'`

  if [[ "" !=  "$PID" ]]; then
    echo "killing $PID"
    kill -9 $PID
  fi
EOF

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
