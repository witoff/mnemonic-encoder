package main

import (
	"bytes"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"log"

	bip32 "github.com/tyler-smith/go-bip32"
	bip39 "github.com/tyler-smith/go-bip39"
)

// ecPrivateKey is an ASN.1 encoded EC key defined here:
// https://tools.ietf.org/html/rfc5915
type ecPrivateKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")
	masterKey, _ := bip32.NewMasterKey(seed)

	// Derive child keys
	// TODO: Use the correct derivation path
	childKey, err := masterKey.NewChildKey(bip32.FirstHardenedChild)
	check(err)
	publicKey := childKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Private key: ", childKey)
	fmt.Println("Public key: ", publicKey)

	// Pack into ASN1 EC Object
	oidPrivateKeySecp256k1 := asn1.ObjectIdentifier{1, 3, 132, 0, 10}
	ec := ecPrivateKey{}
	// Per RFC
	ec.Version = 1
	// TODO: Verify correct format.
	ec.PrivateKey = childKey.Key
	// Per RFC
	ec.NamedCurveOID = oidPrivateKeySecp256k1
	// TODO: Verify bitsring is parsed accoding to asn1.parseBitString
	ec.PublicKey = asn1.BitString{
		Bytes:     publicKey.Key,
		BitLength: len(publicKey.Key),
	}
	asn1Bytes, err := asn1.Marshal(ec)
	check(err)

	// Pack into an ASN1 pkcs8PrivateKey object
	// TODO

	// Pack into a PEM file
	// TODO: Do we want `PRIVATE KEY` or `EC PRIVATE KEY`
	pemBlock := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: asn1Bytes,
	}

	buf := bytes.NewBufferString("")
	if err = pem.Encode(buf, &pemBlock); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}

// REFERENCES
// Importing a wrapped key:
// - https://developers.yubico.com/YubiHSM2/Commands/Import_Wrapped.html
// Import with the shell
// - https://developers.yubico.com/YubiHSM2/Component_Reference/yubihsm-shell/
