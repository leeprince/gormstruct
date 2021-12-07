package cmd

import (
    "fmt"
    "github.com/leeprince/gopublic/tools"
    "github.com/leeprince/gormstruct/internal/app"
    _ "github.com/leeprince/gormstruct/internal/config"
    "github.com/leeprince/gormstruct/internal/constants"
    "github.com/leeprince/gormstruct/internal/logger"
    "github.com/leeprince/gormstruct/internal/model"
    "github.com/leeprince/gormstruct/internal/model/genmysql"
    "github.com/spf13/cobra"
    "os"
    "os/exec"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/10/30 下午3:09
 * @Desc:
 */

type RootCmd struct {
    cmd *cobra.Command
}

var rootCmd = RootCmd{
    cmd: &cobra.Command{
        Use:   "gen",
        Short: "gorm mysql reflect tools",
        Long:  `base on gorm tools for mysql database to golang struct`,
        Run: func(cmd *cobra.Command, args []string) {
            // Start doing things.开始做事情
        },
        RunE: generateGoTag,
    },
}

func NewRootCmd() *RootCmd {
    return &rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (r RootCmd) Execute() {
    if err := r.cmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.cmd.PersistentFlags().StringP("table", "t", "", "Table Name")
    // rootCmd.cmd.MarkFlagRequired("table")
    
    rootCmd.cmd.PersistentFlags().StringP("packageName", "p", "", "Package Name")
    // rootCmd.cmd.MarkFlagRequired("packageName")
    
    rootCmd.cmd.PersistentFlags().StringP("structName", "s", "", "Struct Name")
    // rootCmd.cmd.MarkFlagRequired("structName")
}

func generateGoTag(c *cobra.Command, _ []string) error {
    table, _ := c.PersistentFlags().GetString("table")
    packageName, _ := c.PersistentFlags().GetString("packageName")
    structName, _ := c.PersistentFlags().GetString("structName")
    app.InitFlags(table, packageName, structName)
    logger.Infof("generateGoTag...table=%s;packageName=%s;structName=%s;", table, packageName, structName)
    
    var modeldb model.IModel
    switch app.GetConfigDBType() {
    case constants.DB_TYPE_MYSQL:
        modeldb = genmysql.NewMySQLModel()
    default:
        modeldb = genmysql.NewMySQLModel()
    }
    
    // 生成表相关的所有属性
    dbInfo := modeldb.GenModel()
    
    // 根据表相关的所有属性，生成用于输出到文件的格式
    genOutInfo, _ := model.Generate(dbInfo)
    fmt.Printf("所有输出的信息。generateGoTag.genOutInfo:\n%+v\n", genOutInfo)
    
    for _, i2 := range genOutInfo {
        path := app.GetOutputFileOfPath(i2.FileName)
        tools.WriteFile(path, []string{i2.FileCtx}, true)
        
        goImportByte, _ := exec.Command("goimports", "-l", "-w", path).Output()
        logger.Info("goimports:", string(goImportByte))
        
        // 格式化go文件的格式行并保存到原文件中
        gofmtByte, err := exec.Command("gofmt", "-l", "-w", path).Output()
        if err != nil {
            logger.Errorf("gofmt path:%s, err: %s", path, err.Error())
        } else {
            logger.Info("gofmt:", string(gofmtByte))
        }
    }
    
    return nil
}
