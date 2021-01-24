set -e

echo ">>> Starting New Process"

cd /home/pi/software/onhub-web

./onhub-web >service.log 2>&1 &
