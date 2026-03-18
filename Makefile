start:
	make stop
	docker-compose up --remove-orphans

start-no-cache:
	docker-compose build --no-cache
	docker-compose up --remove-orphans

stop:
	docker-compose down

reset-db:
	docker exec -i db mysql -u demo -pdemo users < data_init.sql

generate-proto:
	make docker-build-protoc
	@docker run --volume "$(PWD)":/protoc_builder --workdir /protoc_builder \
	protoc_builder /bin/bash -c "go mod download && \
    	protoc \
    		--proto_path=./presentation/grpc \
    		--go_out=./presentation/grpc \
    		--go_opt=paths=source_relative  \
    		--go-grpc_out=./presentation/grpc \
    		--go-grpc_opt=paths=source_relative \
    		./presentation/grpc/*/*.proto && \
		go mod tidy"
run-tests:
	make docker-build-tests

docker-build-tests:
	@docker build \
		--tag tests_builder \
		-f tests.dockerfile .

docker-build-protoc:
	@docker build \
		--tag protoc_builder \
		-f protoc_builder.dockerfile .
