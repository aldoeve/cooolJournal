build:
	@echo "Compiling frontend..."
	(cd frontend/ ; npm run build)

run:
	@echo "Attempting to run server..."
	(cd backend/server ; go run . ${f})

clean:
	@echo "Removing compiled files..."
	rm -rf frontend/dist