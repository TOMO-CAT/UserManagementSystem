Version := $(shell git describe --tags --abbrev=0)  # 查看当前分支最近的 tag
Commit := $(shell git rev-parse --short HEAD)
Branch := $(shell git branch | grep \* | cut -d " " -f2 | tr " " "_")
BuildTime := $(shell date +"%Y-%m-%d-%H:%M")
Builder := $(shell git config user.email)
GoVersion := $(shell go version | cut -d " " -f3- | tr " " "_")
LastCommitTime := $(shell git show -s --format=%ci | tr " " "_")
LastCommitAuthor := $(shell git log -1 --pretty=format:'%ae')

BUILD_INFO = -X github.com/TOMO-CAT/UserManagementSystem/pkg/util/app.Version=${Version} \
 	-X github.com/TOMO-CAT/UserManagementSystem/pkg/util/app.Commit=${Commit} \
	-X github.com/TOMO-CAT/UserManagementSystem/pkg/util/app.Branch=${Branch} \
	-X github.com/TOMO-CAT/UserManagementSystem/pkg/util/app.BuildTime=${BuildTime} \
	-X github.com/TOMO-CAT/UserManagementSystem/pkg/util/app.Builder=${Builder} \
	-X github.com/TOMO-CAT/UserManagementSystem/pkg/util/app.GoVersion=${GoVersion} \
	-X github.com/TOMO-CAT/UserManagementSystem/pkg/util/app.LastCommitTime=${LastCommitTime} \
	-X github.com/TOMO-CAT/UserManagementSystem/pkg/util/app.LastCommitAuthor=${LastCommitAuthor}


# 根据 *.proto 文件生成对应的 *.pb.go
.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/config/logger_config.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/service/ums_service.proto

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build:
	go build -ldflags "-w ${BUILD_INFO}" -o bin/ums_server main/ums_server.go

.PHONY: start
start:
	@./bin/ums_server --control start

.PHONY: stop
stop:
	@./bin/ums_server --control stop

.PHONY: info
info:
	@./bin/ums_server --info

.PHONY: restart
restart:
	@./bin/ums_server --control restart

.PHONY: run
run:
	@./bin/ums_server
