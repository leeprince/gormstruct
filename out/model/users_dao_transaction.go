package model

import (
    "fmt"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022-06-30 01:57:04
 * @Desc:   users 表的 DAO 层
 */

func (obj *UsersDAO) DBBeginTest() (err error) {
	// 考虑：该方法的DAO服务在外部是不是独立的（独立：重新初始化进来，并且不再外部公用该DAO服务）？
	// 	- 独立的：则不是必须在该方法结束后重新初始化DOA层中的DB
	// 	- 不独立的：则必须在该方法结束后重新初始化DOA层中的DB，以供外面该DAO服务使用。初始化的方式如下
	// 		```
	// 		方式一
	// 		initDB := obj.GetDB()
	// 		defer func() {
	// 			obj.UpdateDB(initDB)
	// 		}()
	// 		方式二
	// 		defer func() {
	// 			obj.UpdateDB(obj.db)
	// 		}()
	// 		```
	// 	> 考虑到使用便捷性：该方法结束后重新初始化DOA层中的DB
	
	defer func() {
		obj.UpdateDB(obj.db)
	}()

    tx := obj.Begin()
	fmt.Println(">>>>>>>>>>>>>>>>>开启事务")
    // 重新基于该事务的会话更新DAO服务的DB，保证DAO中DB使用事务的会话
    obj.UpdateDB(tx)
    
    user, err := obj.GetFromID(1)
    fmt.Println("GetFromID:", user, err)
    
    user.Age = 4
    rows, err := obj.UpdateByOption(user, obj.WithID(1))
    fmt.Println("UpdateByOption:", rows, err, user)
    if err != nil {
        fmt.Println("xxxxxxxxxxxxxx回滚事务1")
        tx.Rollback()
        fmt.Println("Rollback", err)
        return
    }
    
    // 验证sql正确性：user存在主键ID=2，必定执行错误并执行回滚
    // 包括测试事务是否回滚成功
    user.ID = 2
    rows, err = obj.UpdateByOption(user, obj.WithID(1))
    fmt.Println("UpdateByOption:", rows, err, user)
    if err != nil {
        fmt.Println("xxxxxxxxxxxxxx回滚事务2")
        tx.Rollback()
        fmt.Println("Rollback", err)
        return
    }
    
    fmt.Println("-----------------提交事务")
    tx.Commit()
    
    user, err = obj.GetFromID(1)
    fmt.Println("GetFromID:", user, err)
    
    return
}

func (obj *UsersDAO) ObjBeginTest() (err error) {
    obj.BeginTx()
	fmt.Println(">>>>>>>>>>>>>>>>>开启事务")
    
    user, err := obj.GetFromID(1)
    fmt.Println("GetFromID:", user, err)
    
    user.Age = 5
    rows, err := obj.UpdateByOption(user, obj.WithID(1))
    fmt.Println("UpdateByOption:", rows, err, user)
    if err != nil {
        fmt.Println("xxxxxxxxxxxxxx回滚事务1")
		obj.RollbackTx()
        fmt.Println("Rollback", err)
        return
    }
    
    // 验证sql正确性：user存在主键ID=2，必定执行错误并执行回滚
    // 包括测试事务是否回滚成功
    user.ID = 2
    rows, err = obj.UpdateByOption(user, obj.WithID(1))
    fmt.Println("UpdateByOption:", rows, err, user)
    if err != nil {
        fmt.Println("xxxxxxxxxxxxxx回滚事务2")
        obj.RollbackTx()
        fmt.Println("Rollback", err)
        return
    }
    
    fmt.Println("-----------------提交事务")
	obj.CommitTx()
    
    user, err = obj.GetFromID(1)
    fmt.Println("GetFromID:", user, err)
    
    return
}