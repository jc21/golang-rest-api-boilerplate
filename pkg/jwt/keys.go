package jwt

import (
	"boilerplate/pkg/config"
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

// GetPrivateKey will load the key from config package and return a usable object
// It should only load from file once per program execution
func GetPrivateKey() (*rsa.PrivateKey, error) {
	if privateKey == nil {
		var blankKey *rsa.PrivateKey

		filename := config.Env.Key.Private
		if filename == "" {
			return blankKey, errors.New("Could not get Private Key filename, check environment variables")
		}

		var err error
		privateKey, err = LoadPemPrivateKey(filename)
		if err != nil {
			return blankKey, err
		}
	}

	pub, pubErr := GetPublicKey()
	if pubErr != nil {
		return privateKey, pubErr
	}

	privateKey.PublicKey = *pub

	return privateKey, pubErr
}

// GetPublicKey will load the key from config package and return a usable object
// It should only load from file once per program execution
func GetPublicKey() (*rsa.PublicKey, error) {
	if publicKey == nil {
		var blankKey *rsa.PublicKey

		filename := config.Env.Key.Public
		if filename == "" {
			return blankKey, errors.New("Could not get Public Key filename, check environment variables")
		}

		var err error
		publicKey, err = LoadPemPublicKey(filename)
		if err != nil {
			return blankKey, err
		}
	}

	return publicKey, nil
}

// LoadPemPrivateKey reads a key from a PEM encoded file and returns a private key
func LoadPemPrivateKey(fileName string) (*rsa.PrivateKey, error) {
	var key *rsa.PrivateKey
	privateKeyFile, err := os.Open(fileName)

	if err != nil {
		return key, err
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	if err != nil {
		return key, err
	}
	data, _ := pem.Decode(pembytes)
	privateKeyFile.Close()
	key, err = x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		return key, err
	}
	return key, nil
}

// LoadPemPublicKey reads a key from a PEM encoded file and returns a public key
func LoadPemPublicKey(fileName string) (*rsa.PublicKey, error) {
	var key *rsa.PublicKey

	publicKeyFile, err := os.Open(fileName)
	if err != nil {
		return key, err
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	size := pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		return key, err
	}
	data, _ := pem.Decode(pembytes)
	publicKeyFile.Close()
	publicKeyFileImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		return key, err
	}

	var ok bool
	if key, ok = publicKeyFileImported.(*rsa.PublicKey); !ok {
		return key, fmt.Errorf("Unable to parse Public key")
	}
	return key, nil
}
