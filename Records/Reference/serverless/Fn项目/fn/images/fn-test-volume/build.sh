set -e
docker build --build-arg HTTPS_PROXY --build-arg HTTP_PROXY -t fnproject/fn-test-volume:latest .
