package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config holds configuration for the project.
type Config struct {
	Env       string `env:"APP_ENV,default=development"`
	AppName   string `env:"APP_NAME,default=starter-api"`
	Port      Port
	HashID    HashID
	Google    Google
	Postgres  Postgres
	Redis     Redis
	SMTP      SMTP
	JWTConfig JWTConfig
	Image     Image
	OneSignal OneSignal
	Jaeger    Jaeger
	URL       URL
	MailGun   MailGun
}

// Port holds configuration for project's port.
type Port struct {
	APP string `env:"APP_PORT,default=8080"`
}

// Google holds configuration for the Google.
type Google struct {
	ProjectID          string `env:"GOOGLE_PROJECT_ID"`
	ServiceAccountFile string `env:"GOOGLE_SA"`
	StorageBucketName  string `env:"GOOGLE_STORAGE_BUCKET_NAME"`
	StorageEndpoint    string `env:"GOOGLE_STORAGE_ENDPOINT"`
}

// Image holds configuration for the Image.
type Image struct {
	Host string `env:"IMAGE_HOST"`
}

// URL holds configuration for the URL.
type URL struct {
	ForgotPasswordURL string `env:"FORGOT_PASSWORD_URL"`
}

// HashID holds configuration for HashID.
type HashID struct {
	Salt      string `env:"HASHID_SALT"`
	MinLength int    `env:"HASHID_MIN_LENGTH,default=10"`
}

// Postgres holds all configuration for PostgreSQL.
type Postgres struct {
	Host            string `env:"POSTGRES_HOST,default=localhost"`
	Port            string `env:"POSTGRES_PORT,default=5432"`
	User            string `env:"POSTGRES_USER,required"`
	Password        string `env:"POSTGRES_PASSWORD,required"`
	Name            string `env:"POSTGRES_NAME,required"`
	MaxOpenConns    string `env:"POSTGRES_MAX_OPEN_CONNS,default=5"`
	MaxConnLifetime string `env:"POSTGRES_MAX_CONN_LIFETIME,default=10m"`
	MaxIdleLifetime string `env:"POSTGRES_MAX_IDLE_LIFETIME,default=5m"`
}

// Redis holds configuration for the Redis.
type Redis struct {
	Address  string `env:"REDIS_ADDRESS"`
	Password string `env:"REDIS_PASSWORD"`
}

// SMTP holds configuration for smtp email.
type SMTP struct {
	Host string `env:"SMTP_HOST,required"`
	Port int    `env:"SMTP_PORT,default=587"`
	User string `env:"SMTP_USER,required"`
	Pass string `env:"SMTP_PASS,required"`
	From string `env:"SMTP_FROM,required"`
}

// JWTConfig holds configuration for jwt.
type JWTConfig struct {
	Public    string `env:"JWT_PUBLIC,required"`
	Private   string `env:"JWT_PRIVATE,required"`
	Issuer    string `env:"JWT_ISSUER,required"`
	IssuerCMS string `env:"JWT_ISSUER_CMS,required"`
}

// OneSignal holds configuration for the OneSignal.
type OneSignal struct {
	AppID  string `env:"ONESIGNAL_APP_ID"`
	AppKey string `env:"ONESIGNAL_APP_KEY"`
}

type MailGun struct {
	From   string `env:"MAILGUN_FROM"`
	Domain string `env:"MAILGUN_DOMAIN"`
	APIKey string `env:"MAILGUN_API_KEY"`
}

// Jaeger holds configuration for the Jaeger.
type Jaeger struct {
	Address string `env:"JAEGER_ADDRESS"`
	Port    string `env:"JAEGER_PORT"`
}

func LoadConfig(env string) (*Config, error) {
	// just skip loading env files if it is not exists, env files only used in local dev
	_ = godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
