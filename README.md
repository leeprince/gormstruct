# 根据表名生成表结构体和数据操作方法

> gorm 官方工具，请查阅：https://gorm.io/gen/gen_tool.html
---

# 一、包含功能点

1. yaml 配置文件。包含数据库连接、生成生成的基本方法版本、生成的逻辑方法版本
2.

命令行支持指定连接的数据库名表名、包名、表名对应结构体 `go run main.go -t={表名} [-d={数据库名(命令行设置会覆盖yaml配置的数据库名)}] [-p=包名] [-s=表名结构体]`

3. 生成表的 gorm 结构体，对应文件名：{表名}.go
4. 生成操作表的基本方法，对应文件名：base_dao.go
5. 根据结合 `WithXxxx` 指定值作为 option
   条件或者根据单个字段、字段切片获取单条或者多条记录；根据主键、唯一、非唯一索引，获取单条或多条记录的方法，对应文件名：{表名}_
   dao.go)
6. 通过 `函数选项模式` （ `WithXxxx`）设置 option 条件，最后通过 `XxxxByOptions`(`GetByOption`/`GetByOptions`
   /`GetCountByOptions`/`UpdateByOption`) 方法实现查询、统计、更新
7. 指定字段 `WithSelect`: 设置 option 条件作为查询指定字段或者更新指定字段的值
8. 排序 `WithOrderBy`: 设置 option 条件作为排序条件
9. 分组 `WithGroupBy`: 设置 option 条件作为分组条件
10. 分组筛选 `WithHaving`: 设置 option 条件作为分组筛选
11. SQL `OR` 条件 `WithOrOption`: 设置 option 条件组作为 sql `OR` 条件
12. 分页器 `WithPage`: 设置 option 条件作为 sql `offset`、`limit` 条件
13. 查找单条记录 `GetByOption`
14. 查找多条记录 `GetByOptions`
15. 统计 `Count`
16. 条件统计 `GetCountByOptions`: 设置 option 条件并统计
17. 插入 `Create`
18. 条件更新 `UpdateByOption`: 设置 option 条件作为更新条件，非指针的结构体字段更新为零值更新时需配合 `WithSelect()` 更新

# 二、注意

> 零值：0、''、false

## （一）更新

1. 当通过 struct 结构体更新时，
    - GORM Updates() 方法只会更新非零字段。 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
    - GORM Save() 方法会保存所有的字段，即使字段是零值(// prince@TODO: 添加该方法 2023/7/21 16:15)

## （二）查询

1. 当使用 struct 结构体作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他
   零值，该字段不会被用于构建查询条件

## (三)创建时间/创建时间

1. CreatedAt/UpdatedAt:
    - 创建数据时：CreatedAt/UpdatedAt：设置非零值时覆盖，为零值时会自动生成
    - 更新数据时：CreatedAt 不变；UpdatedAt 自动更新为当前时间戳

# 三、快速使用

## （一）配置 ./config/config.yaml

```
dbinfo:
  host: 127.0.0.1
  port: 3306
  username: root
  password: leeprince
  database: tmp
  type: 0 # 数据库类型:0:mysql

```

## （二）运行指令

### 帮助命令

> go run main.go --help

> go run main.go gen --help


**表名参数(-t)必填！**

```
go run main.go -t=users [-p=model] [-s=Users]
go run main.go --table=users [--packageName=model] [--structName=Users]
  -d, --database string      指定连接的数据库名
  -h, --help                 help for gen
  -p, --packageName string   生成的包名
  -s, --structName string    表名对应结构体，默认是表名的大驼峰命名
  -t, --table string         指定的表名
```

# 四、包含sql查询

```
SELECT TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'tmp' AND TABLE_NAME = 'p_procedure';

SHOW KEYS FROM p_procedure;

SHOW FULL COLUMNS FROM p_procedure;
```

# 五、测试

## （一）测试表：users

```
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '名称',
  `age` int(11) NOT NULL DEFAULT '18' COMMENT '年龄',
  `card_no` varchar(18) NOT NULL DEFAULT '' COMMENT '身份证',
  `head_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `created_at` int(11) NOT NULL COMMENT '创建时间',
  `updated_at` int(11) NOT NULL COMMENT '更新时间',
  `deleted_at` int(11) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `card_no` (`card_no`),
  UNIQUE KEY `unq_name_card` (`name`,`card_no`),
  KEY `age` (`age`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='用户表';
```

```
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'name01', 12, '1', '', 1639411296, 1639411296, 1639411296);
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'name02', 12, '2', '', 1639411296, 1639411296, 0);
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'name01', 13, '3', '', 1639411296, 1639411296, 0);
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'name01', 12, '4', '', 1639411296, 1639411296, 0);
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'name03', 18, '5', '', 1639411296, 1639411296, 0);
```

## （二）测试运行的生成指令

生成的文件路径：out/model/*.go

```
go run main.go -t=users
```

## （三）单元测试

```
out/model_test.go
```


# 六、gorm 官网 gentool 工具的使用

## (一) 安装

```
go install gorm.io/gen/tools/gentool@latest
```

## (二) 使用

```
gentool -h

Usage of gentool:
 -c string
       config file path 
 -db string
       input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html] (default "mysql")
 -dsn string
       consult[https://gorm.io/docs/connecting_to_the_database.html]
 -fieldNullable
       generate with pointer when field is nullable
 -fieldWithIndexTag
       generate field with gorm index tag
 -fieldWithTypeTag
       generate field with gorm column type tag
 -modelPkgName string
       generated model code's package name
 -outFile string
       query code file name, default: gen.go
 -outPath string
       specify a directory for output (default "./dao/query")
 -tables string
       enter the required data table or leave it blank
 -onlyModel
       only generate models (without query file)
 -withUnitTest
       generate unit test for query code
 -fieldSignable
       detect integer field's unsigned type, adjust generated data type
```

## (三) 例子

```
# 基础
gentool -dsn "user:pwd@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local" -tables "orders,doctor"

# 指定输出路径
gentool -dsn "user:pwd@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local" -tables "orders,doctor" -outPath=out/gentool/dao

# 生成用例
gentool -dsn "user:pwd@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local" -tables "orders,doctor" -outPath=out/gentool/dao -withUnitTest

# 指定配置文件
gentool -c "./gen.tool"
```

```
version: "0.1"
database:
  # consult[https://gorm.io/docs/connecting_to_the_database.html]"
  dsn : "username:password@tcp(address:port)/db?charset=utf8mb4&parseTime=true&loc=Local"
  # input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html]
  db  : "mysql"
  # enter the required data table or leave it blank.You can input : orders,users,goods
  tables  : "user"
  # specify a directory for output
  outPath :  "./dao/query"
  # query code file name, default: gen.go
  outFile :  ""
  # generate unit test for query code
  withUnitTest  : false
  # generated model code's package name
  modelPkgName  : ""
  # generate with pointer when field is nullable
  fieldNullable : false
  # generate field with gorm index tag
  fieldWithIndexTag : false
  # generate field with gorm column type tag
  fieldWithTypeTag  : false
```

# 七、修改记录

- [x] 表的结构体的数据类型，根据表的数据类型、默认值、是否为 null、并根据配置是否使用指针类型，去最终设置为指针类型
    - 保持当前设置指针类型的条件：配置允许设置为指针类型&字段允许为null&字段的数据类型为uint|int|float|string
    - WithXxxx() 方法使用 `map[string]interface{}` 的形式支持所有数据类型
    - 更新数据通过更新结构体时，就算结构体的字段不是指针类型，想要更新为零值时，可配合 `WithSelect()` 更新

- [x] 表的 xxxColumns 放在 `dao` 层

- [x] 去除 `base_dao.go` 的 `updateOptionFunc` 方法

- [x] `FetchXxxx()` 方法依赖底层 `GetByOption()`、 `GetByOptions()` 方法

- [x] **DDD 架构**设计思想：`model` 层拆分为：`do(model)`层 + `dao` 层，外部通过 `repository` 调用 dao 层
    - 继续统一在 model 中，通用性更强。具体什么架构，使用者自行移动即可

- [x] 添加保存执行 sql 的日志到文件中的测试: `logWriterFile`

- [x] 生成的表结构体包含：结构体字段、结构体类型、表字段、表数据类型、表字段是否允许为null,表字段默认值、json字段

- [x] 支持配置是否控制台输出数据库表结构文档

- [x] 基础服务的包使用自定义的包 leeprince/goinfra
    - 保持当前
- [x] 支持事务便捷操作。
- 问题场景操作步骤：

1. 查询，更新或插入；
2. 开启事务，查询并更新或插入，提交或者回滚事务；
3. 再次查询，更新或插入。出现报错：`sql: transaction has already been committed or rolled back`

- 原因：当前DAO层的DB已经提交事务，不允许再通过操作该DB

- 注意：开启事务之后，必须使用开启事务返回的*gorm.DB, 而不是开启事务时使用*gorm.DB

- 解决：!!!开始事务时，应基于**会话**的方式操作`*gorm.DB`。如：`tx := db.Begin()`
  开启事务后，必须使用tx而不是db操作事务中的sql!!!

- 在DAO层外开始事务

- - 使用`tx := db.Begin()`开启事务，tx传入DAO层操作DAO层方法

```
	tx := db.Begin()
	// 考虑：xxDAO 是否与事务外公用一个DAO服务？
	//  - xxDAO 不与事务外公用一个DAO服务，写法`xxDAO := model.NewXXXDAO(ctx, tx) `
	//  - xxDAO 与事务外公用一个DAO服务：因为该DAO服务是基于事务的会话开始的，事务结束后，当前会话（即：当前DAO服务的*gorm.DB）会失效
	//      - 如果开始事务的DAO服务与外面公用一个变量会写成`xxDAO = model.NewXXXDAO(ctx, tx) `(没有赋值给新的变量)，就需要在事务结束后，重新恢复xxDAO的`*gorm.DB`，而不是使用会话tx的DB，具体的重新初始化方式如下：
	//          - 重新初始化DAO层DB：`model.NewXXXDAO(ctx, db)`
	//          - 直接更新DAO层DB：`XXXDAO.UpdateDB(db)`
	//  > 推荐：xxDAO 不与事务外公用一个DAO服务,因为如果与外面公用一个变量，会使代码维护变得复杂
	xxDAO := model.NewXXXDAO(ctx, tx) // xxDAO 不与事务外公用一个DAO服务
	    
	xxDAO.XXX()
	...
	tx.Rollback()
	...
	    
	tx.Commit()
```

- - 使用DAO层的事务管理。

```
	xxDAO.BeginTx()
	    
	xxDAO.XXX()
	...
	xxDAO.RollbackTx()
	...

	xxDAO.CommitTx()
```

- 在DAO层中开启事务

- - 使用`tx := db.Begin()`开启事务

```
func (obj *XXXDAO) XXXX() {
    // 考虑：该方法的DAO服务在外部是不是独立的（独立：重新初始化进来，并且不在外部共用该DAO服务）？
    //  - 独立的：则不是必须在该方法结束后重新初始化DOA层中的DB
    //  - 不独立的：则必须在该方法结束后重新初始化DOA层中的DB，以供外面该DAO服务使用。初始化的方式如下
    //      ```
    //      方式一
    //      initDB := obj.GetDB()
    //      defer func() {
    //          obj.UpdateDB(initDB)
    //      }()
    //      方式二
    //      defer func() {
    //          obj.UpdateDB(obj.db)
    //      }()
    //      ```
    //  > 考虑到使用便捷性：该方法结束后重新初始化DOA层中的DB
    defer func() {
        obj.UpdateDB(obj.db)
    }()
    tx := obj.Begin()
  
    // 重新基于该事务的会话更新DAO服务的DB，保证DAO中DB使用事务的会话
    obj.UpdateDB(tx)
        
    obj.XXX()
    ...
    obj.Rollback()
    ...
        
    obj.Commit()
}
```

- - 使用DAO层的事务管理

```
func (obj *XXXDAO) XXXX() {
    obj.BeginTx()
        
    obj.XXX()
    ...
    obj.RollbackTx()
    ...
        
    obj.CommitTx()
}
```

> 注意点：并发场景时，请选择合适的方式！

- [x] select 不指定的情况下取已生成的所有字段代替 `select *`

- [x] 优化生成的模型，满足 DDD 架构设计时对领域实体（模型）的设置

- [x] 默认使用 `gorm.io/gorm` 库，并兼容测试 `github.com/jinzhu/gorm` 库
> [gorm.io/gorm]
- "非指针"的结构体字段更新为零值时，需配合 WithSelect 更新指定字段
- 存在 UpdatedAt 字段，并且 UpdatedAt 未传非零值时，会自动更新（默认：UpdatedAt==当前时间戳）。
- 参考：gorm.io/gorm@v1.22.4/schema/field.go@ParseField
- 取消自动更新 UpdatedAt 字段：通过 UpdateColumns() 方法（obj.prepare(opts...).UpdateColumns(&users)）
  [github.com/jinzhu/gorm]
- 更新零值，需配置字段为"指针"的数据类型。WithSelect 更新指定字段无效
- 更新必须指定 db.Table("table") 或者 db.Molde(&tableStruce{})
- 取消默认的 deleted_at IS NULL: db.Unscoped()...
- 取消自动更新 UpdatedAt 字段：通过 UpdateColumns() 方法（obj.prepare(opts...).UpdateColumns(&users)）
- 查询数据时，`err != nil && errors.Is(err, gorm.ErrRecordNotFound)` 的情况需兼容
  > ⚠️注意：github.com/jinzhu/gorm 使用本 dao 层需要删除部分方法，如：WithContext() 等

- [x] 表的结构体对象统一使用指针
1. 更新/插入时传参
2. 返回时

- [x] 生成的模型中包含获取主键的字段及方法 `PrimaryKey()` 及 `PrimaryKeyValue()`

- [x] 支持批量插入 `Create()`

- [x] 添加 `Save()` 方法：存在则更新，否则插入
- 通过`PrimaryKeyValue()`检查模型主键(如果非整型需要调整下，大部分数据库设计的主键都为整型，`Save()` 方法不使用反射降低性能)存在则更新

- [x] 支持表的模型对应的 *_dao.go 文件归并到 `dao` 目录下，并与 `model` 目录同级。
- 目的：将*_dao.go文件拷贝到业务项目的`dao`目录时，需要手动修改表的模型对应包的引用路径
- 无需项目支持，手动操作IDEA(goland)可实现目的.做法：
- - 1.从该项目拷贝生成的`模型文件（*.go）`及模型文件对应的`DAO文件（*_dao.go）`到业务项目的`model`目录下，
- - 2.手动迁移`DAO文件（*_dao.go）`，并利用IDEA(goland)的迁移`重构功能(Refactor)`，会自动补全model的包路径

- [x] 支持多次更新同一个模型
- 问题场景操作步骤：
- - 1.更新模型
- - 2.再次按条件更新同一模型,出现条件错误：再次更新模型的条件会自动加上上一次模型的主键，如：`WHERE id = 2 AND id = 1`
- 原因：上次模型的数据保存再DB中，模型的主键存在时，gorm会自动根据模型中的主键加上当前加上的条件，作为最后执行的sql
- 解决：在`base_dao.go`的`GetDB()`方法中重新初始化模型，这样DAO层每次执行的sql都是空的模型（非nil）

- [x] 允许自定义行记录的删除字段，存在删除字段时，`查询`方法默认会添加过滤已删除的记录

- [x] 允许通过`WithWhere`自定义查询条件

- [x] 因为使用`Or(options.queryMapOr)`所以目前仅支持一个OR条件；如需使用多个or条件，可以使用`WithWhere` 自定义查询条件；或者拆分为多条sql语句执行

- [x] 支持多条件参数绑定
```userDAO.GetByOptions(userDAO.WithWhere("id >= ? AND id <= ?", 2, 10))```

- [x] GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加添加 gorm:"column:字段的别名;" 标签)
>   注意：是 `Find(result)`, 而不是 `Find(&result)`
```
// 指定开始时间到当前时间的金额统计
func (r *MysqlRepo) SumTradingAmountFCEDataOfOption(opt FundChangeEventWhereOpt) (results *message.SumAmountMonthFCEData, err error) {
	options := r.MakeFundChangeEventWhereOpt(opt)

	selectArr := []string{model.FundChangeEventColumns.TradingTime}
	sum := fmt.Sprintf("SUM(%s) AS %s", model.FundChangeEventColumns.TradingAmount, message.SumTradingAmountAsField)
	selectArr = append(selectArr, sum)

	options = append(options, r.fundChangeEventDAO.WithSelect(selectArr))

	results = &message.SumAmountMonthFCEData{}
	err = r.fundChangeEventDAO.GetCustomeResultByOption(results,
		options...,
	)
	return
}

----
// 函数选项模式获取多条记录到自定义结构体中：支持包含自定义聚合字段
func (obj *FundChangeEventDAO) GetCustomeResultByOption(result interface{}, opts ...Option) (err error) {
    err = obj.prepare(opts...).Find(result).Error
    return
}
```

- [x] `GetByOptions` 改名为 `GetListByOption`

- [x] 解释是否需要 dao 层的上层是否需要 repository 层？答案：简单的CURD不需要，直接调用DAO层；复杂的需要repository层
```
使用 repo 的原因：model 和 dao 都可以生成，并且仅提供简单的方法.
- 不用行不行？：不用也可以，反正都是基于dao 层去做CURD，那就直接去在dao外面组装，但是建议仅简单CURL这么做，涉及条件判断是否需要添加option条件的还是需要接入repo 层;
- 用了之后带来的好处？：封装复杂CURD，并提高复用性。根据model的字段值做条件判断确定是否需要组装option数组
 ```
> repository 层设计参考
```
package mysqlrepo

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/15 18:39
 * @Desc:
 */

type FundChangeEventWhereOpt struct {
	model.FundChangeEvent
	TradingTimeRange *TradingTimeRange
}
type TradingTimeRange struct {
	TradingTimeStart int64
	TradingTimeEnd   int64
}

// 初始化 FundChangeEvent 有关的"查询条件"。其他的select/group by/order by等非查询或者聚合条件的都应该放在repository具体的方法中组装
func (r *MysqlRepo) MakeFundChangeEventWhereOpt(opt FundChangeEventWhereOpt) (options []dao.Option) {
	if opt.OrgID > 0 {
		options = append(options, r.bankHistoryMonthBudgetDAO.WithOrgID(opt.OrgID))
	}
	if opt.EventType != "" {
		options = append(options, r.fundChangeEventDAO.WithEventType(opt.EventType))
	}
	if opt.IncomeBankAccountName != "" {
		options = append(options, r.fundChangeEventDAO.WithIncomeBankAccountName(opt.IncomeBankAccountName))
	}
	if opt.ExpenditureBankAccountName != "" {
		options = append(options, r.fundChangeEventDAO.WithExpenditureBankAccountName(opt.ExpenditureBankAccountName))
	}
	// 注意：DAO 中只有一个 WithWhere 生效
	if opt.TradingTimeRange != nil {
		options = append(options, r.fundChangeEventDAO.WithWhere(fmt.Sprintf("%s >= ? AND %s <= ?",
			model.FundChangeEventColumns.TradingTime, model.FundChangeEventColumns.TradingTime),
			opt.TradingTimeRange.TradingTimeStart, opt.TradingTimeRange.TradingTimeEnd))
	}

	return
}

// 指定开始时间到当前时间的金额统计
func (r *MysqlRepo) SumTradingAmountFCEDataOfOption(opt FundChangeEventWhereOpt) (results *message.SumAmountMonthFCEData, err error) {
	options := r.MakeFundChangeEventWhereOpt(opt)

	selectArr := []string{model.FundChangeEventColumns.TradingTime}
	sum := fmt.Sprintf("SUM(%s) AS %s", model.FundChangeEventColumns.TradingAmount, message.SumTradingAmountAsField)
	selectArr = append(selectArr, sum)

	options = append(options, r.fundChangeEventDAO.WithSelect(selectArr))

	results = &message.SumAmountMonthFCEData{}
	err = r.fundChangeEventDAO.GetCustomeResultByOptions(results,
		options...,
	)
	return
}

// 分组：银行卡号；聚合：金额；排序金额
func (r *MysqlRepo) OrderByCurrentMonthFCEListByOption(opt FundChangeEventWhereOpt) (results []*message.SumAmountMonthFCEData, err error) {
	options := r.MakeFundChangeEventWhereOpt(opt)

	selectArr := model.FundChangeEventAllColumns
	sum := fmt.Sprintf("SUM(%s) AS %s", model.FundChangeEventColumns.TradingAmount, message.SumTradingAmountAsField)
	selectArr = append(selectArr, sum)

	groupField := model.FundChangeEventColumns.IncomeBankNumber
	if opt.EventType == constants.EnventEventTypeExpenditure {
		groupField = model.FundChangeEventColumns.ExpenditureBankNumber
	}

	options = append(options, r.fundChangeEventDAO.WithSelect(selectArr))
	options = append(options, r.fundChangeEventDAO.WithGroupBy(groupField))
	options = append(options, r.fundChangeEventDAO.WithOrderBy(message.SumTradingAmountAsField+" DESC"))

	results = []*message.SumAmountMonthFCEData{}
	err = r.fundChangeEventDAO.GetCustomeResultByOptions(&results,
		options...,
	)

	return
}

// 以租户+银行卡号分组，条件收支类型&&上一个月时间
func (r *MysqlRepo) GoupFCEOrgBankAndTradingAmount(opt FundChangeEventWhereOpt) (results []*message.SumAmountMonthFCEData, err error) {
	options := r.MakeFundChangeEventWhereOpt(opt)

	selectArr := model.FundChangeEventAllColumns
	sum := fmt.Sprintf("SUM(%s) AS %s", model.FundChangeEventColumns.TradingAmount, message.SumTradingAmountAsField)
	selectArr = append(selectArr, sum)

	// 按租户&银行卡号&收支类型分组
	groupField := fmt.Sprintf("%s, %s, %s",
		model.FundChangeEventColumns.OrgID,
		model.FundChangeEventColumns.EventType,
		model.FundChangeEventColumns.IncomeBankNumber,
	)

	options = append(options, r.fundChangeEventDAO.WithSelect(selectArr))
	options = append(options, r.fundChangeEventDAO.WithGroupBy(groupField))

	results = []*message.SumAmountMonthFCEData{}
	err = r.fundChangeEventDAO.GetCustomeResultByOptions(&results,
		options...,
	)

	return
}

// 事务操作：批量插入&更新已统计的记录
func (r *MysqlRepo) TxBatchInsertBudbetAndUpdateFCE(opt FundChangeEventWhereOpt, setFCEData *model.FundChangeEvent, batchBudget []*model.BankHistoryMonthBudget) (err error) {
	options := r.MakeFundChangeEventWhereOpt(opt)

	// 开启事务
	tx := app.GdcEnterpriseDigitalStatisticsDB.Begin()
	defer func() {
		if err != nil {
			fmt.Println(r.logID, "事务回滚!")
			gclog.WithError(err).Error(r.logID, "事务回滚!")
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	// 重新初始化dao
	fundChangeEventDAO := dao.NewFundChangeEventDAO(r.ctx, tx)
	bankHistoryMonthBudgetDAO := dao.NewBankHistoryMonthBudgetDAO(r.ctx, tx)

	// 批量插入
	_, err = bankHistoryMonthBudgetDAO.Create(batchBudget)
	if err != nil {
		gclog.WithError(err).Error(r.logID, "批量插入 error")
		return
	}

	_, err = fundChangeEventDAO.UpdateByOption(setFCEData, options...)
	if err != nil {
		gclog.WithError(err).Error(r.logID, "批量更新 error")
		return
	}

	return
}


```


