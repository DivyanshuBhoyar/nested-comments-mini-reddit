migrate:
	migrate -source file://database/migrations \
		-database postgres://postgres:change_me@127.0.0.1:5432/nested_comments?sslmode=disable up

rollback:
	migrate -source file://database/migrations \
		-database postgres://postgres:change_me@127.0.0.1:5432/nested_comments?sslmode=disable down

drop:
	migrate -source file://database/migrations \
		-database postgres://postgres:change_me@127.0.0.1:5432/nested_comments?sslmode=disable drop


migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir database/migrations $$name

sqlc:
	sqlc generate