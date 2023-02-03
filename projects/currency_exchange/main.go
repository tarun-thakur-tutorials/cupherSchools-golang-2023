package main

import (
	"example.com/database"
	"github.com/spf13/viper"
)

func main() {
	// for json
	/*
			viper.AddConfigPath("./configs")
			viper.SetConfigName("config")
			viper.SetConfigType("json")

			viper.ReadInConfig()

		fmt.Println(" the viper parameter ", viper.Get("dbname"))
	*/

	// for .env file
	/*
		viper.SetConfigFile(".env")

		viper.ReadInConfig()

		fmt.Println(" the viper parameter ", viper.Get("DBNAME"))
	*/

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.ReadInConfig()

	user := viper.Get("user")
	pass := viper.Get("pass")
	dbName := viper.Get("dbname")

	database.RegisterDatabase(user.(string), pass.(string), dbName.(string))
	// _, _ = utils.DownloadXML("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	database.LoadXMLInDatabase()
}
