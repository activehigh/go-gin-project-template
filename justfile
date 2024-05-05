
servide_dir := justfile_directory()

vet:
    cd {{servide_dir}} && go vet ./...

fmt:
    cd {{servide_dir}} && go fmt ./...

tidy:
    cd {{servide_dir}} && go mod tidy

generate:
    docker compose up --build go-generate --remove-orphans

test:
    docker compose up --build service-test --remove-orphans

run: tidy vet fmt
    docker compose up --build service --remove-orphans

go-gen:
    go generate ./...

