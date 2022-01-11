package config

import (
    "github.com/leeprince/gormstruct/internal/constants"
    "github.com/leeprince/gormstruct/internal/logger"
    "github.com/spf13/cobra"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/4 上午12:33
 * @Desc:
 */

type flags struct {
    database    string
    table       string
    packageName string
    structName  string
}

var flagInfo flags

func InitFlags(c *cobra.Command) {
    var err error
    flagInfo.database, err = c.PersistentFlags().GetString(constants.FlagsOfDataBase)
    logger.PanicIfError(err)
    
    flagInfo.table, err = c.PersistentFlags().GetString(constants.FlagsOfTable)
    logger.PanicIfError(err)
    if flagInfo.table == "" {
        logger.Panic("必须指定表名 -t={表名} 或者 --tabel={表名}")
    }
    
    flagInfo.packageName, err = c.PersistentFlags().GetString(constants.FlagsOfPackageName)
    logger.PanicIfError(err)
    
    flagInfo.structName, err = c.PersistentFlags().GetString(constants.FlagsOfStructName)
    logger.PanicIfError(err)
    
    logger.Infof("InitFlags...flagInfo:%+v", flagInfo)
}

func GetFlagsDatabase() string {
    return flagInfo.database
}
