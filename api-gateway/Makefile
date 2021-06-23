pb-gen:
	protoc -I ./protofiles --go_out=plugins=grpc:./ user.proto
	protoc -I ./protofiles --go_out=plugins=grpc:../user user.proto
# api-gatewayでpbで共通定義しても干渉しない？