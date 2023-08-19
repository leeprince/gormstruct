package utils

import (
    "fmt"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/16 上午10:28
 * @Desc:   唯一ID
 */

func UniqID() string {
	now := time.Now()
	// %016x: 用0填充最小宽度16，十六进制格式
	return fmt.Sprintf("%016x", now.UnixNano())
}