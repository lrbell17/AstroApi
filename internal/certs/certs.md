# How to generate certs (local setup)

Generate CA:
```
openssl req -x509 -newkey rsa:4096 -days 8250 -nodes \
  -keyout ca.key -out ca.crt \
  -subj "/C=US/ST=CA/L=San Francisco/O=AstroCo/OU=Dev/CN=AstroAPI Dev CA"
```

Generate private key and CSR:
```
openssl req -newkey rsa:2048 -nodes -keyout localhost.key -out localhost.csr \
  -subj "/C=US/ST=CA/L=San Francisco/O=AstroCo/OU=Dev/CN=localhost"
```

Sign cert with local CA:
```
openssl x509 -req -in localhost.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
  -out localhost.crt -days 8250 -sha256
```

* Add localhost.crt and localhost.key to docker/conf/certs
* Use ca.crt wherever you are sending requests (e.g. Postman)