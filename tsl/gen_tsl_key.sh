#bin/bash

openssl req -x509 -newkey rsa:4096 -keyout  ./tsl/my-private-key.pem -out  ./tsl/my-public-key-cert.pem -days 365 -nodes -subj '/CN=localhost'
