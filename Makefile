# modelごとのprotoファイルを作成する e.g. make protoc s=user
protoc:
	protoc -I ./protofiles ${s}.proto --go_out=plugins=grpc:./services/api-gateway/
	protoc -I ./protofiles ${s}.proto --go_out=plugins=grpc:./services/${s}