build:
	@echo "Compiling frontend..."
	(cd frontend/ ; npm run build)

run:
	@echo "Attempting to run server..."
	(cd backend/server ; go run .)

clean:
	@echo "Removing compiled files..."
	rm -rf frontend/dist
	(cd backend/ ; go mod tidy)

prod:
	rm -rf frontend/dist
	(cd frontend/ ; npm run build)
	(cd backend/ ; go mod tidy; cd server/ ; rm -f server; mod build)