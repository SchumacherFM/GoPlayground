package GoPlayground_test

import "testing"
import (
	"crypto/rand"
	"crypto/rsa"

	"gopkg.in/square/go-jose.v1"
)

func Test_JSON_JOSE(t *testing.T) {

	// Generate a public/private key pair to use for this example. The library
	// also provides two utility functions (LoadPublicKey and LoadPrivateKey)
	// that can be used to load keys from PEM/DER-encoded data.
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Instantiate an encrypter using RSA-OAEP with AES128-GCM. An error would
	// indicate that the selected algorithm(s) are not currently supported.
	publicKey := &privateKey.PublicKey
	encrypter, err := jose.NewEncrypter(jose.RSA_OAEP, jose.A128GCM, publicKey)
	if err != nil {
		panic(err)
	}

	// Encrypt a sample plaintext. Calling the encrypter returns an encrypted
	// JWE object, which can then be serialized for output afterwards. An error
	// would indicate a problem in an underlying cryptographic primitive.
	var plaintext = []byte("Lorem ipsum dolor sit amet")
	object, err := encrypter.Encrypt(plaintext)
	if err != nil {
		panic(err)
	}

	// Serialize the encrypted object using the full serialization format.
	// Alternatively you can also use the compact format here by calling
	// object.CompactSerialize() instead.
	serialized := object.FullSerialize()
	t.Log("serialized", serialized)

	serializedCompact, err := object.CompactSerialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("serializedCompact", serializedCompact)

	// Parse the serialized, encrypted JWE object. An error would indicate that
	// the given input did not represent a valid message.
	object, err = jose.ParseEncrypted(serialized)
	if err != nil {
		panic(err)
	}

	// Now we can decrypt and get back our original plaintext. An error here
	// would indicate the the message failed to decrypt, e.g. because the auth
	// tag was broken or the message was tampered with.
	decrypted, err := object.Decrypt(privateKey)
	if err != nil {
		panic(err)
	}

	t.Log("decrypted", string(decrypted))

}
