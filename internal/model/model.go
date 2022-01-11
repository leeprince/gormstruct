package model

import (
    "bytes"
    "fmt"
    "github.com/leeprince/gormstruct/internal/config"
    "github.com/leeprince/gormstruct/internal/constants"
    "github.com/leeprince/gormstruct/internal/genfunc"
    "github.com/leeprince/gormstruct/internal/genstruct"
    "github.com/leeprince/gormstruct/internal/utils"
    "strings"
    "text/template"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/11/9 下午11:43
 * @Desc:
 */

type GenDBInfo struct {
    info DBInfo
    pkg  *genstruct.GenPackage
}

func Generate(info DBInfo) (out []GenOutInfo, g GenDBInfo) {
    g = GenDBInfo{
        info: info,
    }
    
    // 设置生成表结构体的数据
    outByTable := g.GenerateByTableName()
    out = append(out, outByTable...)
    
    // 添加生成获取表数据的基本方法的文件及内容
    out = append(out, g.generateFunc()...)
    
    return
}

func (g *GenDBInfo) GenerateByTableName() (out []GenOutInfo) {
    if g.pkg == nil {
        tab := g.info.Table
        
        var pkg genstruct.GenPackage
        pkg.SetPackage(g.info.PackageName)
        
        var sct genstruct.GenStruct
        sct.SetStructName(utils.GetCamelName(tab.Name))
        sct.SetStructTableName(tab.Name)
        sct.SetStructComment(tab.Comment)
        sct.SetStructElments(g.genTableElement(tab.ColumnsElement)...)
        
        pkg.SetStructs(sct)
        
        var outInfo GenOutInfo
        outInfo.FileName = fmt.Sprintf("%s.go", tab.Name)
        outInfo.FileCtx = pkg.GenFileCtx()
        out = append(out, outInfo)
    }
    return
}

func (g *GenDBInfo) genTableElement(els []ColumnsElementInfo) (genEls []genstruct.GenElement) {
    for _, el := range els {
        var genEl genstruct.GenElement
        genEl.Name = utils.GetCamelName(el.Name)
        genEl.ColumnName = el.Name
        genEl.Comment = el.Comment
        // 字段类型
        if strings.EqualFold(el.Type, constants.GormModelWord) { // gorm model
            genEl.Type = el.Type
        } else {
            genEl.Type = getTypeName(el.Type, el.IsNull)
        }
        
        isPrimary := false
        for _, i2 := range el.Keys {
            if i2.Key == ColumnsKeyPrimary {
                isPrimary = true
                break
            }
        }
        genEl.TagString = g.getStructTag(isPrimary, genEl.ColumnName)
        
        genEls = append(genEls, genEl)
    }
    return
}

func (g *GenDBInfo) getStructTag(isPrimary bool, clolumnName string) string {
    tagGorm := constants.TagDb
    tagJson := constants.TagJson
    
    var gormTag string
    if isPrimary {
        gormTag = fmt.Sprintf(constants.GenGormtagPrimary, tagGorm, clolumnName)
    } else {
        gormTag = fmt.Sprintf(constants.GenGormtagColumun, tagGorm, clolumnName)
    }
    jsonTag := fmt.Sprintf(constants.GenJson, tagJson, clolumnName)
    return fmt.Sprintf("`%s %s`", gormTag, jsonTag)
}

func (g *GenDBInfo) generateFunc() (genOut []GenOutInfo) {
    // --- 生成操作数据库的基本方法
    tmpl, err := template.New("GetGenBaseTemp").
        Funcs(template.FuncMap{"GetVV": func() string { return "`%v`" }}).
        Parse(genfunc.GetGenBaseTemp())
    if err != nil {
        panic(err)
    }
    var baseBuff bytes.Buffer
    err = tmpl.Execute(&baseBuff, g.info)
    if err != nil {
        panic(err)
    }
    genOut = append(genOut, GenOutInfo{
        FileName: constants.FILE_BASE_NAME,
        FileCtx:  baseBuff.String(),
    })
    // --- 生成操作数据库的基本方法-end
    
    // --- 生成操作数据表的方法
    var pkg genstruct.GenPackage
    pkg.SetPackage(g.info.PackageName)
    // 添加 import 信息
    pkg.AddImport(`"fmt"`)
    pkg.AddImport(`"context"`)
    pkg.AddImport(constants.ImportFile["gorm.Model"])
    
    var data funDef
    data.TableName = g.info.Table.Name
    if config.GetStructName() != "" {
        data.StructName = config.GetStructName()
    } else {
        data.StructName = utils.GetCamelName(g.info.Table.Name)
    }
    
    var primary, unique, uniqueIndex, index []FList
    // 根据表的属性信息构建模版所需字段
    for _, el := range g.info.Table.ColumnsElement {
        // 构建索引（包含主键）的方法需要的信息
        if strings.EqualFold(el.Type, constants.GormModelWord) {
            data.Em = append(data.Em, getGormModelElement()...)
            pkg.AddImport("\"time\"")
            buildFList(&primary, ColumnsKeyPrimary, "", config.GetPrimaryIdType(), "id")
        } else {
            typeName := getTypeName(el.Type, el.IsNull)
            // 该字段值在表中可重复
            isMulti := true
            for _, key := range el.Keys {
                
                switch key.Key {
                case ColumnsKeyPrimary:
                    isMulti = false
                    buildFList(&primary, ColumnsKeyPrimary, key.KeyName, typeName, el.Name)
                case ColumnsKeyUnique:
                    isMulti = false
                    buildFList(&unique, ColumnsKeyUnique, key.KeyName, typeName, el.Name)
                case ColumnsKeyUniqueIndex:
                    buildFList(&uniqueIndex, ColumnsKeyUniqueIndex, key.KeyName, typeName, el.Name)
                case ColumnsKeyIndex:
                    buildFList(&index, ColumnsKeyIndex, key.KeyName, typeName, el.Name)
                }
            }
            
            data.Em = append(data.Em, EmInfo{
                IsMulti:       isMulti,
                Notes:         fixNotes(el.Comment),
                Type:          typeName,
                ColName:       el.Name,
                ColNameEx:     fmt.Sprintf("`%v`", el.Name),
                ColStructName: utils.GetCamelName(el.Name),
            })
            if v2, ok := constants.ImportFile[typeName]; ok {
                if len(v2) > 0 {
                    pkg.AddImport(v2)
                }
            }
            
        }
    }
    data.Primary = append(data.Primary, primary...)
    data.Primary = append(data.Primary, unique...)
    data.Primary = append(data.Primary, uniqueIndex...)
    data.Index = append(data.Index, index...)
    
    logicTmpl, err := template.New("GetGenLogic").
        Funcs(GenLogicTemplateFuncs()).
        Parse(genfunc.GetGenLogic())
    if err != nil {
        panic(fmt.Sprintf("GetGenLogic err:%s", err.Error()))
    }
    var funcBuf bytes.Buffer
    logicErr := logicTmpl.Execute(&funcBuf, data)
    if logicErr != nil {
        panic(fmt.Sprintf("GetGenLogic Execute err:%s", logicErr.Error()))
    }
    pkg.AddFuncStr(funcBuf.String())
    genOut = append(genOut, GenOutInfo{
        FileName: fmt.Sprintf(g.info.DbName+".gen.%v.go", g.info.Table.Name),
        FileCtx:  pkg.GenFileCtx(),
    })
    // --- 生成操作数据表的方法-end
    
    return
}
