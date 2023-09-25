start:
	@echo "Building and running app..."
	docker-compose up --build -d
	@echo "everything is ready!"

stop:
	@echo "Stopping all containers"
	docker-compose down
	@echo "everything is down now!"

log-app:
	@echo "logs in shop app..."
	docker-compose logs -f app

log-db:
	@echo "logs in postgres..."
	docker-compose logs -f postgres

restart: stop start