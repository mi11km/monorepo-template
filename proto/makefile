.PHONY: generate
generate:
	buf dep update
	buf generate --template schema/rpc/buf.go.gen.yaml
	buf generate --template schema/rpc/buf.ts.gen.yaml
