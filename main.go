package main

import (
    "github.com/leeprince/gormstruct/internal/cmd"
    "github.com/leeprince/gormstruct/internal/config"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/10/30 下午3:09
 * @Desc:
 */

func main() {
    config.InitConfig()
    
    cmd.NewRootCmd().Execute()
}