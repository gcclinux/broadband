package handler

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func GetConfig() Config {
	pwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := "/conf/conf.json"

	file, err := os.Open(pwd + filePath)
	if err != nil {
		fmt.Println("file error:", err)
		os.Exit(3)
	}

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error in decoding conf.json:", err)
	}
	return configuration
}

// Config File Structure
type Config struct {
	DB_HOST  []string `json:"db_host"`
	DB_PORT  []string `json:"db_port"`
	DB_USER  []string `json:"db_user"`
	DB_PASS  []string `json:"db_pass"`
	DB_NAME  []string `json:"db_name"`
	TB_NAME  []string `json:"tb_name"`
	Verbose  []bool   `json:"verbose"`
	SaveToDB []bool   `json:"savetodb"`
	TB_SIZE  []string `json:"TB_SIZE"`
}

type IP struct {
	Query string
}
