package config

import (
	"fmt"
	"io/ioutil"
	"path"
	"sap_cert_mgt/infra/utils"
	"sync"

	"gopkg.in/yaml.v2"
)

type ServiceConf struct {
	Lang        string `yaml:"lang"`
	Host        string `yaml:"host"`
	PublicPort  string `yaml:"public_port"`
	PrivatePort string `yaml:"private_port"`
}

type MysqlConf struct {
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	DbHost          string `yaml:"db_host"`
	DbPort          string `yaml:"db_port"`
	DbName          string `yaml:"db_name"`
	Charset         string `yaml:"charset"`
	Timeout         string `yaml:"timeout"`
	TimeoutRead     string `yaml:"timeout_read"`
	TimeoutWrite    string `yaml:"timeout_write"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type LogConf struct {
	LogDir  string `yaml:"log_dir"`
	LogName string `yaml:"log_name"`
}

type Config struct {
	Service ServiceConf
	Mysql   MysqlConf
	Log     LogConf
}

var (
	configOnce sync.Once
	config     *Config
)

// NewConfig .
func NewConfig() *Config {
	configOnce.Do(func() {
		projectPath := path.Dir(utils.GetProjectAbPathByCaller())
		configFilePath := path.Join(projectPath, "server", "config", "config.yaml")
		//configFilePath := "/sysvol/conf/sap-cert-mgt.yaml"
		file, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			panic(fmt.Sprintf("load %v failed: %v", configFilePath, err))
		}

		err = yaml.Unmarshal(file, &config)
		if err != nil {
			panic(fmt.Sprintf("unmarshal yaml file failed: %v", err))
		}
		//fmt.Printf("%v", config)
	})

	return config
}
