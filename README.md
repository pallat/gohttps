https://github.com/denji/golang-tls

# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650


# client
## ref: http://javamemento.blogspot.com/2015/10/using-curl-with-ssl-cert-chain.html
1. open browser -> click "Not Secure" -> drag cer to desktop
2. openssl x509 -inform DES -in localhost.cer -out local.pem -text
3. go run client.go -im local.pem