package app

import (
    "fmt"
    "github.com/leeprince/gormstruct/internal/constants"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/22 下午9:40
 * @Desc:
 */

func GetOutputDir() string {
    return constants.OuputDir
}

func GetOutputFileOfPath(fileName string) string {
    return fmt.Sprintf("%s%s", GetOutputDir(), fileName)
}

func GetSelfTypeDefine() map[string]string {
    return constants.SelfTypeDefine
}

func GetIsNullToPoint() bool {
    return constants.IsNullToPoint
}