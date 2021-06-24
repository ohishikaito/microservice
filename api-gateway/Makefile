# modelごとのprotoファイルを作成する e.g. make protoc s=user
protoc:
	protoc -I ./protofiles ${s}.proto --go_out=plugins=grpc:./
	protoc -I ./protofiles ${s}.proto --go_out=plugins=grpc:../${s}