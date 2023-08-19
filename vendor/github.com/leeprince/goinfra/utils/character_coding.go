package utils

import (
    "bytes"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io/ioutil"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/28 下午11:22
 * @Desc:   字符编码
 */

func GbkToUtf8(s []byte) ([]byte, error) {
     reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
     d, e := ioutil.ReadAll(reader)
     if e != nil {
         return nil, e
     }
     return d, nil
 }

 func Utf8ToGbk(s []byte) ([]byte, error) {
     reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
     d, e := ioutil.ReadAll(reader)
     if e != nil {
         return nil, e
     }
     return d, nil
 }