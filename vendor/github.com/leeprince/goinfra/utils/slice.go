package utils

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 下午10:14
 * @Desc:
 */

func InString(s string, ss []string) bool {
    for _, i2 := range ss {
        if i2 == s {
            return true
        }
    }
    return false
}

func InInt(s int, ss []int) bool {
    for _, i2 := range ss {
        if i2 == s {
            return true
        }
    }
    return false
}

func InUint8(s uint8, ss []uint8) bool {
    for _, i2 := range ss {
        if i2 == s {
            return true
        }
    }
    return false
}

func InInt32(s int32, ss []int32) bool {
    for _, i2 := range ss {
        if i2 == s {
            return true
        }
    }
    return false
}

func InInt64(s int64, ss []int64) bool {
    for _, i2 := range ss {
        if i2 == s {
            return true
        }
    }
    return false
}
