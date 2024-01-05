commit:
	./commit.sh

# e.g. make jwt-gen username=user
jwt-gen: username := $(or $(username), super.admin.one)
jwt-gen:
	go run -tags devtools . jwt-gen $(username)

# e.g. make jwt-key-gen flag="--force"
jwt-key-gen:
	go run -tags tools . jwt-key-gen $(flag)

# e.g. make migrate flag="--rollback --steps=17"
migrate:
	go run -tags tools . migrate $(flag)

migrate-fresh:
	go run -tags devtools . migrate-fresh

# e.g. make migrate-gen filename=create_table_name
migrate-gen:
	go run -tags devtools . migrate-gen -f $(filename)

# e.g. make seed name=table_name
seed:
	go run -tags devtools . seed $(name)

# e.g. make seed-production name=table_name
seed-production:
	go run -tags devtools . seed --production $(name)

sync-permission:
	go run -tags tools . sync-permission

# e.g. make test name=FuncTestName
test: flag := $(if $(name), -run $(name),)
test:
	go clean -testcache && go test -timeout 3h `go list ./... | grep -v tool$$`$(flag) -v > test.out

generate-docs:
	go run myapp/tool/swag fmt -d main.go,./delivery/api && go run myapp/tool/swag init -d ./,./delivery/api --outputTypes json,yaml
