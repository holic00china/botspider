package structure

type Config struct {
	UserAgent []string `yaml:"User-agent"`
}
type Env struct {
	Host    string `yaml:"URL"`
	Timeout int    `yaml:"TIMEOUT"`
}
