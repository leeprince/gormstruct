package generate

import (
    "fmt"
    "github.com/spf13/cast"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/9 下午11:44
 * @Desc:
 */
// 添加打印的一行信息
func (p *PrintAtom) Add(strs ...interface{}) {
    var tmp string
    for _, i2 := range strs {
        tmp += fmt.Sprintf("%s \t", cast.ToString(i2))
    }
    p.lines = append(p.lines, tmp)
}

// 获取打印的多行信息
func (p *PrintAtom) Generates() []string {
	return p.lines
}


