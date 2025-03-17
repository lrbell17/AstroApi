package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lrbell17/astroapi/impl/conf"
	log "github.com/sirupsen/logrus"
)

var publicKey *rsa.PublicKey

// Load JWK from file
func LoadJwk() error {

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

// Get JWK
func GetJwkPublicKey() *rsa.PublicKey {
	return publicKey
}
