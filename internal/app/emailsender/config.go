package emailsender

type Config struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	TLSPort  string `json:"tls_port"`
}

func NewConfig() *Config {
	return &Config{}
}
