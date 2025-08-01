up:
	docker-compose up -d

down:
	docker-compose down

migrate:
	docker-compose exec app ./main --migrate

logs:
	docker-compose logs -f app

test:
	go test -v ./...