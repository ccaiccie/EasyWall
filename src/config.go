package main

import (
	"os"

	"github.com/go-ini/ini"
)

// ConfigGetValue - Read Value from Configuration File
func ConfigGetValue(section string, key string) string {
	checkConfigExists()
	cfg, err := ini.Load("config/easywall.ini")
	if err != nil {
		Error(err)
		return ""
	}
	if cfg.Section(section).HasKey(key) {
		return cfg.Section(section).Key(key).String()
	}
	return ""
}

// ConfigSetValue - Place a string value into configuration file
func ConfigSetValue(section string, key string, value string) {
	checkConfigExists()
	cfg, err := ini.Load("config/easywall.ini")
	if err != nil {
		Error(err)
	}
	if cfg.Section(section).HasKey(key) {
		cfg.Section(section).Key(key).SetValue(value)
	} else {
		cfg.Section(section).NewKey(key, value)
	}
	err = cfg.SaveTo("config/easywall.ini")
	if err != nil {
		Error(err)
	}
}

func checkConfigExists() {
	if _, err := os.Stat("config"); os.IsNotExist(err) {
		err = os.Mkdir("config", os.ModePerm)
		if err != nil {
			Error(err)
		}
	}
	if _, err := os.Stat("config/easywall.ini"); os.IsNotExist(err) {
		createDefaultConfigFile()
	}
}

func createDefaultConfigFile() {
	cfg := ini.Empty()

	// Section log
	cfg.Section("log").NewKey("enable", "true")
	cfg.Section("log").NewKey("logfile", "easywall.log")
	cfg.Section("log").NewKey("logdir", "./log/")

	// Section Server
	cfg.Section("server").NewKey("ssl", "true")
	cfg.Section("server").NewKey("address", "0.0.0.0")
	cfg.Section("server").NewKey("port", "8080")
	cfg.Section("server").NewKey("certfile", "cert.pem")
	cfg.Section("server").NewKey("keyfile", "key.pem")

	err := cfg.SaveTo("config/easywall.ini")
	if err != nil {
		Error(err)
	}

	Log("Creating new configuration file easywall.ini in folder config using default values.")
	Log("Edit easywall.ini and restart EasyWall after changing the configuration")
}
