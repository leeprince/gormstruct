package test

import (
    "context"
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "github.com/leeprince/goinfra/utils"
    "github.com/leeprince/gormstruct/test/model"
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

func InitLooger() {
    // err := plog.SetOutputFile("./logs", "gorm.log", false)
    err := plog.SetOutputFile("./logs", "gorm.log", true)
    if err != nil {
        panic(fmt.Sprintf("plog.SetOutputFile err:%v", err))
    }
    // plog.SetReportCaller(true)
    
    plog.Debugln("0001", "0002")
    plog.WithFiledLogID(utils.UniqidID()).WithField("key01", "v001").Debugln("0001", "0002")
}

var logWriterStdout = log.New(os.Stdout, "\r\n", log.LstdFlags) // io writer（日志输出的目标，前缀和日志包含的内容——译者注）

var DBLogger = logger.New(
    // logWriterStdout, // 标准输出
    plog.GetLogger(), // 指定日志文件输出
    logger.Config{
        SlowThreshold:             time.Second, // 慢 SQL 阈值
        LogLevel:                  logger.Warn, // 日志级别
        IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
        Colorful:                  true,        // 彩色打印
    },
)

func GetGorm() *gorm.DB {
    InitLooger()
    
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
    
    userTableName := model.NewUsersDao(context.Background(), db).GetTableName()
    fmt.Println("userTableName:", userTableName)
}

func TestModelGet(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var err error
    
    user, err = model.NewUsersDao(context.Background(), db).Get()
    fmt.Println("user, err:", user, err)
    
    user, err = model.NewUsersDao(context.Background(), db.Where("id = ?", 1)).Get()
    fmt.Println("user, err:", user, err)
    
    user, err = model.NewUsersDao(context.Background(), db.Where("id = ?", 2)).Get()
    fmt.Println("user, err:", user, err)
    
    user, err = model.NewUsersDao(context.Background(), db.Where("id = ?", 20000)).Get()
    fmt.Println("user, err:", user, err)
}

func TestModelGets(t *testing.T) {
    db := GetGorm()
    
    var users []*model.Users
    var err error
    
    users, err = model.NewUsersDao(context.Background(), db.Where(model.Users{ID: 0})).Gets()
    fmt.Printf("err:%v, users:%+v \n", err, users)
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    
    users, err = model.NewUsersDao(context.Background(), db.Where("id = ?", 0)).Gets()
    fmt.Printf("err:%v, users:%+v \n", err, users)
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    
    users, err = model.NewUsersDao(context.Background(), db.Where("id = ?", 1)).Gets()
    fmt.Printf("err:%v, users:%+v \n", err, users)
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    
    users, err = model.NewUsersDao(context.Background(), db.Where("id = ?", 1000)).Gets()
    fmt.Printf("err:%v, users:%+v \n", err, users)
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    
    users, err = model.NewUsersDao(context.Background(), db.Where("name = ?", "name01")).Gets()
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
}

func TestModelCount(t *testing.T) {
    db := GetGorm()
    
    var count int64
    
    db1 := model.NewUsersDao(context.Background(), db).Count(&count)
    fmt.Printf("count:%+v, db.err:%v \n", count, db1.Error)
    
    // 根据 option 条件统计数量
    usersDao := model.NewUsersDao(context.Background(), db)
    name1 := "name01"
    count = usersDao.GetCountByOptions(
        usersDao.WithName(&name1),
    )
    fmt.Println("count", count)
    name2 := "name010000"
    count = usersDao.GetCountByOptions(
        usersDao.WithName(&name2),
    )
    fmt.Println("count", count)
}

func TestModelCreate(t *testing.T) {
    db := GetGorm()
    
    // CreatedAt/UpdatedAt：设置值时覆盖，为0时会自动生成
    deletedAt := int32(1)
    name := "insert-prince03"
    users := model.Users{
        // ID:        0,
        Name: &name,
        // Age:       18,
        Age:       0,
        CardNo:    "46100001",
        HeadImg:   "https://dd.xx",
        CreatedAt: 1643399938,
        UpdatedAt: 0,
        DeletedAt: deletedAt,
    }
    rowsAffected, err := model.NewUsersDao(context.Background(), db).Create(&users)
    fmt.Printf("users:%+v rowsAffected:%d err:%v \n", users, rowsAffected, err)
}

// or 条件查询
func TestModelOr(t *testing.T) {
    db := GetGorm()
    userDao := model.NewUsersDao(context.Background(), db)
    
    var users []*model.Users
    var err error
    
    name := "name01"
    
    userCol := model.UsersColumns
    users, err = userDao.GetByOptions(
        userDao.WithSelect([]string{userCol.ID, userCol.Age}),
        userDao.WithID(1),
        userDao.WithOrOption(
            userDao.WithAge(18),
            userDao.WithName(&name),
        ),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
}

// 查询指定字段
func TestModelSelect(t *testing.T) {
    db := GetGorm()
    userDao := model.NewUsersDao(context.Background(), db)
    
    var user model.Users
    var users []*model.Users
    var err error
    userCol := model.UsersColumns
    // userAllCol := model.UsersAllColumns
    user, err = userDao.GetByOption(
        // userDao.WithSelect(userAllCol),
        userDao.WithSelect([]string{userCol.ID, userCol.Age}),
        // userDao.WithSelect(fmt.Sprintf("%s, %s", userCol.ID, userCol.Age)),
        // userDao.WithSelect(fmt.Sprintf("%s, sum(%s) AS age", userCol.ID, userCol.Age)),
        // userDao.WithSelect(fmt.Sprintf("IFNULL(%s, %d) AS age", userCol.Age, 100)),
        // userDao.WithSelect(fmt.Sprintf("IFNULL(%s, ?) AS age", userCol.Age), 100),
        // userDao.WithSelect(fmt.Sprintf("%s, IFNULL(%s, ?) AS age", userCol.ID, userCol.Age), 100),
        // userDao.WithSelect(fmt.Sprintf("%s, IF(%s, %s, ?) AS age", userCol.ID, userCol.Age, userCol.Age), 100),
        userDao.WithID(1),
    )
    fmt.Println("user, err:", user, err)
    users, err = userDao.GetByOptions(
        // userDao.WithSelect(fmt.Sprintf("%s, %s", userCol.ID, userCol.Age)),
        // userDao.WithSelect([]string{userCol.ID, userCol.Age}),
        // userDao.WithSelect(fmt.Sprintf("%s, sum(%s) AS age", userCol.ID, userCol.Age)),
        // userDao.WithSelect(fmt.Sprintf("IFNULL(%s, %d) AS age", userCol.Age, 100)),
        // userDao.WithSelect(fmt.Sprintf("IFNULL(%s, ?) AS age", userCol.Age), 100),
        // userDao.WithSelect(fmt.Sprintf("%s, IFNULL(%s, ?) AS age", userCol.ID, userCol.Age), 100),
        // userDao.WithSelect(fmt.Sprintf("%s, IF(%s, %s, ?) AS age", userCol.ID, userCol.Age, userCol.Age), 100),
        userDao.WithIDs([]int32{1, 2}),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
}

// 更新
func TestModelUpdateByOption(t *testing.T) {
    var user model.Users
    var err error
    var count int64
    
    db := GetGorm()
    userDao := model.NewUsersDao(context.Background(), db)
    
    user, err = userDao.GetByOption(
        userDao.WithID(1),
    )
    fmt.Printf("err:%v, users:%+v \n", err, user)
    
    // dtime := 0
    name := ""
    dtime := int32(1642337297)
    usesUpdate := model.Users{
        // ID: 100,
        Name:      &name,
        Age:       0, // 需要把年龄更新为0，注意 当通过 struct 更新时，GORM 只会更新非零字段。 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
        HeadImg:   "",
        DeletedAt: dtime,
    }
    count, err = userDao.UpdateByOption(
        usesUpdate,
        userDao.WithID(1),
    )
    fmt.Printf("err:%v, count:%d, users:%+v \n", err, count, usesUpdate)
    
    userCol := model.UsersColumns
    count, err = userDao.UpdateByOption(
        usesUpdate,
        userDao.WithSelect([]string{userCol.Name, userCol.Age, userCol.Age, userCol.HeadImg, userCol.DeletedAt}), // 更新指定字段
        userDao.WithID(1),
    )
    fmt.Printf("err:%v, count:%d, users:%+v \n", err, count, usesUpdate)
}

// 分组+筛选
func TestModelGetByOptionsOfGroup(t *testing.T) {
    db := GetGorm()
    
    var users []*model.Users
    var err error
    name := "name01"
    
    userDao := model.NewUsersDao(context.Background(), db)
    
    users, err = userDao.GetByOptions(
        userDao.WithName(&name),
        userDao.WithAge(12),
        userDao.WithOrOption(
            userDao.WithAge(18),
            userDao.WithName(&name),
        ),
        userDao.WithOrderBy(fmt.Sprintf("%s desc", model.UsersColumns.Name)),
        userDao.WithGroupBy(model.UsersColumns.Age),
        userDao.WithHaving(fmt.Sprintf("%s > ?", model.UsersColumns.Age), 15),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
}

// option 条件查询
func TestModelGetByOptions(t *testing.T) {
    db := GetGorm()
    
    var users []*model.Users
    var err error
    name := "name01"
    
    userDao := model.NewUsersDao(context.Background(), db)
    
    user, err := userDao.GetByOption(userDao.WithID(1))
    fmt.Println("user, err:", user, err)
    userDao = model.NewUsersDao(context.Background(), db)
    users, err = userDao.GetByOptions(userDao.WithIDs([]int32{1, 2}))
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    users, err = userDao.GetByOptions(userDao.WithID(1))
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    users, err = userDao.GetByOptions(
        userDao.WithName(&name),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
    users, err = userDao.GetByOptions(
        userDao.WithName(&name),
        userDao.WithAge(12),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
}

// 分页
func TestModelPage(t *testing.T) {
    db := GetGorm()
    
    var users []*model.Users
    var err error
    name := "name01"
    
    userDao := model.NewUsersDao(context.Background(), db)
    users, err = userDao.GetByOptions(
        userDao.WithName(&name),
        // userDao.WithName("name01"),
        // userDao.WithPage(0, 2),
        userDao.WithPage(1, 2),
    )
    for _, i2 := range users {
        fmt.Printf("err:%v, users:%+v \n", err, i2)
    }
}

// GetFromXxx 返回单条记录时，传入的参数为空值（0，""，nil）时会忽略为查询条件
func TestModelFrom(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var users []*model.Users
    var err error
    
    user, err = model.NewUsersDao(context.Background(), db).GetFromID(1)
    fmt.Println("GetFromID..user, err:", user, err)
    
    name := "ddd"
    users, err = model.NewUsersDao(context.Background(), db).GetFromName(&name)
    fmt.Println("GetFromName..users, err:", users, err)
    for _, i2 := range users {
        fmt.Printf("GetFromName..err:%v, users:%+v \n", err, i2)
    }
    
    var deletedAt int
    deletedAt = 0
    // deletedAt = 1
    users, err = model.NewUsersDao(context.Background(), db).GetFromDeletedAt(int32(deletedAt))
    for _, i2 := range users {
        fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, i2)
    }
    
    deletedAt1 := 1639411296
    deletedAt2 := 1639411297
    deletedAt3 := 0
    deletedAts := []int32{
        int32(deletedAt1),
        int32(deletedAt2),
        int32(deletedAt3),
    }
    users, err = model.NewUsersDao(context.Background(), db).GetsFromDeletedAt(deletedAts)
    fmt.Println("GetFromID..users, err:", users, err)
    
    user, err = model.NewUsersDao(context.Background(), db).GetFromID(10000)
    fmt.Println("GetFromID..user, err:", user, err)
    
    users, err = model.NewUsersDao(context.Background(), db).GetsFromID([]int32{1, 2})
    for _, i2 := range users {
        fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, i2)
    }
    
    name01 := "name01"
    name02 := "name01"
    users, err = model.NewUsersDao(context.Background(), db).GetFromName(&name01)
    fmt.Println("GetFromName..user, err:", user, err)
    users, err = model.NewUsersDao(context.Background(), db).GetsFromName([]*string{&name01, &name02})
    fmt.Println("GetFromName..user, err:", user, err)
    
    users, err = model.NewUsersDao(context.Background(), db).GetsFromID([]int32{1, 2})
    for _, i2 := range users {
        fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, i2)
    }
    
    user, err = model.NewUsersDao(context.Background(), db).GetFromCardNo("1")
    fmt.Println("GetFromCardNo..user, err:", user, err)
    
    users, err = model.NewUsersDao(context.Background(), db).GetsFromCardNo([]string{"1", "2"})
    for _, i2 := range users {
        fmt.Printf("GetBatchFromCardNo..err:%v, users:%+v \n", err, i2)
    }
    
}

// 通过索引获取数据
func TestModelFetch(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var users []*model.Users
    var err error
    
    user, err = model.NewUsersDao(context.Background(), db).FetchByPrimaryKey(1)
    fmt.Println("FetchByPrimaryKey..user, err:", user, err)
    
    user, err = model.NewUsersDao(context.Background(), db).FetchUniqueByCardNo("1ooo")
    fmt.Println("FetchUniqueByCardNo..user, err:", user, err)
    
    name01 := "name01"
    user, err = model.NewUsersDao(context.Background(), db).FetchUniqueIndexByUnqNameCard(&name01, "1")
    fmt.Println("FetchUniqueIndexByUnqNameCard..user, err:", user, err)
    
    users, err = model.NewUsersDao(context.Background(), db).FetchIndexByAge(120)
    for _, i2 := range users {
        fmt.Printf("FetchIndexByAge..err:%v, users:%+v \n", err, i2)
    }
}

// 重置连接
func TestModelReset(t *testing.T) {
    db := GetGorm()
    
    var user model.Users
    var err error
    
    userDao := model.NewUsersDao(context.Background(), db)
    
    name01 := "name01"
    name02 := "name02"
    user, err = userDao.GetByOption(
        // userDao.WithID(1),
        userDao.WithName(&name01),
        userDao.WithAge(18),
    )
    fmt.Println("userDao.GetByOption(userDao.WithID(1)):", user, err)
    
    user, err = userDao.GetByOption(userDao.WithName(&name02))
    fmt.Println("userDao.GetByOption(userDao.WithID(2)):", user, err)
    
}

// 支持事务便捷操作
func TestTracsaction(t *testing.T) {
    db := GetGorm()
    
    ctx := context.Background()
    
    var user model.Users
    var err error
    var rows int64
    
    // 1. 查询，更新；
    usersDao := model.NewUsersDao(ctx, db)
    user, err = usersDao.GetFromID(1)
    fmt.Println("+ model.NewUsersDao(ctx, db).GetFromID(1):", user, err)
    
    user.Age = 10
    rows, err = usersDao.UpdateByOption(user, usersDao.WithID(1))
    fmt.Println("+ usersDao.UpdateByOption(user, usersDao.WithID(1)):", rows, err)
    
    userBasseDao := model.NewUserBaseDao(ctx, db)
    rows, err = userBasseDao.Create(&model.UserBase{
        Name: "tt-01",
        Age:  10,
    })
    fmt.Println("+ userBasseDao.Save(&model.UserBase):", rows, err)
    
    // 2. 开启事务，查询并更新，提交或者回滚事务；
    tx := db.Begin()
    usersDao = model.NewUsersDao(ctx, tx)
    user, err = usersDao.GetFromID(1)
    fmt.Println("》model.NewUsersDao(ctx, db).GetFromID(1) tx:", user, err)
    
    user.Age = 20
    rows, err = usersDao.UpdateByOption(user, usersDao.WithID(1))
    fmt.Println("》usersDao.UpdateByOption(user, usersDao.WithID(1)) tx:", rows, err)
    if err != nil {
        tx.Rollback()
        fmt.Println("》tx.Rollback()", err)
        return
    }
    
    userBasseDao = model.NewUserBaseDao(ctx, tx)
    rows, err = userBasseDao.UpdateByOption(model.UserBase{
        Name: "tt-0101",
        Age:  10,
    }, userBasseDao.WithName("tt-01"))
    fmt.Println("》 userBasseDao.Save(&model.UserBase) tx:", rows, err)
    // err = errors.New("》》》 userBasseDao.Save to err")
    if err != nil {
        tx.Rollback()
        fmt.Println("》tx.Rollback()", err)
        return
    }
    
    tx.Commit()
    
    // 3. 再次查询，更新或插入
    // usersDao = model.NewUsersDao(ctx, usersDao.NewDB()) // `sql: transaction has already been committed or rolled back`
    // usersDao.New() // `sql: transaction has already been committed or rolled back`
    usersDao = model.NewUsersDao(ctx, db)
    
    user, err = usersDao.GetFromID(1)
    fmt.Println("+ model.NewUsersDao(ctx, db).GetFromID(1):", user, err)
    
    user.Age = 30
    rows, err = usersDao.UpdateByOption(user, usersDao.WithID(1))
    fmt.Println("+ usersDao.UpdateByOption(user, usersDao.WithID(1)):", rows, err)
    
    
}