package master

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	ApiPort string `json:"apiPort"`
	ApiReadTiemout int `json:"apiReadTimeout"`
	ApiWriteTimeout int `json:"apiWriteTimeout"`
	EtcdEndpoints []string `json:"etcdEndpoints"`
	EtcdDialTimeout  int `json:"etcdDialTimeout"`
}

var(
	//单例
	G_config *Config
)
func InitConfig(fielname string)(err error){
	var(
		content []byte
		conf Config
	)
	//1.读取配置文件
	if content,err = ioutil.ReadFile(fielname); err != nil {
		return
	}

	//JOSN反序列化
	if err = json.Unmarshal(content,&conf); err != nil {
		return
	}
	G_config = &conf
	fmt.Println(conf)
	return
}