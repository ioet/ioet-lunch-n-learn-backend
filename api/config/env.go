package config

import "sync"

type Config struct {
	AppPort                         string `mapstructure:"APP_PORT"`
	FirebaseAppURL                  string `mapstructure:"FIREBASE_APP_URL"`
	FirebaseType                    string `mapstructure:"FIREBASE_TYPE"`
	FirebaseProjectID               string `mapstructure:"FIREBASE_PROJECT_ID"`
	FirebasePrivateKeyID            string `mapstructure:"FIREBASE_PRIVATE_KEY_ID"`
	FirebasePrivateKey              string `mapstructure:"FIREBASE_PRIVATE_KEY"`
	FirebaseClientEmail             string `mapstructure:"FIREBASE_CLIENT_EMAIL"`
	FirebaseClientID                string `mapstructure:"FIREBASE_CLIENT_ID"`
	FirebaseAuthURI                 string `mapstructure:"FIREBASE_AUTH_URI"`
	FirebaseTokenURI                string `mapstructure:"FIREBASE_TOKEN_URI"`
	FirebaseAuthProviderX509CertURL string `mapstructure:"FIREBASE_AUTH_PROVIDER_X509_CERT_URL"`
	FirebaseClientX509CertURL       string `mapstructure:"FIREBASE_CLIENT_X509_CERT_URL"`
	FirebaseUniverseDomain          string `mapstructure:"FIREBASE_UNIVERSE_DOMAIN"`
}

var (
	config *Config
	once   sync.Once
	envs   = []string{
		"APP_PORT",
		"FIREBASE_APP_URL",
		"FIREBASE_TYPE",
		"FIREBASE_PROJECT_ID",
		"FIREBASE_PRIVATE_KEY_ID",
		"FIREBASE_PRIVATE_KEY",
		"FIREBASE_CLIENT_EMAIL",
		"FIREBASE_CLIENT_ID",
		"FIREBASE_AUTH_URI",
		"FIREBASE_TOKEN_URI",
		"FIREBASE_AUTH_PROVIDER_X509_CERT_URL",
		"FIREBASE_CLIENT_X509_CERT_URL",
		"FIREBASE_UNIVERSE_DOMAIN",
	}
)
