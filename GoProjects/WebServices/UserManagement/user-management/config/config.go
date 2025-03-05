package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port         string        `mapstructure:"PORT"`
		ReadTimeout  time.Duration `mapstructure:"READ_TIMEOUT"`
		WriteTimeout time.Duration `mapstructure:"WRITE_TIMEOUT"`
	} `mapstructure:"SERVER"`

	Database struct {
		Host     string `mapstructure:"HOST"`
		Port     string `mapstructure:"PORT"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		Name     string `mapstructure:"NAME"`
	} `mapstructure:"DATABASE"`

	Security struct {
		HTTPSEnabled      bool   `mapstructure:"HTTPS_ENABLED"`
		JWTSecretKey      string `mapstructure:"JWT_SECRET_KEY"`
		PasswordHashCost  int    `mapstructure:"PASSWORD_HASH_COST"`
		MFAEnabled        bool   `mapstructure:"MFA_ENABLED"`
		AuditLogEnabled   bool   `mapstructure:"AUDIT_LOG_ENABLED"`
		PasswordMinLength int    `mapstructure:"PASSWORD_MIN_LENGTH"`
	} `mapstructure:"SECURITY"`

	Roles struct {
		Admin   string `mapstructure:"ADMIN"`
		User    string `mapstructure:"USER"`
		Guest   string `mapstructure:"GUEST"`
		Default string `mapstructure:"DEFAULT"`
	} `mapstructure:"ROLES"`

	Reporting struct {
		ReportInterval time.Duration `mapstructure:"REPORT_INTERVAL"`
		ReportPath     string        `mapstructure:"REPORT_PATH"`
	} `mapstructure:"REPORTING"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// init 函数在 Go 语言中是一个特殊的函数，它会在包被导入时自动执行，并且在程序运行期间只会执行一次。
func init() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Println("Config loaded successfully:", config)
}
