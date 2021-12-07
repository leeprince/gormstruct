package utils

import "github.com/leeprince/gopublic/mybigcamel"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/13 下午12:31
 * @Desc:
 */

// 大驼峰或者首字母大写
func GetCamelName(name string) string {
	return mybigcamel.Marshal(name)
}