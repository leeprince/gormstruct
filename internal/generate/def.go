package generate

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/14 下午3:00
 * @Desc:
 */
// 生成打印接口
type IGenerate interface {
    Generate() string
}

// 原始打印的文件内容
type PrintAtom struct {
	lines []string
}