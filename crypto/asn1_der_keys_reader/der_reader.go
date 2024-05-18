package main

import (
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"os"
)

func main() {
	// RSA
	pwd, _ := os.Getwd()
	file, err := os.ReadFile(pwd + "/crypto/asn1_der_keys_reader/keys/public.der")
	if err != nil {
		return
	}

	data, err := x509.ParsePKIXPublicKey(file)
	if err != nil {
		return
	}

	key := data.(*rsa.PublicKey)
	_ = key

	// Private Key
	pwd, _ = os.Getwd()
	file, err = os.ReadFile(pwd + "/crypto/asn1_der_keys_reader/keys/private.der")
	if err != nil {
		return
	}

	data, err = x509.ParsePKCS8PrivateKey(file)
	if err != nil {
		return
	}

	rsaPrivateKey := data.(*rsa.PrivateKey)
	_ = rsaPrivateKey

	// ed25519
	pwd, _ = os.Getwd()
	file, err = os.ReadFile(pwd + "/crypto/asn1_der_keys_reader/keys/ed25519/public.der")
	if err != nil {
		return
	}

	data, err = x509.ParsePKIXPublicKey(file)
	if err != nil {
		return
	}

	edkey := data.(ed25519.PublicKey)
	_ = edkey

	pwd, _ = os.Getwd()
	file, err = os.ReadFile(pwd + "/crypto/asn1_der_keys_reader/keys/ed25519/private.der")
	if err != nil {
		return
	}

	data, err = x509.ParsePKCS8PrivateKey(file)
	if err != nil {
		return
	}

	edPrivatekey := data.(ed25519.PrivateKey)
	_ = edPrivatekey
}
