package config

type Config struct {
	App struct {
		Name string
		Env  string
	}
	Server struct {
		Host string
		Port string
	}
	MySQL struct {
		Host        string
		Port        string
		Database    string
		UserName    string
		Password    string
	}
	Redis struct {
		Host     string
		Port     string
		Database string
		Auth     string
	}
	GitHub struct {
		ClientId       string
		ClientKey      string
		ClientRedirect string
	}
	QiNiu struct {
		AccessKey string
		SecretKey string
		Bucket    string
		Domain    string
	}
}
