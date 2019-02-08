all: tfctl

# Run tests
test: fmt vet
	go test -v ./pkg/... ./cmd/...

# Build tfctl binary
tfctl: fmt vet
	go build -o bin/tfctl github.com/stpabhi/tfctl

# Download vendor dependencies
deps:
	dep ensure -v

# Run against the configured Kubernetes cluster in ~/.kube/config
run: tfctl
	bin/tfctl

# Run go fmt against code
fmt:
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet:
	go vet ./pkg/... ./cmd/...