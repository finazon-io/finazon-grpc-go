PATH_THIS:=$(realpath $(dir $(lastword ${MAKEFILE_LIST})))

PROTOC_VERSION:=21.12
PROTOC_ASSETS:=${PATH_THIS}/download/protoc_${PROTOC_VERSION}
PROTOC:=${PROTOC_ASSETS}/bin/protoc

# https://protobuf.dev/reference/go/go-generated/#package
GO_PACKAGE:="github.com/finazon/finazon-grpc-go"

${PROTOC}:
	@mkdir -p ${PROTOC_ASSETS}
	@wget -q https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -P ${PATH_THIS}/download
	@unzip -q ${PATH_THIS}/download/protoc-${PROTOC_VERSION}-linux-x86_64.zip -d ${PROTOC_ASSETS}
	@rm ${PATH_THIS}/download/protoc-${PROTOC_VERSION}-linux-x86_64.zip

.PHONY: generate
generate: ${PROTOC}
	@[ "${VERSION}" ] || ( echo "VERSION is not set"; exit 1 )
	@$(eval mappingFinazon:=$(shell cd ${PATH_THIS}/proto && find . -iname "*.proto" -exec echo --go_opt=M{}=${GO_PACKAGE} \; | awk '{sub(/M.\//,"M")}1'))
	@$(eval mappingFinazonGrpc:=$(shell cd ${PATH_THIS}/proto && find . -iname "*.proto" -exec echo --go-grpc_opt=M{}=${GO_PACKAGE} \; | awk '{sub(/M.\//,"M")}1'))
	@rm -rf ${PATH_THIS}/pb/*
	@${PROTOC} ${PATH_THIS}/proto/*.proto \
	  --proto_path=${PATH_THIS}/proto \
	  ${mappingFinazon} \
	  --go_out=${PATH_THIS}/pb \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=${PATH_THIS}/pb \
	  ${mappingFinazonGrpc} \
	  --go-grpc_opt=paths=source_relative
	 @${PATH_THIS}/generate_extra_code.bash ${VERSION}

.PHONY: bump_version
bump_version:
	@[ "${VERSION}" ] || ( echo "VERSION is not set"; exit 1 )

clean:
	@rm -rf ${PATH_THIS}/download

.PHONY: publish
publish:
	@echo -n ""
