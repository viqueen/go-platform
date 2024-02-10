harness:
	docker compose up -d
	go run _harness/server/cmd/main.go

schema:
	cd domains/_schema && buf generate