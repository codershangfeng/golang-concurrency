# Crypto

## RSA

## AES

## RSA & AES


## PEM file extention

[PEM, DER, CRT, and CER: X.509 Encodings and Conversions]

数字签名文件含有多种文件扩展名，比如：`.crt`、`.cer`、`.pem`或`.der`。

这些扩展对应着两类**X.509**编码方式:

- PEM (Base64 ASCII)
- DER (binary)

实际中，不能仅凭扩展名来确定哪种编码方式，通常还是需要用文本编辑器打开自己看看。

> PKCS#12, PKCS#7???

### PEM

PEM （originally "Privacy Enhanced Mail"）是X.509证书、Certificate Signing Request（CSR）和cryptographic keys的最常见格式。

PEM文件是个文本文件，含有一个或多个经Base64 ASCII编码的item，每个item有一个明文（纯文本）的header和footer，比如：`-----BEGIN RSA PRIVATE KEY-----`和`-----END RSA PRIVATE KEY-----`。

### DER

DER（Distinguished Encoding Rules）是X.509证书和私钥的二进制编码。

不同于PEM，DER编码的文件不包含明文（纯文本）。

DER文件比较常见于Java中。





[PEM, DER, CRT, and CER: X.509 Encodings and Conversions]:https://www.ssl.com/guide/pem-der-crt-and-cer-x-509-encodings-and-conversions/#:~:text=PEM%20(originally%20%E2%80%9CPrivacy%20Enhanced%20Mail,%2D%2D%2D%2D%2D%20).

