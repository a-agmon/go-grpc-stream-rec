BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
RM_F_CMD = rm -f
RM_RF_CMD = ${RM_F_CMD} -r
SERVER_BIN = ${SERVER_DIR}
CLIENT_BIN = ${CLIENT_DIR}

.PHONY: compile clean


compile:
	@echo "Executing $@"
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. ${PROTO_DIR}/*.proto
	go build -o ${BIN_DIR}/${SERVER_BIN} ./${SERVER_DIR}
	go build -o ${BIN_DIR}/${CLIENT_BIN} ./${CLIENT_DIR}

clean:
	${RM_F_CMD} ${PROTO_DIR}/*.pb.go

