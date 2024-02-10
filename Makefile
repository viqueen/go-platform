harness:
	docker compose up -d
	go run _harness/server/cmd/main.go

webapp:
	yarn workspace @todo/webapp start

schema:
	cd domains/_schema && buf generate