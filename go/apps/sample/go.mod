module github.com/mi11km/monorepo-template/go/apps/sample

go 1.25.0

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.7-20250717185734-6c6e0d3c608e.1
	connectrpc.com/connect v1.18.1
	golang.org/x/net v0.43.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250811230008-5f3141c8851a
	google.golang.org/protobuf v1.36.7
)

require golang.org/x/text v0.28.0 // indirect
