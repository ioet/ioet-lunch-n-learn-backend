package config

import (
	"encoding/json"
)

func FormatCredsJSON(cfg *Config) ([]byte, error) {
	creds := map[string]string{
		"type":                        cfg.FirebaseType,
		"project_id":                  cfg.FirebaseProjectID,
		"private_key_id":              cfg.FirebasePrivateKeyID,
		"private_key":                 cfg.FirebasePrivateKey,
		"client_email":                cfg.FirebaseClientEmail,
		"client_id":                   cfg.FirebaseClientID,
		"auth_uri":                    cfg.FirebaseAuthURI,
		"token_uri":                   cfg.FirebaseTokenURI,
		"auth_provider_x509_cert_url": cfg.FirebaseAuthProviderX509CertURL,
		"client_x509_cert_url":        cfg.FirebaseClientX509CertURL,
		"universe_domain":             cfg.FirebaseUniverseDomain,
	}

	return json.Marshal(creds)
}
