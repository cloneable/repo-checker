.PHONY: generate
generate:
	go run github.com/Khan/genqlient ./genqlient.yaml

.PHONY: download-schema
download-schema:
	curl --create-dirs --output internal/github/api/schema.docs.graphql https://docs.github.com/public/schema.docs.graphql
