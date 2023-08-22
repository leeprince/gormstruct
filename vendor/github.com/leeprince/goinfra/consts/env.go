package consts

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/2 下午11:50
 * @Desc:
 */

const (
    ENVLocal   = "local"
    ENVDev     = "dev"
    ENVTest    = "test"
    ENVUat     = "uat"
    ENVSandbox = "sandbox"
    ENVProd    = "prod"
)

var AllENV = []string{
    ENVLocal,
    ENVDev,
    ENVTest,
    ENVUat,
    ENVSandbox,
    ENVProd,
}
