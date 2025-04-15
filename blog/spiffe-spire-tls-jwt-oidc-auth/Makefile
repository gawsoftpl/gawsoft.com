build:
	go build "./gencerts.go"
	chmod +x ./gencerts

tls:
	./gencerts ./services/spire-server ./services/spire-agent

clean:
	find . -type f \( -iname "*.crt" -o -iname "*.pem" -o -iname "*.key" -o -iname "*.token" \) -exec rm -f {} \;
	docker-compose down