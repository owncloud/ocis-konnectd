package flagset

import (
	"github.com/micro/cli"
	"github.com/owncloud/ocis-konnectd/pkg/config"
)

// RootWithConfig applies cfg to the root flagset
func RootWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "config-file",
			Value:       "",
			Usage:       "Path to config file",
			EnvVar:      "KONNECTD_CONFIG_FILE",
			Destination: &cfg.File,
		},
		&cli.StringFlag{
			Name:        "log-level",
			Value:       "info",
			Usage:       "Set logging level",
			EnvVar:      "KONNECTD_LOG_LEVEL",
			Destination: &cfg.Log.Level,
		},
		&cli.BoolTFlag{
			Name:        "log-pretty",
			Usage:       "Enable pretty logging",
			EnvVar:      "KONNECTD_LOG_PRETTY",
			Destination: &cfg.Log.Pretty,
		},
		&cli.BoolTFlag{
			Name:        "log-color",
			Usage:       "Enable colored logging",
			EnvVar:      "KONNECTD_LOG_COLOR",
			Destination: &cfg.Log.Color,
		},
	}
}

// HealthWithConfig applies cfg to the root flagset
func HealthWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:9134",
			Usage:       "Address to debug endpoint",
			EnvVar:      "KONNECTD_DEBUG_ADDR",
			Destination: &cfg.Debug.Addr,
		},
	}
}

// ServerWithConfig applies cfg to the root flagset
func ServerWithConfig(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "tracing-enabled",
			Usage:       "Enable sending traces",
			EnvVar:      "KONNECTD_TRACING_ENABLED",
			Destination: &cfg.Tracing.Enabled,
		},
		&cli.StringFlag{
			Name:        "tracing-type",
			Value:       "jaeger",
			Usage:       "Tracing backend type",
			EnvVar:      "KONNECTD_TRACING_TYPE",
			Destination: &cfg.Tracing.Type,
		},
		&cli.StringFlag{
			Name:        "tracing-endpoint",
			Value:       "",
			Usage:       "Endpoint for the agent",
			EnvVar:      "KONNECTD_TRACING_ENDPOINT",
			Destination: &cfg.Tracing.Endpoint,
		},
		&cli.StringFlag{
			Name:        "tracing-collector",
			Value:       "",
			Usage:       "Endpoint for the collector",
			EnvVar:      "KONNECTD_TRACING_COLLECTOR",
			Destination: &cfg.Tracing.Collector,
		},
		&cli.StringFlag{
			Name:        "tracing-service",
			Value:       "konnectd",
			Usage:       "Service name for tracing",
			EnvVar:      "KONNECTD_TRACING_SERVICE",
			Destination: &cfg.Tracing.Service,
		},
		&cli.StringFlag{
			Name:        "debug-addr",
			Value:       "0.0.0.0:9134",
			Usage:       "Address to bind debug server",
			EnvVar:      "KONNECTD_DEBUG_ADDR",
			Destination: &cfg.Debug.Addr,
		},
		&cli.StringFlag{
			Name:        "debug-token",
			Value:       "",
			Usage:       "Token to grant metrics access",
			EnvVar:      "KONNECTD_DEBUG_TOKEN",
			Destination: &cfg.Debug.Token,
		},
		&cli.BoolFlag{
			Name:        "debug-pprof",
			Usage:       "Enable pprof debugging",
			EnvVar:      "KONNECTD_DEBUG_PPROF",
			Destination: &cfg.Debug.Pprof,
		},
		&cli.BoolFlag{
			Name:        "debug-zpages",
			Usage:       "Enable zpages debugging",
			EnvVar:      "KONNECTD_DEBUG_ZPAGES",
			Destination: &cfg.Debug.Zpages,
		},
		&cli.StringFlag{
			Name:        "http-addr",
			Value:       "0.0.0.0:9130",
			Usage:       "Address to bind http server",
			EnvVar:      "KONNECTD_HTTP_ADDR",
			Destination: &cfg.HTTP.Addr,
		},
		&cli.StringFlag{
			Name:        "http-root",
			Value:       "/",
			Usage:       "Root path of http server",
			EnvVar:      "KONNECTD_HTTP_ROOT",
			Destination: &cfg.HTTP.Root,
		},
		&cli.StringFlag{
			Name:        "http-namespace",
			Value:       "com.owncloud.web",
			Usage:       "Set the base namespace for service discovery",
			EnvVar:      "KONNECTD_HTTP_NAMESPACE",
			Destination: &cfg.HTTP.Namespace,
		},
		&cli.StringFlag{
			Name:        "identity-manager",
			Value:       "ldap",
			Usage:       "Identity manager (one of ldap,kc,cookie,dummy)",
			EnvVar:      "KONNECTD_IDENTITY_MANAGER",
			Destination: &cfg.Konnectd.IdentityManager,
		},
		&cli.StringFlag{
			Name:        "transport-tls-cert",
			Value:       "",
			Usage:       "Certificate file for transport encryption",
			EnvVar:      "KONNECTD_TRANSPORT_TLS_CERT",
			Destination: &cfg.HTTP.TLSCert,
		},
		&cli.StringFlag{
			Name:        "transport- tls-key",
			Value:       "",
			Usage:       "Secret file for transport encryption",
			EnvVar:      "KONNECTD_TRANSPORT_TLS_KEY",
			Destination: &cfg.HTTP.TLSKey,
		},
		&cli.StringFlag{
			Name:        "iss",
			Usage:       "OIDC issuer URL",
			EnvVar:      "KONNECTD_ISS",
			Value:       "https://127.0.0.1:9130",
			Destination: &cfg.Konnectd.ISS,
		},
		&cli.StringSliceFlag{
			Name:   "signing-private-key",
			Usage:  "Full path to PEM encoded private key file (must match the --signing-method algorithm)",
			EnvVar: "KONNECTD_SIGNING_PRIVATE_KEY",
			Value:  nil,
		},
		&cli.StringFlag{
			Name:        "signing-kid",
			Usage:       "Value of kid field to use in created tokens (uniquely identifying the signing-private-key)",
			EnvVar:      "KONNECTD_SIGNING_KID",
			Value:       "",
			Destination: &cfg.Konnectd.SigningKid,
		},
		&cli.StringFlag{
			Name:        "validation-keys-path",
			Usage:       "Full path to a folder containg PEM encoded private or public key files used for token validaton (file name without extension is used as kid)",
			EnvVar:      "KONNECTD_VALIDATION_KEYS_PATH",
			Value:       "",
			Destination: &cfg.Konnectd.ValidationKeysPath,
		},
		&cli.StringFlag{
			Name:        "encryption-secret",
			Usage:       "Full path to a file containing a %d bytes secret key",
			EnvVar:      "KONNECTD_ENCRYPTION_SECRET",
			Value:       "",
			Destination: &cfg.Konnectd.EncryptionSecretFile,
		},
		&cli.StringFlag{
			Name:        "signing-method",
			Usage:       "JWT default signing method",
			EnvVar:      "KONNECTD_SIGNING_METHOD",
			Value:       "PS256",
			Destination: &cfg.Konnectd.SigningMethod,
		},
		&cli.StringFlag{
			Name:        "uri-base-path",
			Usage:       "Custom base path for URI endpoints",
			EnvVar:      "KONNECTD_URI_BASE_PATH",
			Value:       "",
			Destination: &cfg.Konnectd.URIBasePath,
		},
		&cli.StringFlag{
			Name:        "sign-in-uri",
			Usage:       "Custom redirection URI to sign-in form",
			EnvVar:      "KONNECTD_SIGN_IN_URI",
			Value:       "",
			Destination: &cfg.Konnectd.SignInUri,
		},
		&cli.StringFlag{
			Name:        "signed-out-uri",
			Usage:       "Custom redirection URI to signed-out goodbye page",
			EnvVar:      "KONNECTD_SIGN_OUT_URI",
			Value:       "",
			Destination: &cfg.Konnectd.SignedOutUri,
		},
		&cli.StringFlag{
			Name:        "authorization-endpoint-uri",
			Usage:       "Custom authorization endpoint URI",
			EnvVar:      "KONNECTD_ENDPOINT_URI",
			Value:       "",
			Destination: &cfg.Konnectd.AuthorizationEndpointURI,
		},
		&cli.StringFlag{
			Name:        "endsession-endpoint-uri",
			Usage:       "Custom endsession endpoint URI",
			EnvVar:      "KONNECTD_ENDSESSION_ENDPOINT_URI",
			Value:       "",
			Destination: &cfg.Konnectd.EndsessionEndpointURI,
		},
		&cli.StringFlag{
			Name:        "asset-path",
			Value:       "",
			Usage:       "Path to custom assets",
			EnvVar:      "KONNECTD_ASSET_PATH",
			Destination: &cfg.Asset.Path,
		},
		&cli.StringFlag{
			Name:        "identifier-client-path",
			Usage:       "Path to the identifier web client base folder",
			EnvVar:      "KONNECTD_IDENTIFIER_CLIENT_PATH",
			Value:       "./assets/identifier",
			Destination: &cfg.Konnectd.IdentifierClientPath,
		},
		&cli.StringFlag{
			Name:        "identifier-registration-conf",
			Usage:       "Path to a identifier-registration.yaml configuration file",
			EnvVar:      "KONNECTD_IDENTIFIER_REGISTRATION_CONF",
			Value:       "./identifier-registration.yaml",
			Destination: &cfg.Konnectd.IdentifierRegistrationConf,
		},
		&cli.StringFlag{
			Name:        "identifier-scopes-conf",
			Usage:       "Path to a scopes.yaml configuration file",
			EnvVar:      "KONNECTD_IDENTIFIER_SCOPES_CONF",
			Value:       "",
			Destination: &cfg.Konnectd.IdentifierScopesConf,
		},
		&cli.BoolFlag{
			Name:        "insecure",
			Usage:       "Disable TLS certificate and hostname validation",
			EnvVar:      "KONNECTD_INSECURE",
			Destination: &cfg.Konnectd.Insecure,
		},
		&cli.StringSliceFlag{
			Name:   "trusted-proxy",
			Usage:  "Trusted proxy IP or IP network (can be used multiple times)",
			EnvVar: "KONNECTD_TRUSTED_PROXY",
			Value:  nil,
		},
		&cli.StringSliceFlag{
			Name:   "allow-scope",
			Usage:  "Allow OAuth 2 scope (can be used multiple times, if not set default scopes are allowed)",
			EnvVar: "KONNECTD_ALLOW_SCOPE",
			Value:  nil,
		},
		&cli.BoolFlag{
			Name:        "allow-client-guests",
			Usage:       "Allow sign in of client controlled guest users",
			EnvVar:      "KONNECTD_ALLOW_CLIENT_GUESTS",
			Destination: &cfg.Konnectd.AllowClientGuests,
		},
		&cli.BoolFlag{
			Name:        "allow-dynamic-client-registration",
			Usage:       "Allow dynamic OAuth2 client registration",
			EnvVar:      "KONNECTD_ALLOW_DYNAMIC_CLIENT_REGISTRATION",
			Destination: &cfg.Konnectd.AllowDynamicClientRegistration,
		},
	}
}
