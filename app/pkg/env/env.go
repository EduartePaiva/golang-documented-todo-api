package env

import (
	"log"
	"sync"

	"github.com/joeshaw/envdecode"
)

type envVariables struct {
	Database struct {
		DbHost      string `env:"DB_HOST"`
		DbUser      string `env:"DB_USER"`
		DbName      string `env:"DB_NAME"`
		DbPort      string `env:"DB_PORT"`
		DbPassword  string `env:"DB_PASSWORD"`
		DatabaseUrl string `env:"DATABASE_URL,required"`
	}
	OAuth2 struct {
		GitHub struct {
			ClientID     string `env:"GITHUB_CLIENT_ID,required"`
			ClientSecret string `env:"GITHUB_CLIENT_SECRET,required"`
			RedirectURI  string `env:"GITHUB_REDIRECT_URI,required"`
		}
		Google struct {
			ClientID     string `env:"GOOGLE_CLIENT_ID,required"`
			ClientSecret string `env:"GOOGLE_CLIENT_SECRET,required"`
			RedirectURI  string `env:"GOOGLE_REDIRECT_URI,required"`
		}
	}
	BasePath string `env:"BASE_PATH,default=."`
	GoEnv    string `env:"GO_ENV,default=production"`
}

var (
	instance *envVariables
	once     sync.Once
)

// loadAndValidateEnv loads environment variables and validates their presence.
func loadAndValidateEnv() *envVariables {
	env := envVariables{}
	err := envdecode.Decode(&env)
	if err != nil {
		log.Fatal("❌ ERROR DECODING ENVIRONMENT VARIABLES: ", err)
	}

	// return the env struct
	return &env
}

// Get provides a thread-safe singleton instance of envVariables.
func Get() *envVariables {
	once.Do(func() {
		instance = loadAndValidateEnv()
	})
	return instance
}
