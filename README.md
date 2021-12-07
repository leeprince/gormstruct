根据表名生成表结构体和数据操作方法
---

# 包含功能点
1. yaml 配置文件。包含数据库连接，表前缀
2. go run main.go ...支持指定表名、包名、表名对应结构体
3. 生成表的 gorm 结构体(对应文件名：{表名}.go)
4. 生成操作表的基本方法(对应文件名：gen.base.go)
5. 根据字段获取单条或者多条记录；根据主键、唯一、非唯一索引，获取单条或多条记录(对应文件名：{数据库名}.gen.{表名}.go)
    
### 指定表
```
go run main.go -t=users
```

# 快速使用
> 表名参数必填！
```
go run main.go -t=users [-p=model] [-s=Users]
go run main.go --table=users [--packageName=model] [--structName=Users]
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
  PRIMARY KEY (`id`),
  UNIQUE KEY `card_no` (`card_no`),
  UNIQUE KEY `unq_name_card` (`name`,`card_no`),
  KEY `age` (`age`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='用户表';
```
```
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`) VALUES (1, 'name01', 12, '1', '');
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`) VALUES (2, 'name02', 12, '2', '');
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`) VALUES (3, 'name01', 13, '3', '');
INSERT INTO `tmp`.`users` (`id`, `name`, `age`, `card_no`, `head_img`) VALUES (4, 'name01', 12, '4', '');
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