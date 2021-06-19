`openssl genrsa -aes256 -out private_key.pem 2048 # will ask for a pass phrase`

`openssl rsa -in private_key.pem -outform PEM -pubout -out public_key.pem`

`openssl rsautl -encrypt -pubin -inkey public_key.pem -in plaintext.txt -out encrypted.txt`

`openssl rsautl -decrypt -passin pass:${PASS_PHRASE} -inkey private_key.pem -in encrypted.txt -out plaintext.txt`


1. https://stackoverflow.com/questions/57331931/what-does-this-command-do-openssl-genrsa-aes256-out-example-key-2048