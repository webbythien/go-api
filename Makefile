server:
	go run main.go server
worker:
	go run main.go worker core_api_queue:test
abigen:
	abigen --abi ./pkg/abis/RefCode.json --pkg contract --type RefCode --out ./app/models/contract/refcode.go