package out

import (
    "fmt"
    model_bak "github.com/leeprince/gormstruct/out/model.bak"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "testing"
    "time"
)

var mysqlDns = "root:leeprince@tcp(127.0.0.1:3306)/tmp?charset=utf8&parseTime=True&loc=Local&interpolateParams=True"

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
    
    var user model_bak.Users
    var err error
    
    user, err = model_bak.UsersMgr(db).Get()
    fmt.Println("user, err:", user, err)
    
    user, err = model_bak.UsersMgr(db.Where("id = ?", 1)).Get()
    fmt.Println("user, err:", user, err)
}

func TestModelGets(t *testing.T) {
    db := GetGorm()
    
    var users []*model_bak.Users
    var err error
    
    users, err = model_bak.UsersMgr(db.Where("id = ?", 1)).Gets()
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, *i2)
    }
    
    users, err = model_bak.UsersMgr(db.Where("name = ?", "name01")).Gets()
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, *i2)
    }
}

func TestModelCout(t *testing.T) {
    db := GetGorm()
    
    var count int64
    
    db1 := model_bak.UsersMgr(db.Where("id = ?", 11)).Count(&count)
    fmt.Printf("count:%+v, db.err:%v \n", count, db1.Error)
    
    db2 := model_bak.UsersMgr(db.Where("id in (?)", []int64{1, 2})).Count(&count)
    fmt.Printf("count:%+v, db.err:%v \n", count, db2.Error)
}

func TestModelWith(t *testing.T) {
    db := GetGorm()
    
    var user model_bak.Users
    var users []*model_bak.Users
    var err error
    
    usersMgr := model_bak.UsersMgr(db)
    user, err = usersMgr.GetByOption(usersMgr.WithID(1))
    fmt.Println("user, err:", user, err)
    
    users, err = usersMgr.GetByOptions(usersMgr.WithID(1))
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
    
}

func TestModelFrom(t *testing.T) {
    db := GetGorm()
    
    var user model_bak.Users
    var users []*model_bak.Users
    var err error
    
    user, err = model_bak.UsersMgr(db).GetFromID(1)
    fmt.Println("GetFromID..user, err:", user, err)
    
    users, err = model_bak.UsersMgr(db).GetBatchFromID([]int{1, 2})
    for _, i2 := range users {
        fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, *i2)
    }
    
    users, err = model_bak.UsersMgr(db).GetFromName("name01")
    fmt.Println("GetFromName..user, err:", user, err)
    
    users, err = model_bak.UsersMgr(db).GetBatchFromID([]int{1, 2})
    for _, i2 := range users {
        fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, *i2)
    }
    
    user, err = model_bak.UsersMgr(db).GetFromCardNo("1")
    fmt.Println("GetFromCardNo..user, err:", user, err)
    
    users, err = model_bak.UsersMgr(db).GetBatchFromCardNo([]string{"1", "2"})
    for _, i2 := range users {
        fmt.Printf("GetBatchFromCardNo..err:%v, users:%+v \n", err, *i2)
    }
    
}

func TestModelFetch(t *testing.T) {
    db := GetGorm()
    
    var user model_bak.Users
    var users []*model_bak.Users
    var err error
    
    user, err = model_bak.UsersMgr(db).FetchByPrimaryKey(1)
    fmt.Println("FetchByPrimaryKey..user, err:", user, err)
    
    user, err = model_bak.UsersMgr(db).FetchUniqueByCardNo("1")
    fmt.Println("FetchUniqueByCardNo..user, err:", user, err)
    
    user, err = model_bak.UsersMgr(db).FetchUniqueIndexByUnqNameCard("name01", "1")
    fmt.Println("FetchUniqueIndexByUnqNameCard..user, err:", user, err)
    
    users, err = model_bak.UsersMgr(db).FetchIndexByAge(12)
    for _, i2 := range users {
        fmt.Printf("FetchIndexByAge..err:%v, users:%+v \n", err, *i2)
    }
}
