export PATH :=$(PATH):$(GOPATH)/bin

APP=$(shell basename "$(PWD)")
POD=$(shell kubectl get pod -l project=${APP} -o jsonpath="{.items[0].metadata.name}")
PROTOC ?= protoc

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

tools: ##install dependencies
	@echo 'Installing protoc-gen-go...'
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	@echo 'Success'
	@echo 'Installing protoc-gen-go-grpc...'
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	@echo 'Success'

start: ## start one-node local cluster
	minikube start
stop: ## stop one-node local cluster
	minikube stop

build-image: ## build image and restart app
	minikube image build -t ${APP}:latest -f build/Dockerfile .
	kubectl apply -f deployments/dev
	kubectl rollout restart deployment/${APP}

run: ##  run instance
	minikube service ${APP}

serve: start build-image run ## build and run application

grpc-gen: grpc-get-proto ## autogenerate grpc code
	$(PROTOC) --go_out=internal/interface/controllers/gRPC/ --go-grpc_out=internal/interface/controllers/grpc/ -I$(shell go list -f '{{ .Dir }}' -m gitlab.doslab.ru/sell-and-buy/sb-proto) delivery/delivery.proto
	$(PROTOC) --go_out=internal/interface/controllers/gRPC/ --go-grpc_out=internal/interface/controllers/grpc/ -I$(shell go list -f '{{ .Dir }}' -m gitlab.doslab.ru/sell-and-buy/sb-proto) module/module.proto
	$(PROTOC) --go_out=internal/interface/controllers/gRPC/ --go-grpc_out=internal/interface/controllers/grpc/ -I$(shell go list -f '{{ .Dir }}' -m gitlab.doslab.ru/sell-and-buy/sb-proto) health/health_check.proto
	go mod tidy

grpc-get-proto: ## download repository with proto files
	GOSUMDB=off go get gitlab.doslab.ru/sell-and-buy/sb-proto

exec: ## enter to app container
	kubectl exec -it ${POD} -- /bin/sh

app-logs: ## show app logs in console
	kubectl logs po/${POD} -f