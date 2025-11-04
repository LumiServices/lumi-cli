package config

type Config struct {
	Profiles map[string]*Profile `toml:"profiles"`
}

type Profile struct {
	Endpoint  string `toml:"endpoint"`
	Output    string `toml:"output"` // xml | json
	VerifySSL bool   `toml:"verify_ssl"`
}

type Credentials struct {
	AccessKey string `toml:"access_key"`
	SecretKey string `toml:"secret_key"`
}
