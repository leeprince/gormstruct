package utils

import (
    "github.com/leeprince/goinfra/consts"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/25 下午9:21
 * @Desc:   环境检测
 */

func IsProd(env string) (is bool) {
    is = false
    if env == consts.ENVProd {
        is = true
    }
    return
}

func IsProdOrSandbox(env string) (is bool) {
    return env == consts.ENVProd || env == consts.ENVUat || env == consts.ENVSandbox
}

func IsLocal(env string) (is bool) {
    is = false
    if env == consts.ENVLocal {
        is = true
    }
    return
}