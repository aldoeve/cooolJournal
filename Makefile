build:
	@echo "Compiling frontend..."
	(cd frontend/ ; npm run build)

run:
	(cd backend/server ; go run .)

clean:
	@echo "Removing compiled files..."
	rm -rf frontend/dist