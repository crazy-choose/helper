CUR_DIR=$(pwd)
echo $CUR_DIR

META_DIR=$CUR_DIR"/meta/"
echo $META_DIR


DST_DIR=$CUR_DIR
echo $DST_DIR

META="*.proto"

protoc -I=$META_DIR --go_out=plugins=grpc:${DST_DIR} ${META_DIR}${META}