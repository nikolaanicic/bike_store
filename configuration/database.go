package configuration

type Database struct {
	User           string `yaml:"user"`
	Password       string `yaml:"password"` // password is a file field and contents should be read and stored in the field
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	DbName         string `yaml:"db_name"`
	MigrationsPath string `yaml:"migrations_path"`
}
