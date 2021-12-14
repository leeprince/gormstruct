根据表名生成表结构体和数据操作方法
---

# 包含功能点
1. yaml 配置文件。
    > 包含数据库连接、生成生成的基本方法版本、生成的逻辑方法版本
2. 支持指定表名、包名、表名对应结构体
    > go run main.go -t={表名} [-p=包名] [-s=表名结构体]
3. 生成表的 gorm 结构体
    > 对应文件名：{表名}.go
4. 生成操作表的基本方法
    > 对应文件名：gen.base.go
5. 根据字段获取单条或者多条记录根据主键、唯一、非唯一索引，获取单条或多条记录
    > 对应文件名：{数据库名}.gen.{表名}.go)
6. 重置 gorm 会话; 
    > 要求配置：gen_logic_func_version >= V2
7. 支持表中的字段或者字段切片作为 option 条件
    > 要求配置：gen_logic_func_version >= V2
8. option 条件查询支持分页、支持统计记录总数
    > 要求配置：gen_logic_func_version >= V2
    > 支持分页的方法通过 `函数选项模式` 设置 option, 最后通过 `GetByOptions` 方法实现    

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
  `deleted_at` int(11) DEFAULT NULL COMMENT '删除 时间',
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