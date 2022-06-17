ROOT=$PWD

SRC_BIN=$ROOT/cmd/conv
SRC_PLUGINS=$ROOT/plugins

TARGET_BIN=$ROOT/bin
TARGET_PLUGINS=$TARGET_BIN/plugins

echo "mkdir -p $TARGET_PLUGINS"
mkdir -p $TARGET_PLUGINS

echo "cd $SRC_BIN"
cd $SRC_BIN

echo "go build -o $TARGET_BIN/conv main.go"
go build -o $TARGET_BIN/conv main.go

echo "cd $SRC_PLUGINS/ip"
cd $SRC_PLUGINS/ip

echo "go build -buildmode=plugin -o $TARGET_PLUGINS/ip.so ip.go"
go build -buildmode=plugin -o $TARGET_PLUGINS/ip.so ip.go

echo "cd $SRC_PLUGINS/bits"
cd $SRC_PLUGINS/bits

echo "go build -buildmode=plugin -o $TARGET_PLUGINS/bits.so bits.go"
go build -buildmode=plugin -o $TARGET_PLUGINS/bits.so bits.go