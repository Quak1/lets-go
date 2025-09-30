# Snippetbox

## Generating a self-signed TLS certificate
```bash
mkdir tls
cd tls

# Find out where go is installed
which go

go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```
