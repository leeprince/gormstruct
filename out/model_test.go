package out

import (
    "fmt"
    model "github.com/leeprince/gormstruct/out/model"
    // model "github.com/leeprince/gormstruct/out/model_bv2_lv2"
    // model "github.com/leeprince/gormstruct/out/model_bv1_lv1"
    // model "github.com/leeprince/gormstruct/out/model_bv2_lv1"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "testing"
    "time"
)

var mysqlDns = "root:leeprince@tcp(127.0.0.1:3306)/tmp?charset=utf8&parseTime=true&loc=Local&interpolateParams=True"

func GetGorm() *gorm.DB {
    mysqlConnDns := mysqlDns
    
    db, err := gorm.Open(mysql.Open(mysqlConnDns), &gorm.Config{PrepareStmt: false})
    if err != nil {
        panic("gorm open err:" + err.Error())
    }
    sqlDB, err := db.DB()
    if err != nil {
        panic("db.DB() err:" + err.Error())
    }
    // SetMaxIdleConns 设置空闲连接池中连接的最大数量
    sqlDB.SetMaxIdleConns(10)
    // SetMaxOpenConns 设置打开数据库连接的最大数量。
    sqlDB.SetMaxOpenConns(100)
    // SetConnMaxLifetime 设置了连接可复用的最大时间。
    sqlDB.SetConnMaxLifetime(time.Hour)
    return db.Debug()
}

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/5 下午8:40
 * @Desc:
 */

func TestModelGet(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var err error
    
    user, err = model.NewUsersModel(db).Get()
    fmt.Println("user, err:", user, err)
    
    user, err = model.NewUsersModel(db.Where("id = ?", 1)).Get()
    fmt.Println("user, err:", user, err)
    
    user, err = model.NewUsersModel(db.Where("id = ?", 2)).Get()
    fmt.Println("user, err:", user, err)
}

func TestModelGets(t *testing.T) {
    db := GetGorm()
    
    var users []*model.Users
    var err error
    
    users, err = model.NewUsersModel(db.Where("id = ?", 1)).Gets()
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, *i2)
    }
    
    users, err = model.NewUsersModel(db.Where("name = ?", "name01")).Gets()
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, *i2)
    }
}

func TestModelCount(t *testing.T) {
    db := GetGorm()
    
    var count int64
    
    db1 := model.NewUsersModel(db.Where("id = ?", 11)).Count(&count)
    fmt.Printf("count:%+v, db.err:%v \n", count, db1.Error)
    
    db2 := model.NewUsersModel(db.Where("id in (?)", []int64{1, 2})).Count(&count)
    fmt.Printf("count:%+v, db.err:%v \n", count, db2.Error)
}

// 依赖 gen_base_func_version == V1
// 暂注释
// func TestModelCondition(t *testing.T) {
//     db := GetGorm()
//
//     var users []*model.Users
//     var err error
//     var condition model.Condition
//     var where string
//     var args []interface{}
//
//     // condition.And("id", "in", []int{1, 2})
//     condition.And("id", "in", []int{1})
//     // condition.And("name", "in", []string{"name01", "name02"})
//     condition.Or("name", "in", []string{"name01", "name02"})
//     condition.And("age", "=", 13)
//     where, args = condition.Get()
//     fmt.Println(where, args)
//     users, err = model.NewUsersModel(db.Where(where, args...)).Gets()
//     for _, i2 := range users {
//         fmt.Printf("err:%v, users:%+v \n", err, *i2)
//     }
//     fmt.Println("---1")
//     condition = model.Condition{}
//     // condition.And("id", "in", []int{1, 2})
//     condition.And("id", "in", []int{1})
//     // condition.And("name", "in", []string{"name01", "name02"})
//     condition.Or("name", "in", []string{"name01", "name02"})
//     whereTmp, argsTmp := condition.Get()
//     fmt.Println(whereTmp, argsTmp)
//     fmt.Println("---2")
//     condition = model.Condition{}
//     condition.And("age", "=", 13)
//     where, args = condition.Get()
//     fmt.Println(where, args)
//     fmt.Println("---3")
//     where = fmt.Sprintf("(%s) AND %s", whereTmp, where)
//     argsTmp = append(argsTmp, args...)
//     fmt.Println(where, argsTmp)
//     users, err = model.NewUsersModel(db.Where(where, argsTmp...)).Gets()
//     for _, i2 := range users {
//         fmt.Printf("err:%v, users:%+v \n", err, *i2)
//     }
// }

func TestModelWith(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var users []*model.Users
    var err error
    
    usersMgr := model.NewUsersModel(db)
    
    user, err = usersMgr.GetByOption(usersMgr.WithID(1))
    fmt.Println("user, err:", user, err)

    usersMgr = model.NewUsersModel(db)
    users, err = usersMgr.GetByOptions(usersMgr.WithBatchID([]int{1, 2}))
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, *i2)
    }
    
    users, err = usersMgr.GetByOptions(usersMgr.WithID(1))
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, *i2)
    }
    
    users, err = usersMgr.GetByOptions(
        usersMgr.WithName("name01"),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, *i2)
    }
    
    users, err = usersMgr.GetByOptions(
        usersMgr.WithName("name01"),
        usersMgr.WithAge(12),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, *i2)
    }
    
    // 分页
    users, err = usersMgr.GetByOptions(
        usersMgr.WithName("name01"),
        usersMgr.WithPage(1, 2),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, *i2)
    }
}

func TestModelFrom(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var users []*model.Users
    var err error
    
    user, err = model.NewUsersModel(db).GetFromID(1)
    fmt.Println("GetFromID..user, err:", user, err)
    
    user, err = model.NewUsersModel(db).GetFromID(10000)
    fmt.Println("GetFromID..user, err:", user, err)
    
    users, err = model.NewUsersModel(db).GetBatchFromID([]int{1, 2})
    for _, i2 := range users {
        fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, *i2)
    }
    
    users, err = model.NewUsersModel(db).GetFromName("name01")
    fmt.Println("GetFromName..user, err:", user, err)
    
    users, err = model.NewUsersModel(db).GetBatchFromID([]int{1, 2})
    for _, i2 := range users {
        fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, *i2)
    }
    
    user, err = model.NewUsersModel(db).GetFromCardNo("1")
    fmt.Println("GetFromCardNo..user, err:", user, err)
    
    users, err = model.NewUsersModel(db).GetBatchFromCardNo([]string{"1", "2"})
    for _, i2 := range users {
        fmt.Printf("GetBatchFromCardNo..err:%v, users:%+v \n", err, *i2)
    }
    
}

func TestModelFetch(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var users []*model.Users
    var err error
    
    user, err = model.NewUsersModel(db).FetchByPrimaryKey(1)
    fmt.Println("FetchByPrimaryKey..user, err:", user, err)
    
    user, err = model.NewUsersModel(db).FetchUniqueByCardNo("1")
    fmt.Println("FetchUniqueByCardNo..user, err:", user, err)
    
    user, err = model.NewUsersModel(db).FetchUniqueIndexByUnqNameCard("name01", "1")
    fmt.Println("FetchUniqueIndexByUnqNameCard..user, err:", user, err)
    
    users, err = model.NewUsersModel(db).FetchIndexByAge(12)
    for _, i2 := range users {
        fmt.Printf("FetchIndexByAge..err:%v, users:%+v \n", err, *i2)
    }
}

func TestModelReset(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var err error
    
    userModel := model.NewUsersModel(db.Where("id = ?", 1))
    fmt.Println("user, err:", user, err)
    
    user, err = userModel.Get()
    fmt.Println("user, err:", user, err)
    user, err = userModel.Reset().Get()
    fmt.Println("user, err:", user, err)
    user, err = userModel.Reset().Get()
    fmt.Println("user, err:", user, err)
    user, err = model.NewUsersModel(userModel.Reset().DB).Get()
    fmt.Println("user, err:", user, err)
    
    fmt.Println("---------------")
    userModel = model.NewUsersModel(db.Where("id = ?", 2))
    fmt.Println("user, err:", user, err)
    
    user, err = userModel.Get()
    fmt.Println("user, err:", user, err)
    user, err = userModel.Reset().Get()
    fmt.Println("user, err:", user, err)
    user, err = userModel.Get()
    fmt.Println("user, err:", user, err)
    user, err = userModel.Reset().Get()
    fmt.Println("user, err:", user, *user.CreatedAt, err)
}
