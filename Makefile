ADS_VERSION=v8
PROTO_ROOT_DIR=googleapis/
PROTO_SRC_DIR=/google/ads/googleads/$(ADS_VERSION)/**/*.proto
# PROTO_OUT_DIR=$$GOPATH/src/github.com/kritzware/google-ads-go/
PROTO_OUT_DIR=${HOME}/workspace/gjg/google-ads-go/
PKG_PATH=paths=source_relative
PROTOC_GO_ARGS=--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative

ENTRY=main.go
BIN=gads

build:
	go build -o $$GOPATH/bin/$(BIN) $(ENTRY)

run:
	go run $(ENTRY)

run-debug:
	GODEBUG=http2debug=2 GRPC_GO_LOG_SEVERITY_LEVEL=info GRPC_GO_LOG_VERBOSITY_LEVEL=2 go run $(ENTRY)

test:
	go test -v -cover ./...

.SILENT protos: clean-protos clean-gen-protos
	echo "converting protos for version $(ADS_VERSION)"
	echo $(PROTO_ROOT_DIR)$(PROTO_SRC_DIR)
	@for file in $(shell ls -R googleapis/google/ads/googleads/v8/**/*.proto); do \
		echo "converting proto $$(basename $$file)"; \
		protoc --proto_path=$(PROTO_ROOT_DIR) $(PROTOC_GO_ARGS) $$file; \
	done; \
	sh ./fix-package-paths.sh; \
	rm -rf google/
	@echo "built proto files to $$(basename $(PROTO_OUT_DIR))"

clean-protos:
	rm -rf common/
	rm -rf enums/
	rm -rf errors/
	rm -rf resources/
	rm -rf services/

clean-gen-protos:
	rm -rf google/

clone-googleapis:
	cd $(PROTO_ROOT_DIR)
	git submodule update --init --recursive

update-googleapis:
	cd $(PROTO_ROOT_DIR)
	git submodule update --recursive --remote

.PHONY: protos clone-googleapis update-googleapis
