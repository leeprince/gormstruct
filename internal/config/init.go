package config

import (
    "github.com/leeprince/gopublic/tools"
    "github.com/leeprince/gormstruct/internal/logger"
    "gopkg.in/yaml.v3"
    "io/ioutil"
    "path"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/1 上午12:42
 * @Desc:
 */

// 引入文件初始化
func InitConfig() {
    once.Do(func() {
        configPath = path.Join(tools.GetModelPath(), "config/config.yaml")
        logger.Info("configPath:", configPath)
        if !tools.CheckFileIsExist(configPath) {
            logger.Panic("config.yaml not exit. using default config")
        }
        
        onInit()
    })
}

// 初始化
func onInit() {
    err := initConfigFile(configPath)
    if err != nil {
        // fmt.Println("Load config file error: ", err.Error())
        return
    }
}

// 初始化配置文件
func initConfigFile(filename string) error {
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        return err
    }
    
    if err := yaml.Unmarshal(bs, &config); err != nil {
        // fmt.Println("read config file error: ", err.Error())
        return err
    }
    
    logger.Info("Config:", config)
    return nil
}
