##@ Testing

test:
	godotenv -f .env go test ./... -race -v

test-fmt:
	set -eu pipefail && godotenv -f .env go test -json -v ./... 2>&1 | tee /tmp/gotest.log | gotestfmt -hide empty-packages

##@Live reload
ifeq ($(OS),Windows_NT)
    AIR_CONFIG=.\\.air.windows.conf
else
    AIR_CONFIG=.air.linux.conf
endif

air:
	air -c $(AIR_CONFIG)
