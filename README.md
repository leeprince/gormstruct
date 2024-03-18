# 根据表名生成模型和 DAO 层数据操作方法
---

# 一、快速使用

## （一）使用步骤
1. 克隆项目到本地（https://github.com/leeprince/gormstruct）
2. 重命名`./config/config.bak.yaml`为`./config/config.yaml`
3. 修改`dbinfo`为自己的数据库链接配置
```
# 数据库设置
dbinfo:
  host: 127.0.0.1
  port: 3306
  username: root
  password: root
  database: ticket # 该参数可通过命令行设置，命令行设置会覆盖该值。如果未设置则必须在命令中指定
  type: 0 # 数据库类型:0:mysql
```

4. 运行指令`go run main.go -t={数据库下的表名}`
会在`./out/model`下生成3个文件

- 模型：表名xxx.go
- DAO层：表名_DAO.go
- 基础DAO层：base.go（只有在第一次该工具时生成）

## （二）帮助命令
> go run main.go --help
> go run main.go gen --help

**表名参数(-t)必填！**

```
go run main.go -t=users [-p=model] [-s=Users]
go run main.go --table=users [--packageName=model] [--structName=Users]
  -d, --database string      指定连接的数据库名【在yaml配置时可选，命令设置优先级大于配置】
  -t, --table string         指定的表名【必填】
  -h, --help                 help for gen
  -p, --packageName string   生成的包名【可选，默认：model】
  -s, --structName string    表名对应结构体【可选，默认：表名的大驼峰命名】
```

> gorm 官网也提供 gentool 工具的使用。请查阅：https://gorm.io/gen/gen_tool.html
> 使用过后，不是我想要的，特别是 DAO 层的使用方式，感兴趣的自行查看

# 二、注意

> 零值：0、''、false

## （一）更新 Updates

当通过表模型结构体更新数据时，
- GORM Updates() 方法只会更新非零字段。该工具底层就是用了这个方法，如果您想确保指定字段被更新，你应该使用 Select 更新选定字段、或者允许为零值的字段设置为指针、或使用 map 来完成更新操作

> ⚠️注意：工具会根据是否允许为空，判断允许为空则设置为指针，反之不设置；如果表设计是不允许为空，但是有存在零值，只能手动改为指针

## （一）保存或者更新 Save
gorm 原生的 Save 是一个组合函数。 如果保存值不包含主键，它将执行 Create，否则它将执行 Update (包含所有字段)

> ⚠️注意：Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建。

- 保存时：字段是指针（允许为空）
- - 字段有非零值默认值：保存为默认值
- - 字段有零值或者无默认值：因为是指针，不会保存零值，只会保存为NULL #细节待验证

- 更新时：字段是指针（允许为空）
- - 字段有非零值默认值： 保存为默认值 #细节待验证
- - 字段有零值或者无默认值： 因为是指针，不会保存零值，只会保存为NULL #细节待验证

## （三）查询

当使用表模型结构体作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他零值，该字段不会被用于构建查询条件。
但通过该工具生成的DAO层`WithXxx`和`WithXxxs`作为条件则可以忽略这个，因为底层是用 map 

## （四）创建时间/创建时间
CreatedAt/UpdatedAt:
- 创建数据时：CreatedAt/UpdatedAt：设置非零值时覆盖，为零值时会自动生成
- 更新数据时：CreatedAt 不变；UpdatedAt 自动更新为当前时间戳

# 三、包含功能点

> 自动判断是否存在软删除字段，存在则默认过滤非软删除记录

## （一）关于模型

1. 生成模型 ：生成模型，包含每个字段的注释，注释统一在右边//对齐，方便全局观察字段的定义情况。
2. 生成表的所有字段映射：使用字段自定义查询时，统一使用映射的字段，在字段的维护上更方便；同时默认查询所有字段，不使用*，不使用反射查询所有字段。
3. 为所有字段生成`SetXxx()`、`GetXxx()`方法：用户设置/获取模型下指定字段的值

## （二）关于DAO层
生成操作表的基本方法，对应文件名：base_DAO.go
对应的表 DAO 层：xxx_dao.go 里面包含多种操作表数据的简单、易用、易维护方法。

- 生成`UpdateOrCreate()`方法：存在则更新（只会更新非零值），否则插入（会忽略零值字段）
- 生成`Save()`方法：gorm 原生的 Save 会保存所有的字段，即使字段是零值。仅会判断主键是否存在，存在则更新，不存在则创建
- 生成`Create()`方法：创建数据:允许单条/批量创建，批量创建时传入切片

- 为所有字段生成 `WithXxx`方法：将单个字段的值作为查询条件
- 为所有字段生成 `WithXxxs`方法（多一个s）：将单个字段的列表值作为查询条件

- 生成`GetByOption()`方法：函数选项模式获取单条记录
- 生成`GetListByOption`方法：函数选项模式获取多条记录：支持分页

- 生成`GetCountByOption()`方法：函数选项模式获取多条记录：支持统计记录总数

- 生成`GetCustomeResultByOption()`方法：函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名；" 标签)

- 生成`UpdateByOption()`：更新数据

- 生成`GetFromXxx()`方法：简单直接通过单个字段值获取数据（自动判断该字段是否设置唯一索引，否则自动获取单条记录，反之则是多条记录）
- 生成`GetListFromXxx()`方法：简单直接通过单个字段的列表值获取多条记录

- 生成`FetchByXxx()`方法：通过索引（唯一索引（主键、唯一索引、唯一复合索引）、非唯一索引（普通索引））作为查询条件获取数据（自动判断索引类型确定获取单条记录还是多条记录）

> 所有的方法：在项目中都包含相应的测试用例`test/xxx_test.go`


# 四、包含sql查询
```
SELECT TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'tmp' AND TABLE_NAME = 'p_procedure'；

SHOW KEYS FROM p_procedure；

SHOW FULL COLUMNS FROM p_procedure；
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='用户表'；
```

```
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'name01', 12, '1', '', 1639411296, 1639411296, 1639411296)；
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'name02', 12, '2', '', 1639411296, 1639411296, 0)；
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'name01', 13, '3', '', 1639411296, 1639411296, 0)；
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'name01', 12, '4', '', 1639411296, 1639411296, 0)；
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'name03', 18, '5', '', 1639411296, 1639411296, 0)；
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


# 六、修改记录

- [x] 表的结构体的数据类型，根据表的数据类型、默认值、是否为 null、并根据配置是否使用指针类型，去最终设置为指针类型
    - 保持当前设置指针类型的条件：配置允许设置为指针类型&字段允许为null&字段的数据类型为uint|int|float|string
    - WithXxxx() 方法使用 `map[string]interface{}` 的形式支持所有数据类型
    - 更新数据通过更新结构体时，就算结构体的字段不是指针类型，想要更新为零值时，可配合 `WithSelect()` 更新

- [x] 表的 xxxColumns 放在 `DAO` 层

- [x] 去除 `base_DAO.go` 的 `updateOptionFunc` 方法

- [x] `FetchXxxx()` 方法依赖底层 `GetByOption()`、 `GetListByOption()` 方法

- [x] **DDD 架构**设计思想：`model` 层拆分为：`do(model)`层 + `DAO` 层，外部通过 `repository` 调用 DAO 层
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
  > ⚠️注意：github.com/jinzhu/gorm 使用本 DAO 层需要删除部分方法，如：WithContext() 等

- [x] 表的结构体对象统一使用指针
1. 更新/插入时传参
2. 返回时

- [x] 生成的模型中包含获取主键的字段及方法 `PrimaryKey()` 及 `PrimaryKeyValue()`

- [x] 支持批量插入 `Create()`

- [x] 添加 `Save()` 方法：存在则更新，否则插入
- 通过`PrimaryKeyValue()`检查模型主键(如果非整型需要调整下，大部分数据库设计的主键都为整型，`Save()` 方法不使用反射降低性能)存在则更新

- [x] 支持表的模型对应的 *_DAO.go 文件归并到 `DAO` 目录下，并与 `model` 目录同级。
- 目的：将*_DAO.go文件拷贝到业务项目的`DAO`目录时，需要手动修改表的模型对应包的引用路径
- 无需项目支持，手动操作IDEA(goland)可实现目的.做法：
- - 1.从该项目拷贝生成的`模型文件（*.go）`及模型文件对应的`DAO文件（*_DAO.go）`到业务项目的`model`目录下，
- - 2.手动迁移`DAO文件（*_DAO.go）`，并利用IDEA(goland)的迁移`重构功能(Refactor)`，会自动补全model的包路径

- [x] 支持多次更新同一个模型
- 问题场景操作步骤：
- - 1.更新模型
- - 2.再次按条件更新同一模型,出现条件错误：再次更新模型的条件会自动加上上一次模型的主键，如：`WHERE id = 2 AND id = 1`
- 原因：上次模型的数据保存再DB中，模型的主键存在时，gorm会自动根据模型中的主键加上当前加上的条件，作为最后执行的sql
- 解决：在`base_DAO.go`的`GetDB()`方法中重新初始化模型，这样DAO层每次执行的sql都是空的模型（非nil）

- [x] 允许自定义行记录的删除字段，存在删除字段时，`查询`方法默认会添加过滤已删除的记录

- [x] 允许通过`WithWhere`自定义查询条件

- [x] 因为使用`Or(options.queryMapOr)`所以目前仅支持一个OR条件；如需使用多个or条件，可以使用`WithWhere` 自定义查询条件；或者拆分为多条sql语句执行

- [x] 支持多条件参数绑定
```userDAO.GetByOptions(userDAO.WithWhere("id >= ? AND id <= ?", 2, 10))```

- [x] GetCustomeResultByOption 函数选项模式获取多条记录到自定义结构体(result:务必使用指针变量)：支持包含自定义聚合字段(自定义的聚合字段务必添加 gorm:"column:字段的别名；" 标签)
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

- [x] 解释是否需要 DAO 层的上层是否需要 repository 层？答案：简单的CURD不需要，直接调用DAO层；复杂的需要repository层
```
使用 repo 的原因：model 和 DAO 都可以生成，并且仅提供简单的方法.
- 不用行不行？：不用也可以，反正都是基于DAO 层去做CURD，那就直接去在DAO外面组装，但是建议仅简单CURL这么做，涉及条件判断是否需要添加option条件的还是需要接入repo 层；
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
func (r *MysqlRepo) MakeFundChangeEventWhereOpt(opt FundChangeEventWhereOpt) (options []DAO.Option) {
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
	// 重新初始化DAO
	fundChangeEventDAO := DAO.NewFundChangeEventDAO(r.ctx, tx)
	bankHistoryMonthBudgetDAO := DAO.NewBankHistoryMonthBudgetDAO(r.ctx, tx)

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

- [x] 当时时间格式未使用整型时间戳，而是 timestamp/datatiem 时：结构体字段类型使用 *time.Time，而不是 time.Time, 避免初始化变量或者数据json转化后使用time.Time的默认值：`0001-01-01 00:00:00 +0000` 插入数据，导致查询删除字段是否为空错误
另外一种方案是使用：gorm.DeletedAt 字段类型---本库未使用

- [x] UpdateOrCreate 方法中，Fixed #1：有些表的主键不是以`id`命名，但是在此建议非常量定义，都需要建好自增主键ID，这个与存储引擎的数据结构有关。

---

### 如果您觉得本工具对您有帮助，不妨在右上角点亮一颗小星星，以示鼓励！

也欢迎关注我的公众号

![加我微信](./公众号《皇子谈技术》.png)