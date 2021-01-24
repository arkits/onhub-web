set -e

PID=$(ps -eaf | grep "onhub-web" | grep -v grep | awk '{print $2}')

if [[ "" != "$PID" ]]; then
    echo "Killing $PID"
    kill -9 $PID
fi
