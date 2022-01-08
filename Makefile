GENQLIENT = $(shell go env GOPATH)/bin/genqlient

.PHONY: generate
generate:
	$(GENQLIENT) ./genqlient.yaml

.PHONY: install-tools
install-tools:
	go install github.com/Khan/genqlient@latest

.PHONY: download-schema
download-schema:
	curl --create-dirs --output internal/github/api/schema.docs.graphql https://docs.github.com/public/schema.docs.graphql
