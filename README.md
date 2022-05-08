根据表名生成表结构体和数据操作方法
---

# 包含功能点
1. yaml 配置文件。包含数据库连接、生成生成的基本方法版本、生成的逻辑方法版本
2. 命令行支持指定连接的数据库名表名、包名、表名对应结构体 `go run main.go -t={表名} [-d={数据库名(命令行设置会覆盖yaml配置的数据库名)}] [-p=包名] [-s=表名结构体]`
3. 生成表的 gorm 结构体，对应文件名：{表名}.go
4. 生成操作表的基本方法，对应文件名：base_dao.go
5. 根据结合 `WithXxxx` 指定值作为 option 条件或者根据单个字段、字段切片获取单条或者多条记录；根据主键、唯一、非唯一索引，获取单条或多条记录的方法，对应文件名：{表名}_dao.go)
6. 通过 `函数选项模式` （ `WithXxxx`）设置 option 条件，最后通过 `XxxxByOptions`(`GetByOption`/`GetByOptions`/`GetCountByOptions`/`UpdateByOption`) 方法实现查询、统计、更新   
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

# 注意
1. 当通过 struct 结构体更新时，GORM 只会更新非零字段。 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
2. 当使用 struct 结构体作为条件查询时，GORM 只会查询非零值字段。这意味着如果您的字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件
3. 创建数据时：CreatedAt/UpdatedAt：设置非零值时覆盖，为零值时会自动生成

# 快速使用
## 配置 ./config/config.yaml
```
dbinfo:
  host: 127.0.0.1
  port: 3306
  username: root
  password: leeprince
  database: tmp
  type: 0 # 数据库类型:0:mysql

# 生成的基本方法版本:V1,V2
gen_base_func_version: V2
# 生成的逻辑方法版本:V1,V2
gen_logic_func_version: V2
```

## 运行指令
> 表名参数必填！
```
go run main.go -t=users [-p=model] [-s=Users]
go run main.go --table=users [--packageName=model] [--structName=Users]
    -t/--table: 表名。表名必填！
    -p/--packageName: 包名
    -s/--structName: 表名对应结构体
```
```
./gormstruc -t=users [-p=model] [-s=Users]
./gormstruc --table=users [--packageName=model] [--structName=Users]
    -t/--table: 表名。表名必填！
    -p/--packageName: 包名
    -s/--structName: 表名对应结构体
```

# 包含sql查询
```
SELECT TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'tmp' AND TABLE_NAME = 'p_procedure';

SHOW KEYS FROM p_procedure;

SHOW FULL COLUMNS FROM p_procedure;
```

# 测试
## 测试表：users
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
## 测试运行的生成指令
生成的文件路径：out/model/*.go
```
go run main.go -t=users
```
## 单元测试
```
out/model_test.go
```

# 修改记录
- [x] 表的结构体的数据类型，根据表的数据类型、默认值、是否为 null、并根据配置是否使用指针类型，去最终设置为指针类型
    - 保持当前设置指针类型的条件：配置允许设置为指针类型&字段允许为null&字段的数据类型为uint|int|float|string
    - WithXxxx() 方法使用 `map[string]interface{}` 的形式支持所有数据类型
    - 更新数据通过更新结构体时，就算结构体的字段不是指针类型，想要更新为零值时，可配合 `WithSelect()` 更新
- [x] 表的 xxxColumns 放在 `dao` 层
- [x] 去除 `base_dao.go` 的 `updateOptionFunc` 方法
- [x] `FetchXxxx()` 方法依赖底层 `GetByOption()`、 `GetByOptions()` 方法
- [x] **DDD 架构**设计思想：`model` 层拆分为：`do(model)`层 + `dao` 层，外部通过 `repository` 调用 dao 层
    - 继续统一在 model 中，具体什么架构使用者自行移动即可，方便使用
- [x] 添加保存执行 sql 的日志到文件中的测试: `logWriterFile`
- [x] 生成的表结构体包含：结构体字段、结构体类型、表字段、表数据类型、表字段是否允许为null,表字段默认值、json字段
- [x] 支持配置是否控制台输出数据库表结构文档
- [x] 基础服务的包使用自定义的包 leeprince/goinfra
    - 保持当前
- [x] 支持事务便捷操作。
        场景操作步骤：
            1. 查询，更新或插入；
            2. 开启事务，查询并更新更新或插入，提交或者回滚事务；
            3. 再次查询，更新或插入。避免出现报错：`sql: transaction has already been committed or rolled back`
- [x] select 不指定的情况下取已生成的所有字段代替 `select *`
- [x] 优化生成的模型，满足 DDD 架构设计时对领域实体（模型）的设置