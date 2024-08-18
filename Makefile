GOOSE_BIN=goose
UNAME_S=$(shell uname -s)

.PHONY: logging
logging:
	@echo "UNAME_S: $(UNAME_S)"
	@echo "GOOSE_BIN: $(GOOSE_BIN)"
	
build:
	GOOS=linux GOARCH=amd64 go build -o dist/main_linux cmd/main.go

act: build
	@act pull_request -j 'review' \
	--eventpath github.input \
	--container-architecture linux/amd64 \
	--secret-file github.secrets \
	--pull=false \
	-W .github/workflows/ai-review.yml
