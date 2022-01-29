package test

import (
    "fmt"
    "github.com/leeprince/gormstruct/out/model"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "testing"
    "time"
)

var mysqlDns = "root:leeprince@tcp(127.0.0.1:3306)/tmp?charset=utf8&parseTime=true&loc=Local&interpolateParams=True"

const IsDebug = true

var DBLogger = logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
    logger.Config{
        SlowThreshold:             time.Second, // 慢 SQL 阈值
        LogLevel:                  logger.Warn, // 日志级别
        IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
        Colorful:                  true,        // 彩色打印
    },
)

func GetGorm() *gorm.DB {
    mysqlConnDns := mysqlDns
    
    db, err := gorm.Open(mysql.Open(mysqlConnDns), &gorm.Config{
        PrepareStmt: false,
        Logger:      DBLogger,
    })
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
    
    if IsDebug {
        return db.Debug()
    }
    return db
}

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/5 下午8:40
 * @Desc:
 */
func TestModelGetTableName(t *testing.T) {
    db := GetGorm()
    
    userTableName := model.NewUsersDao(db).GetTableName()
    fmt.Println("userTableName:", userTableName)
}

func TestModelGet(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var err error
    
    user, err = model.NewUsersDao(db).Get()
    fmt.Println("user, err:", user, err)
    
    user, err = model.NewUsersDao(db.Where("id = ?", 1)).Get()
    fmt.Println("user, err:", user, err)
    
    user, err = model.NewUsersDao(db.Where("id = ?", 2)).Get()
    fmt.Println("user, err:", user, err)
    
    user, err = model.NewUsersDao(db.Where("id = ?", 20000)).Get()
    fmt.Println("user, err:", user, err)
}

func TestModelGets(t *testing.T) {
    db := GetGorm()
    
    var users []*model.Users
    var err error
    
    users, err = model.NewUsersDao(db.Where("id = ?", 1)).Gets()
    fmt.Printf("err:%v, users:%+v \n", err, users)
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    
    users, err = model.NewUsersDao(db.Where("id = ?", 1000)).Gets()
    fmt.Printf("err:%v, users:%+v \n", err, users)
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    
    users, err = model.NewUsersDao(db.Where("name = ?", "name01")).Gets()
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
}

func TestModelCount(t *testing.T) {
    db := GetGorm()
    
    var count int64
    
    db1 := model.NewUsersDao(db.Where("id = ?", 11)).Count(&count)
    fmt.Printf("count:%+v, db.err:%v \n", count, db1.Error)
    
    db2 := model.NewUsersDao(db.Where("id in (?)", []int64{1, 2})).Count(&count)
    fmt.Printf("count:%+v, db.err:%v \n", count, db2.Error)
}

func TestModelCreate(t *testing.T) {
    db := GetGorm()
    
    // CreatedAt/UpdatedAt：设置值时覆盖，为0时会自动生成
    deletedAt := 1
    users := model.Users{
        // ID:        0,
        Name: "insert-prince03",
        // Age:       18,
        Age:       0,
        CardNo:    "46100001",
        HeadImg:   "https://dd.xx",
        CreatedAt: 1643399938,
        UpdatedAt: 0,
        DeletedAt: int64(deletedAt),
    }
    rowsAffected, err := model.NewUsersDao(db).Create(&users)
    fmt.Printf("users:%+v rowsAffected:%d err:%v \n", users, rowsAffected, err)
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
//     users, err = model.NewUsersDao(db.Where(where, args...)).Gets()
//     for _, i2 := range users {
//         fmt.Printf("err:%v, users:%+v \n", err, i2)
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
//     users, err = model.NewUsersDao(db.Where(where, argsTmp...)).Gets()
//     for _, i2 := range users {
//         fmt.Printf("err:%v, users:%+v \n", err, i2)
//     }
// }

func TestModelWith(t *testing.T) {
    db := GetGorm()
    
    // var user model.Users
    // var users []*model.Users
    // var err error
    // var count int64
    
    usersMgr := model.NewUsersDao(db)
    /*
       // or 条件查询
       userCol := model.UsersColumns
       users, err = usersMgr.GetByOptions(
           usersMgr.WithSelect([]string{userCol.ID, userCol.Age}),
           usersMgr.WithID(1),
           usersMgr.WithOrOption(
               usersMgr.WithAge(18),
               usersMgr.WithName("name01"),
           ),
       )
       for _, i2 := range users {
           fmt.Printf("err:%v, users:%+v \n", err, i2)
       }
    */
    
    /*
       // 查询指定字段
       userCol := model.UsersColumns
       user, err = usersMgr.GetByOption(
           usersMgr.WithSelect([]string{userCol.ID, userCol.Age}),
           // usersMgr.WithSelect(fmt.Sprintf("%s, %s", userCol.ID, userCol.Age)),
           // usersMgr.WithSelect(fmt.Sprintf("%s, sum(%s) AS age", userCol.ID, userCol.Age)),
           // usersMgr.WithSelect(fmt.Sprintf("IFNULL(%s, %d) AS age", userCol.Age, 100)),
           // usersMgr.WithSelect(fmt.Sprintf("IFNULL(%s, ?) AS age", userCol.Age), 100),
           // usersMgr.WithSelect(fmt.Sprintf("%s, IFNULL(%s, ?) AS age", userCol.ID, userCol.Age), 100),
           // usersMgr.WithSelect(fmt.Sprintf("%s, IF(%s, %s, ?) AS age", userCol.ID, userCol.Age, userCol.Age), 100),
           usersMgr.WithID(1),
       )
       fmt.Println("user, err:", user, err)
       users, err = usersMgr.GetByOptions(
           // usersMgr.WithSelect(fmt.Sprintf("%s, %s", userCol.ID, userCol.Age)),
           // usersMgr.WithSelect([]string{userCol.ID, userCol.Age}),
           // usersMgr.WithSelect(fmt.Sprintf("%s, sum(%s) AS age", userCol.ID, userCol.Age)),
           // usersMgr.WithSelect(fmt.Sprintf("IFNULL(%s, %d) AS age", userCol.Age, 100)),
           // usersMgr.WithSelect(fmt.Sprintf("IFNULL(%s, ?) AS age", userCol.Age), 100),
           // usersMgr.WithSelect(fmt.Sprintf("%s, IFNULL(%s, ?) AS age", userCol.ID, userCol.Age), 100),
           usersMgr.WithSelect(fmt.Sprintf("%s, IF(%s, %s, ?) AS age", userCol.ID, userCol.Age, userCol.Age), 100),
           usersMgr.WithBatchID([]int{1, 2}),
       )
       for _, i2 := range users {
           fmt.Printf("err:%v, users:%+v \n", err, i2)
       }
    */
    
    /*
       // 更新方法
       users, err = usersMgr.GetByOptions(
           usersMgr.WithID(1),
       )
       for _, i2 := range users {
           fmt.Printf("err:%v, users:%+v \n", err, i2)
       }
       // dtime := 0
       dtime := 1642337297
       usesUpdate := model.Users{
           // ID: 100,
           Name:      "prince-update",
           Age:       0, // 需要把年龄更新为0，注意 当通过 struct 更新时，GORM 只会更新非零字段。 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
           HeadImg:   "",
           DeletedAt: &dtime,
       }
       count, err = usersMgr.UpdateByOption(
           usesUpdate,
           usersMgr.WithID(1),
       )
       userCol := model.UsersColumns
       count, err = usersMgr.UpdateByOption(
           usesUpdate,
           usersMgr.WithSelect([]string{userCol.Name, userCol.Age, userCol.Age, userCol.HeadImg, userCol.DeletedAt}), // 更新指定字段
           usersMgr.WithID(1),
       )
       fmt.Printf("UpdateByOption count:%v, err:%+v \n", count, err)
    */
    
    /*
    // 分组+筛选
    users, err = usersMgr.GetByOptions(
        usersMgr.WithHaveUpdateName("name01"),
        usersMgr.WithName("name01"),
        usersMgr.WithAge(12),
        usersMgr.WithOrOption(
            usersMgr.WithAge(18),
            usersMgr.WithName("name02"),
        ),
        usersMgr.WithOrderBy(fmt.Sprintf("%s desc", model.UsersColumns.Name)),
        usersMgr.WithGroupBy(model.UsersColumns.Age),
        usersMgr.WithHaving(fmt.Sprintf("%s > ?", model.UsersColumns.Age), 15),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    */
    
    /*
    // option 条件查询
    user, err := usersMgr.GetByOption(usersMgr.WithID(1))
    fmt.Println("user, err:", user, err)
    usersMgr = model.NewUsersDao(db)
    users, err := usersMgr.GetByOptions(usersMgr.WithBatchID([]int{1, 2}))
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    users, err = usersMgr.GetByOptions(usersMgr.WithID(1))
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    users, err = usersMgr.GetByOptions(
        usersMgr.WithName("name01"),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    users, err = usersMgr.GetByOptions(
        usersMgr.WithName("name01"),
        usersMgr.WithAge(12),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    */
    
    /*
    // 分页
    users, err := usersMgr.GetByOptions(
        usersMgr.WithName("name01"),
        // usersMgr.WithName("name01"),
        // usersMgr.WithPage(0, 2),
        usersMgr.WithPage(1, 2),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    */
    
    // 根据 option 条件统计数量
    var count int64
    count = usersMgr.GetCountByOptions(
        usersMgr.WithName("name01"),
    )
    fmt.Println("count", count)
    count = usersMgr.GetCountByOptions(
        usersMgr.WithName("name010000"),
    )
    fmt.Println("count", count)
}

// GetFromXxx 返回单条记录时，传入的参数为空值（0，""，nil）时会忽略为查询条件
func TestModelFrom(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var users []*model.Users
    var err error
    
    user, err = model.NewUsersDao(db).GetFromID(1)
    fmt.Println("GetFromID..user, err:", user, err)
    
    users, err = model.NewUsersDao(db).GetFromName("ddd")
    fmt.Println("GetFromName..users, err:", users, err)
    for _, i2 := range users {
        fmt.Printf("GetFromName..err:%v, users:%+v \n", err, i2)
    }
    
    var deletedAt int
    deletedAt = 0
    // deletedAt = 1
    users, err = model.NewUsersDao(db).GetFromDeletedAt(int64(deletedAt))
    for _, i2 := range users {
        fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, i2)
    }
    
    deletedAt1 := 1639411296
    deletedAt2 := 1639411297
    deletedAt3 := 0
    deletedAts := []int64{
        int64(deletedAt1),
        int64(deletedAt2),
        int64(deletedAt3),
    }
    users, err = model.NewUsersDao(db).GetBatchFromDeletedAt(deletedAts)
    fmt.Println("GetFromID..users, err:", users, err)
    
    // user, err = model.NewUsersDao(db).GetFromID(10000)
    // fmt.Println("GetFromID..user, err:", user, err)
    //
    // users, err = model.NewUsersDao(db).GetBatchFromID([]int{1, 2})
    // for _, i2 := range users {
    //     fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, i2)
    // }
    //
    // users, err = model.NewUsersDao(db).GetFromName("name01")
    // fmt.Println("GetFromName..user, err:", user, err)
    //
    // users, err = model.NewUsersDao(db).GetBatchFromID([]int{1, 2})
    // for _, i2 := range users {
    //     fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, i2)
    // }
    //
    // user, err = model.NewUsersDao(db).GetFromCardNo("1")
    // fmt.Println("GetFromCardNo..user, err:", user, err)
    //
    // users, err = model.NewUsersDao(db).GetBatchFromCardNo([]string{"1", "2"})
    // for _, i2 := range users {
    //     fmt.Printf("GetBatchFromCardNo..err:%v, users:%+v \n", err, i2)
    // }
    
}

func TestModelFetch(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var users []*model.Users
    var err error
    
    user, err = model.NewUsersDao(db).FetchByPrimaryKey(1)
    fmt.Println("FetchByPrimaryKey..user, err:", user, err)
    
    user, err = model.NewUsersDao(db).FetchUniqueByCardNo("1")
    fmt.Println("FetchUniqueByCardNo..user, err:", user, err)
    
    user, err = model.NewUsersDao(db).FetchUniqueIndexByUnqNameCard("name01", "1")
    fmt.Println("FetchUniqueIndexByUnqNameCard..user, err:", user, err)
    
    users, err = model.NewUsersDao(db).FetchIndexByAge(12)
    for _, i2 := range users {
        fmt.Printf("FetchIndexByAge..err:%v, users:%+v \n", err, i2)
    }
}

func TestModelReset(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var err error
    
    userModel := model.NewUsersDao(db.Where("id = ?", 2))
    user, err = userModel.Get()
    fmt.Println("user, err:", user, err)
    
    fmt.Println("---------------")
    userModel = model.NewUsersDao(db.Where("id = ?", 3))
    user, err = userModel.Get()
    fmt.Println("user006------, err:", user, err)
    user, err = model.NewUsersDao(userModel.NewDB()).Get()
    fmt.Println("user007------, err:", user, err)
    user, err = userModel.Get()
    fmt.Println("user008------, err:", user, err)
}
