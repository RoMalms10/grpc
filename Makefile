.PHONY: generate clean

# Directory structure
PROTO_DIR = ./proto
OUT_DIR = ./meme

generate:
	@echo "Generating Go code from protobuf definitions..."
	@mkdir -p $(OUT_DIR)
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/**/*.proto
	@echo "Code generation complete"

clean:
	@echo "Cleaning generated files..."
	@rm -rf $(OUT_DIR)