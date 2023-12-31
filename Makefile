.PHONY: format
format:
	@#go install golang.org/x/tools/cmd/goimports@latest
	@#goimports -local "github.com/hodlgap/captive-portal" -w .
	@go install -v github.com/incu6us/goimports-reviser/v3@latest
	@goimports-reviser -rm-unused \
		-company-prefixes 'github.com/hodlgap' \
		-excludes 'db' \
		-project-name 'github.com/hodlgap/captive-portal' \
		-format \
		./...
	@gofmt -s -w .
	@go mod tidy

.PHONY: lint
lint:
	@golangci-lint run -v ./...

.PHONY: test
test:
	@go install github.com/rakyll/gotest@latest
	@gotest -race -cover -v ./...

.PHONY: update
update:
	@go get -u all
	@go mod tidy

.PHONY: generate
generate:
	@go install go.uber.org/mock/mockgen@latest
	@go generate ./...

.PHONY: models
models:
	@go install github.com/volatiletech/sqlboiler/v4@latest
	@go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
	sqlboiler psql -c sqlboiler.toml \
		--templates ${GOPATH}/pkg/mod/github.com/volatiletech/sqlboiler/*/templates/main \
		--templates ${GOPATH}/pkg/mod/github.com/volatiletech/sqlboiler/*/drivers/sqlboiler-psql/driver/override/main \
		--templates db/templates

.PHONY: dump-db
dump-db:
	@# This dumps your local postgres to db/schema.sql
	@PGPASSWORD=example pg_dump --no-owner --schema-only --no-privileges --host=localhost --username=postgres --dbname=captive-portal > db/schema.sql
	echo "db/schema.sql"

.PHONY: restore-db
restore-db:
	# This restores your local postgres to db/schema.sql
	#PGPASSWORD=example psql --host=localhost --username=postgres --dbname=captive-portal -c "drop database if exists \"captive-portal\";"
	@PGPASSWORD=example psql -h localhost -U postgres -d captive-portal < db/schema.sql
	echo "Completed"

