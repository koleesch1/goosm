// config.go
package config

type Config struct {
	AppName         string           `json:"app_name" toml:"appName" xml:"app_name" yaml:"appName"`
	Verbose         bool             `json:"verbose" toml:"verbose" xml:"verbose" yaml:"verbose"`
	Automigrate     bool             `json:"autmigrate" toml:"automigrate" xml:"autmigrate" yaml:"automigrate"`
	Database        string           `json:"database" toml:"database" xml:"database" yaml:"database"`
	DatabaseDialect string           `json:"databaseDialect" toml:"databaseDialect" xml:"databaseDialect" yaml:"databaseDialect"`
	Senseboxes      []SenseboxConfig `json:"senseboxes" toml:"senseboxes" xml:"senseboxes" yaml:"senseboxes"`
}

type SenseboxConfig struct {
	Name string `json:"name" toml:"name" xml:"name" yaml:"name"`
	Id   string `json:"id" toml:"id" xml:"id" yaml:"id"`
}

func defaultConfig() *Config {
	return &Config{
		AppName:         `Sensebox`,
		Verbose:         false,
		Automigrate:     true,
		Database:        `mysensebox.sqlite`,
		DatabaseDialect: `sqlite`,
		Senseboxes{Name:`Stefans Sensebox`,Id: `5a436432ec97390020d9c2ee`}
	}
}

