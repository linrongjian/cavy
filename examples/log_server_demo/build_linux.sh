CURRENT_DIR=$(cd $(dirname $0); pwd)
cd $CURRENT_DIR

GOOS=linux CGO_ENABLED=0 GOARCH=amd64 \
go build \
-o log_server_demo_linux \
-tags "release" \
-ldflags "-X main.Developer=zl" .