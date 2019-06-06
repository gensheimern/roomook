package config

import(
	"github.com/spf13/viper"
	"fmt"
)

type Config struct {
	Db string
	Db_name string
	Db_user string
	Db_password string
	Db_host string

	server_port int
}

func ReadConfig() Config{
	viper.SetConfigName("config") 
	viper.AddConfigPath("./config") 
	
	var cfg Config

	err := viper.ReadInConfig()
	if err != nil {

		fmt.Println("Config not found...")
		return cfg
	} else {

		cfg.Db = viper.GetString("database.Db")
		cfg.Db_name  = viper.GetString("database.Db_name")
		cfg.Db_user  = viper.GetString("database.Db_user")
		cfg.Db_password = viper.GetString("database.Db_password")
		cfg.Db_host = viper.GetString("database.Db_host")
		cfg.server_port = viper.GetInt("database.server_port")

		fmt.Println("======= Config found ======= ")
		fmt.Println("database: ", cfg.Db)
		fmt.Println("table: ", cfg.Db_name)
		fmt.Println("user: ", cfg.Db_user)
		fmt.Println("password: ", cfg.Db_password)
		fmt.Println("host: ", cfg.Db_host)
		fmt.Println("server_port: ", cfg.Db_host)
		return cfg
	}

}