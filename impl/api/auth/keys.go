package auth

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
)

var (
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

// Load keys
func LoadKeys() error {
	if err := loadJwk(); err != nil {
		return err
	}
	if err := loadPrivateKey(); err != nil {
		return err
	}
	return nil
}

// Load JWK from file
func loadJwk() error {

	config, err := conf.GetConfig()
	if err != nil {
		return err
	}
	path := config.Api.JwkPath

	log.Infof("Loading JWK from %v", path)

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read JWK file from %v: %v", path, err)
	}

	keySet, err := jwk.Parse(data)
	if err != nil {
		return fmt.Errorf("unable to parse JWK file: %v", err)
	}

	if keySet.Len() == 0 {
		return errors.New("no keys found in JWK")
	}

	rawKey, ok := keySet.Get(0)
	if !ok {
		return errors.New("key at index 0 not found in JWK")
	}

	var pubKey rsa.PublicKey
	if err := rawKey.Raw(&pubKey); err != nil {
		return fmt.Errorf("unable to get raw JWK: %v", err)
	}
	publicKey = &pubKey
	return nil
}

// Load RSA private key from file
func loadPrivateKey() error {
	config, err := conf.GetConfig()
	if err != nil {
		return err
	}
	path := config.Api.RSAPrivatePath

	log.Infof("Loading RSA private key from %v", path)

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read RSA file from %v: %v", path, err)
	}

	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PRIVATE KEY" {
		return errors.New("failed to decode PEM block containing private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("unable to parse private key: %v", err)
	}

	var ok bool
	privateKey, ok = key.(*rsa.PrivateKey)
	if !ok {
		return errors.New("not an RSA private key")
	}

	return nil
}

// Get public key
func GetPublicKey() *rsa.PublicKey {
	return publicKey
}

// Get private key
func GetPrivateKey() *rsa.PrivateKey {
	return privateKey
}
