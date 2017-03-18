package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"gopkg.in/gorp.v1"
	// libreria de Oracle
	_ "gopkg.in/rana/ora.v3"
)

var DbmapOracle = InitDbOracle()

// InitDbOracle Funcion que carga la conecccion a al BD Global
func InitDbOracle() *gorp.DbMap {
	viper.SetConfigName("taller") // name of config file (without extension) envoca un archivo json y actualiza la informacion
	viper.AddConfigPath(".")         // path to look for the config file in

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config not found...taller.json")
		fmt.Println(err)
		return nil
	}
	var cadenaCon string
	if viper.GetString("Produccion") == "SI" {
		cadenaCon = viper.GetString("pro.user") + "/" + viper.GetString("pro.pass") + "@" + viper.GetString("pro.server") + ":" + viper.GetString("pro.port") + "/" + viper.GetString("pro.sid")
	} else {
		cadenaCon = viper.GetString("dev.user") + "/" + viper.GetString("dev.pass") + "@" + viper.GetString("dev.server") + ":" + viper.GetString("dev.port") + "/" + viper.GetString("dev.sid")
	}
	log.Println("Config found, name = ", cadenaCon)
	// fmt.Println("--------")
	// fmt.Println(cadenaCon)

	// db, err := sql.Open("ora", "system/admin@localhost:1521/xe")
	// db, err := sql.Open("ora", "pro/pro@192.168.36.6:1521/orcl")
	db, err := sql.Open("ora", cadenaCon)
	if err != nil {
		log.Println(err)
	}
	// defer db.Close()
	DbmapOracle := &gorp.DbMap{Db: db, Dialect: gorp.OracleDialect{}}
	return DbmapOracle
}
