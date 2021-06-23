pb-gen:
	protoc -I ./protofiles ${s}.proto --go_out=plugins=grpc:./pb
	protoc -I ./protofiles ${s}.proto --go_out=plugins=grpc:../user