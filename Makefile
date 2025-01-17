##@ Testing

test:
	godotenv -f .env go test ./... -race -v