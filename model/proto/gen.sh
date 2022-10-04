CUR_DIR=$(pwd)
echo $CUR_DIR

META_DIR=$CUR_DIR"/meta/"
echo $META_DIR

RPC_DIR=$CUR_DIR"/rpc/"
echo RPC_DIR


DST_DIR=$CUR_DIR
echo $DST_DIR

META="*.proto"

protoc --go_out=${DST_DIR} --proto_path=${META_DIR} ${META_DIR}${META}
#protoc -I=$META_DIR --go_out=plugins=grpc:${DST_DIR} ${META_DIR}${META}