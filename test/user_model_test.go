package test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/leeprince/gormstruct/out/model"
	"testing"
	"time"
	
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2021/12/5 下午8:40
 * @Desc:
 */

func TestModelGetTableName(t *testing.T) {
	db := InitDB()
	
	userTableName := model.NewUsersDAO(context.Background(), db).GetTableName()
	fmt.Println("userTableName:", userTableName)
}

func TestModelCount(t *testing.T) {
	db := InitDB()
	
	var count int64
	
	// db1 := model.NewUsersDAO(context.Background(), db).Count(&count)
	// fmt.Printf("count:%+v, db.err:%v \n", count, db1.Error)
	
	// 根据 option 条件统计数量
	usersDAO := model.NewUsersDAO(context.Background(), db)
	name1 := "name01"
	count = usersDAO.GetCountByOption(
		usersDAO.WithName(name1),
	)
	fmt.Println("count", count)
	
	name2 := "name010000"
	count = usersDAO.GetCountByOption(
		usersDAO.WithName(name2),
	)
	fmt.Println("count", count)
}

// GetByOption 条件查询
func TestModelGetByOptionWithID(t *testing.T) {
	db := InitDB()
	
	fmt.Printf("--------------TestModelGetByOptionWithID \n\n")
	
	var err error
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	user, err := userDAO.GetByOption(
		userDAO.WithID(1),
	)
	
	fmt.Println(">>>>>>>>>>> 1 user, err:", user, err)
	
	user, err = userDAO.GetByOption(
		userDAO.WithID(0),
	)
	
	fmt.Println(">>>>>>>>>>> 2 user, err:", user, err)
}

// GetByOption 条件查询
func TestModelGetByOptionWithSelect(t *testing.T) {
	db := InitDB()
	
	fmt.Printf("--------------TestModelGetByOptionWithID \n\n")
	
	var err error
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	user, err := userDAO.GetByOption(
		userDAO.WithSelect([]string{
			model.UsersColumns.ID,
			model.UsersColumns.Age,
		}),
		userDAO.WithID(1),
	)
	fmt.Println(">>>>>>>>>>> 1 user, err:", user, err)
	
	user, err = userDAO.GetByOption(
		userDAO.WithSelect(fmt.Sprintf("%s, %s",
			model.UsersColumns.ID,
			model.UsersColumns.Age,
		)),
		userDAO.WithID(1),
	)
	fmt.Println(">>>>>>>>>>> 2 user, err:", user, err)
}

// GetByOption 条件查询
func TestCustomerSelectField(t *testing.T) {
	db := InitDB()
	
	fmt.Printf("--------------TestModelGetByOptionWithID \n\n")
	
	var err error
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	type userSumResult struct {
		*model.Users
		SumAge int64 `gorm:"column:sum_age" json:"sum_age"` // 新增的聚合字段。结合模型添加
	}
	userResult := &userSumResult{}
	err = userDAO.GetCustomeResultByOption(userResult,
		userDAO.WithSelect([]string{
			model.UsersColumns.ID,
			model.UsersColumns.Age,
			fmt.Sprintf("sum(%s) AS sum_age", model.UsersColumns.Age), // 聚合并且与userSumResult对应
		}),
		userDAO.WithID(1),
	)
	fmt.Println(">>>>>>>>>>> userResult, err:", userResult, err)
	
}

// GetByOption 条件查询
func TestModelGetByOption(t *testing.T) {
	db := InitDB()
	
	fmt.Printf("--------------TestModelGetByOption \n\n")
	
	var users []*model.Users
	var err error
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	user, err := userDAO.GetByOption(userDAO.WithID(1))
	fmt.Println(">>>>>>>>>>> 1 user, err:", user, err)
	
	user, err = userDAO.GetByOption(userDAO.WithID(1000))
	fmt.Println(">>>>>>>>>>> 1.1 user, err:", user, err)
	
	users, err = userDAO.GetListByOption(userDAO.WithIDs([]int64{1, 2}))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------ 2 err:%v, users:%+v \n", err, i2)
	}
	
	user, err = userDAO.GetByOption(userDAO.WithIDs([]int64{1, 2}))
	fmt.Println("=========== 3 user, err:", user, err)
}

// GetListByOption 条件查询
func TestModelGetByOptionWithWhereLike(t *testing.T) {
	db := InitDB()
	
	fmt.Printf("--------------TestModelGetByOption \n\n")
	
	var users []*model.Users
	var err error
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	users, err = userDAO.GetListByOption(userDAO.WithWhere("id >= 2"))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------1 err:%v, users:%+v \n", err, i2)
	}
	
	users, err = userDAO.GetListByOption(userDAO.WithWhere("name like '%a%'"))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------1.1 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	users, err = userDAO.GetListByOption(userDAO.WithWhere("name like ?", "%a%"))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------1.2 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	users, err = userDAO.GetListByOption(userDAO.WithWhere("name like ?", "%"+"a"+"%"))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------1.2.1 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	users, err = userDAO.GetListByOption(userDAO.WithWhere("name like ?", fmt.Sprintf("%%%s%%", "a")))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------1.3 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	users, err = userDAO.GetListByOption(userDAO.WithWhere(fmt.Sprintf("name like '%%%s%%'", "a")))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------1.4 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	// 推荐
	users, err = userDAO.GetListByOption(userDAO.WithWhere(fmt.Sprintf("%s like '%%%s%%'", model.UsersColumns.Name, "a")))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------1.5 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
}

// GetListByOption
func TestModelGetByOptionWithWhereList(t *testing.T) {
	db := InitDB()
	
	fmt.Printf("--------------TestModelGetByOption \n\n")
	
	var users []*model.Users
	var err error
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	users, err = userDAO.GetListByOption(userDAO.WithWhere("id >= ?", 2))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------1 err:%v, users:%+v \n", err, i2)
	}
	
	users, err = userDAO.GetListByOption(userDAO.WithWhere("id >= ? AND id <= ?", 2, 10))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------2 err:%v, users:%+v \n", err, i2)
	}
	
	opt := []model.Option{
		userDAO.WithWhere("id >= ? AND id <= ?", 2, 10),
	}
	users, err = userDAO.GetListByOption(opt...)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------3 err:%v, users:%+v \n", err, i2)
	}
	
	opt = []model.Option{
		userDAO.WithWhere("id >= ?", 2),
		userDAO.WithWhere("id <= ?", 10),
	}
	users, err = userDAO.GetListByOption(opt...)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------4 err:%v, users:%+v \n", err, i2)
	}
	
	opt = []model.Option{
		userDAO.WithWhere("id >= ? AND id <= ?", 2, 10),
		userDAO.WithWhere(fmt.Sprintf("name like '%%%s%%'", "name01")),
	}
	users, err = userDAO.GetListByOption(opt...)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------5 err:%v, users:%+v \n", err, i2)
	}
	
	users, err = userDAO.GetListByOption(userDAO.WithWhere("id in ? ", []int64{2, 10}))
	for _, i2 := range users {
		i2 := i2
		fmt.Printf("------6 err:%v, users:%+v \n", err, i2)
	}
}

// GetListByOption 条件查询
func TestModelGetListByOption(t *testing.T) {
	db := InitDB()
	
	var users []*model.Users
	var err error
	name := "name01"
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	users, err = userDAO.GetListByOption(userDAO.WithID(1))
	for _, i2 := range users {
		fmt.Printf("err:%v, users:%+v \n", err, i2)
	}
	users, err = userDAO.GetListByOption(
		userDAO.WithName(name),
	)
	for _, i2 := range users {
		fmt.Printf("err:%v, users:%+v \n", err, i2)
	}
	users, err = userDAO.GetListByOption(
		userDAO.WithName(name),
		userDAO.WithAge(12),
	)
	for _, i2 := range users {
		fmt.Printf("err:%v, users:%+v \n", err, i2)
	}
}

// or 条件查询
func TestModelOr(t *testing.T) {
	db := InitDB()
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	var users []*model.Users
	var err error
	
	name := "name01"
	userCol := model.UsersColumns
	
	// SELECT `id`,`age` FROM `users` WHERE `id` = 1 OR `age` = 18
	users, err = userDAO.GetListByOption(
		userDAO.WithSelect([]string{userCol.ID, userCol.Name, userCol.Age}),
		userDAO.WithID(1),
		userDAO.WithOrOption(
			userDAO.WithAge(18),
		),
	)
	for _, i2 := range users {
		fmt.Printf(">1 err:%v, users:%+v \n", err, i2)
	}
	
	// SELECT `id`,`name`,`age` FROM `users` WHERE `id` = 1 OR `name` = 'name01'
	users, err = userDAO.GetListByOption(
		userDAO.WithSelect([]string{userCol.ID, userCol.Name, userCol.Age}),
		userDAO.WithID(1),
		// 只会保留最后一个or
		userDAO.WithOrOption(
			userDAO.WithAge(18),
		),
		userDAO.WithOrOption(
			userDAO.WithName(name),
		),
	)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf(">2 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	// SELECT `id`,`name`,`age` FROM `users` WHERE `id` = 1 AND `name` = 'name01' OR `age` = 18
	users, err = userDAO.GetListByOption(
		userDAO.WithSelect([]string{userCol.ID, userCol.Name, userCol.Age}),
		userDAO.WithID(1),
		userDAO.WithOrOption(
			userDAO.WithAge(18),
		),
		userDAO.WithName(name),
	)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf(">3 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	// SELECT `id`,`name`,`age` FROM `users` WHERE `id` = 1 AND `name` = 'name01' OR `age` = 18
	users, err = userDAO.GetListByOption(
		userDAO.WithSelect([]string{userCol.ID, userCol.Name, userCol.Age}),
		userDAO.WithID(1),
		// 只会保留最后一个or
		userDAO.WithOrOption(
			userDAO.WithAge(18),
		),
		userDAO.WithName(name),
		userDAO.WithOrOption(
			userDAO.WithName(name),
		),
	)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf(">31 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	// SELECT `id`,`name`,`age` FROM `users` WHERE `name` = 'name01'
	users, err = userDAO.GetListByOption(
		userDAO.WithSelect([]string{userCol.ID, userCol.Name, userCol.Age}),
		// 只会保留最后一个or
		userDAO.WithOrOption(
			userDAO.WithID(1),
		),
		userDAO.WithOrOption(
			userDAO.WithAge(18),
		),
		userDAO.WithOrOption(
			userDAO.WithName(name),
		),
	)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf(">32 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	// SELECT `id`,`name`,`age` FROM `users`
	users, err = userDAO.GetListByOption(
		userDAO.WithSelect([]string{userCol.ID, userCol.Name, userCol.Age}),
		// 只会保留最后一个or
		userDAO.WithOrOption(
			userDAO.WithOrOption(
				userDAO.WithAge(18),
			),
			userDAO.WithOrOption(
				userDAO.WithName(name),
			),
		),
	)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf(">33 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	// SELECT `id`,`name`,`age` FROM `users` WHERE `id` = 1 AND `name` = 'name01' OR `age` = 18
	users, err = userDAO.GetListByOption(
		userDAO.WithSelect([]string{userCol.ID, userCol.Name, userCol.Age}),
		// 只会保留最后一个or
		userDAO.WithOrOption(
			userDAO.WithAge(18),
			userDAO.WithName(name),
		),
	)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf(">34 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
	
	// SELECT `id`,`age` FROM `users` WHERE `id` = 1 OR (`age` = 18 AND `name` = 'name01')
	users, err = userDAO.GetListByOption(
		userDAO.WithSelect([]string{userCol.ID, userCol.Name, userCol.Age}),
		userDAO.WithID(1),
		userDAO.WithOrOption(
			userDAO.WithAge(18),
			userDAO.WithName(name),
		),
	)
	for _, i2 := range users {
		fmt.Printf(">4 err:%v, users:%+v \n", err, i2)
	}
	
	// SELECT `id`,`name`,`age` FROM `users` WHERE `id` = 1 AND `name` = 'name01' OR (`age` = 18 AND `name` = 'name01')
	users, err = userDAO.GetListByOption(
		userDAO.WithSelect([]string{userCol.ID, userCol.Name, userCol.Age}),
		userDAO.WithID(1),
		userDAO.WithOrOption(
			userDAO.WithAge(18),
			userDAO.WithName(name),
		),
		userDAO.WithName(name),
	)
	for _, i2 := range users {
		i2 := i2
		fmt.Printf(">5 err:%v, users:%+v name:%s \n", err, i2, i2.Name)
	}
}

// 查询指定字段
func TestModelSelect(t *testing.T) {
	db := InitDB()
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	var user *model.Users
	var users []*model.Users
	var err error
	userCol := model.UsersColumns
	// userAllCol := model.UsersAllColumns
	user, err = userDAO.GetByOption(
		// userDAO.WithSelect(userAllCol),
		userDAO.WithSelect([]string{userCol.ID, userCol.Age}),
		// userDAO.WithSelect(fmt.Sprintf("%s, %s", userCol.ID, userCol.Age)),
		// userDAO.WithSelect(fmt.Sprintf("%s, sum(%s) AS age", userCol.ID, userCol.Age)),
		// userDAO.WithSelect(fmt.Sprintf("IFNULL(%s, %d) AS age", userCol.Age, 100)),
		// userDAO.WithSelect(fmt.Sprintf("IFNULL(%s, ?) AS age", userCol.Age), 100),
		// userDAO.WithSelect(fmt.Sprintf("%s, IFNULL(%s, ?) AS age", userCol.ID, userCol.Age), 100),
		// userDAO.WithSelect(fmt.Sprintf("%s, IF(%s, %s, ?) AS age", userCol.ID, userCol.Age, userCol.Age), 100),
		userDAO.WithID(1),
	)
	fmt.Println("user, err:", user, err)
	
	users, err = userDAO.GetListByOption(
		// userDAO.WithSelect(fmt.Sprintf("%s, %s", userCol.ID, userCol.Age)),
		userDAO.WithSelect([]string{userCol.ID, userCol.Age}),
		userDAO.WithSelect(fmt.Sprintf("%s, sum(%s) AS age", userCol.ID, userCol.Age)),
		// userDAO.WithSelect(fmt.Sprintf("IFNULL(%s, %d) AS age", userCol.Age, 100)),
		// userDAO.WithSelect(fmt.Sprintf("IFNULL(%s, ?) AS age", userCol.Age), 100),
		// userDAO.WithSelect(fmt.Sprintf("%s, IFNULL(%s, ?) AS age", userCol.ID, userCol.Age), 100),
		// userDAO.WithSelect(fmt.Sprintf("%s, IF(%s, %s, ?) AS age", userCol.ID, userCol.Age, userCol.Age), 100),
		// userDAO.WithIDs([]int64{1, 2}),
	)
	for _, i2 := range users {
		fmt.Printf("err:%v, users:%+v \n", err, i2)
	}
	
	aggratorData := struct {
		SumAge int64 `json:"sum_age,omitempty"`
	}{}
	err = userDAO.GetCustomeResultByOption(&aggratorData,
		// userDAO.WithSelect(fmt.Sprintf("%s, %s", userCol.ID, userCol.Age)),
		// userDAO.WithSelect([]string{userCol.ID, userCol.Age}),
		userDAO.WithSelect(fmt.Sprintf("sum(%s) AS sum_age", userCol.Age)),
		// userDAO.WithSelect(fmt.Sprintf("IFNULL(%s, %d) AS age", userCol.Age, 100)),
		// userDAO.WithSelect(fmt.Sprintf("IFNULL(%s, ?) AS age", userCol.Age), 100),
		// userDAO.WithSelect(fmt.Sprintf("%s, IFNULL(%s, ?) AS age", userCol.ID, userCol.Age), 100),
		// userDAO.WithSelect(fmt.Sprintf("%s, IF(%s, %s, ?) AS age", userCol.ID, userCol.Age, userCol.Age), 100),
		// userDAO.WithIDs([]int64{1, 2}),
	)
	fmt.Printf("err:%v, aggratorData:%+v \n", err, aggratorData)
}

func TestUpdateOrCreate(t *testing.T) {
	db := InitDB()
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	users := &model.Users{
		ID: 50,
		// Name: nil,
		// Age:       18,
		// Age:       0,
		CardNo: "32131da40",
		// HeadImg:   "https://dd.xx",
		// School:    nil,
		CreatedAt: 164339993,
		UpdatedAt: 1643399938,
		// DeletedAt: deletedAt,
	}
	rowsAffected, err := userDAO.UpdateOrCreate(users)
	fmt.Printf("users:%+v rowsAffected:%d err:%v \n", users, rowsAffected, err)
	
}
func TestModelSave(t *testing.T) {
	db := InitDB()
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	name := "insert-prince2"
	fmt.Printf("name:%+v \n", name)
	school := "insert-school"
	fmt.Printf("school:%+v \n", school)
	
	users := &model.Users{
		ID: 102,
		// Name: nil,
		// Age:       18,
		// Age:       0,
		CardNo: "32134da40",
		// HeadImg:   "https://dd.xx",
		// School:    nil,
		CreatedAt: 164339993,
		UpdatedAt: 1643399938,
		// DeletedAt: deletedAt,
	}
	rowsAffected := userDAO.Save(users)
	fmt.Printf("users:%+v rowsAffected:%d \n", users, rowsAffected)
	
	users.Name = ""
	users.School = ""
	rowsAffected, err := userDAO.UpdateOrCreate(users)
	fmt.Printf("users:%+v rowsAffected:%d err:%v \n", users, rowsAffected, err)
	
	time.Sleep(time.Second * 2)
	users.Age = 19
	users.UpdatedAt = 1643399938
	rowsAffected = userDAO.Save(users)
	// db.Save(users)
	fmt.Printf("users:%+v rowsAffected:%d \n", users, rowsAffected)
}

// 更新
func TestModelUpdate(t *testing.T) {
	var err error
	var count int64
	
	db := InitDB()
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	name := "insert-prince01"
	dtime := int64(1642337297)
	usesUpdate := &model.Users{
		// ID: 100,
		Name:      name,
		Age:       0, // 需要把年龄更新为0，注意 当通过表模型结构体更新时，GORM 只会更新非零字段。 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
		HeadImg:   "",
		DeletedAt: dtime,
	}
	count, err = userDAO.UpdateByOption(
		usesUpdate,
		userDAO.WithID(1),
	)
	fmt.Printf("err:%v, count:%d, users:%+v \n", err, count, usesUpdate)
	
	count, err = userDAO.UpdateByOption(
		usesUpdate,
		userDAO.WithID(2),
	)
	fmt.Printf("err:%v, count:%d, users:%+v \n", err, count, usesUpdate)
	
	userCol := model.UsersColumns
	count, err = userDAO.UpdateByOption(
		usesUpdate,
		userDAO.WithSelect([]string{userCol.Name, userCol.Age, userCol.HeadImg, userCol.DeletedAt}), // 更新指定字段
		userDAO.WithID(3),
	)
	fmt.Printf("err:%v, count:%d, users:%+v \n", err, count, usesUpdate)
}

// 多次更新
func TestModelMoreUpdate(t *testing.T) {
	var err error
	var count int64
	
	db := InitDB()
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	// usesUpdate, err := userDAO.GetFromID(1)
	usesUpdate, err := userDAO.GetFromCardNo("46000")
	if err != nil {
		fmt.Println("000000000 err:", err)
		return
	}
	fmt.Printf("= users:%+v \n", usesUpdate)
	
	fmt.Println(">>1 出现`WHERE `id` = 2 AND `id` = 1` 的问题")
	count, err = userDAO.UpdateByOption(
		usesUpdate,
		// userDAO.WithID(2),
		userDAO.WithCardNo("leeprince"),
	)
	fmt.Printf("err:%v, count:%d, users:%+v \n", err, count, usesUpdate)
	
	fmt.Println(">>>2 出现`WHERE `id` = 2 AND `id` = 1` 的问题")
	// usesUpdate, err = userDAO.GetFromID(2)
	count, err = userDAO.UpdateByOption(
		usesUpdate,
		// userDAO.WithID(2),
		userDAO.WithCardNo("leeprince"),
	)
	fmt.Printf("err:%v, count:%d, users:%+v \n", err, count, usesUpdate)
	
	fmt.Println(">>>>>>> 解决：出现`WHERE `id` = 2 AND `id` = 1` 的问题")
	// 分析：之前的查询条件查询之后，条件被保留在当前的userDAO中了
	// 解决：重新初始化userDAO
	//  但是这种解决方式需要每次都是重新初始化userDAO，有没有办法不需要重新初始化呢？
	//      考虑：dao 层每次执行sql初始化DB对应的模型`DB=db.Model(&Users{})`的话，外部通过UpdateDB()的方式会不会被`DB=db.Model(&Users{})`覆盖了
	//      最终解决：在`func (obj *_BaseDAO) GetDB() *gorm.DB {`每次初始化模型
	// 下列是手动解决方式，最终解决：在`func (obj *_BaseDAO) GetDB() *gorm.DB {`每次初始化模型
	// userDAO = model.NewUsersDAO(context.Background(), db) // 解决方式1
	// userDAO = model.NewUsersDAO(context.Background(), userDAO.GetDB()) // 解决方式2
	// userDAO.UpdateDB(userDAO.GetDB().Model(&model.Users{})) // 解决方式3
	// userDAO = model.NewUsersDAO(context.Background(), userDAO.NewDB()) // 解决方式4
	fmt.Printf("= users:%+v \n", usesUpdate)
	count, err = userDAO.UpdateByOption(
		usesUpdate,
		// userDAO.WithID(2),
		userDAO.WithCardNo("leeprince"),
	)
	fmt.Printf("err:%v, count:%d, users:%+v \n", err, count, usesUpdate)
}

// 分组+筛选
func TestModelGetListByOptionOfGroup(t *testing.T) {
	db := InitDB()
	
	var users []*model.Users
	var err error
	name := "name01"
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	users, err = userDAO.GetListByOption(
		userDAO.WithName(name),
		// userDAO.WithAge(12),
		userDAO.WithOrOption(
			userDAO.WithAge(18),
			// userDAO.WithName(name),
		),
		userDAO.WithOrderBy(fmt.Sprintf("%s desc", model.UsersColumns.Name)),
		userDAO.WithGroupBy(model.UsersColumns.Age),
		userDAO.WithHaving(fmt.Sprintf("%s > ?", model.UsersColumns.Age), 10),
	)
	for _, i2 := range users {
		fmt.Printf("err:%v, users:%+v \n", err, i2)
	}
}

// 分页
func TestModelPage(t *testing.T) {
	db := InitDB()
	
	var users []*model.Users
	var err error
	name := "name01"
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	users, err = userDAO.GetListByOption(
		userDAO.WithName(name),
		// userDAO.WithName("name01"),
		// userDAO.WithPage(0, 2),
		userDAO.WithPage(1, 2),
	)
	for _, i2 := range users {
		fmt.Printf("err:%v, users:%+v \n", err, i2)
	}
}

// GetFromXxx 返回单条记录时，传入的参数为空值（0，""，nil）时会忽略为查询条件
func TestModelFrom(t *testing.T) {
	db := InitDB()
	
	var user *model.Users
	var users []*model.Users
	var err error
	
	user, err = model.NewUsersDAO(context.Background(), db).GetFromID(1)
	fmt.Println("GetFromID..user, err:", user, err)
	
	user, err = model.NewUsersDAO(context.Background(), db).GetFromID(1000)
	fmt.Println("GetFromID..user, err:", user, err)
	
	name := "ddd"
	user, err = model.NewUsersDAO(context.Background(), db).GetFromName(name)
	fmt.Println("GetFromName..users, err:", users, err)
	for _, i2 := range users {
		fmt.Printf("GetFromName..err:%v, users:%+v \n", err, i2)
	}
	
	var deletedAt int
	deletedAt = 0
	// deletedAt = 1
	users, err = model.NewUsersDAO(context.Background(), db).GetFromDeletedAt(int64(deletedAt))
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
	users, err = model.NewUsersDAO(context.Background(), db).GetListFromDeletedAt(deletedAts)
	fmt.Println("GetFromID..users, err:", users, err)
	
	user, err = model.NewUsersDAO(context.Background(), db).GetFromID(10000)
	fmt.Println("GetFromID..user, err:", user, err)
	
	users, err = model.NewUsersDAO(context.Background(), db).GetListFromID([]int64{1, 2})
	for _, i2 := range users {
		fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, i2)
	}
	
	name01 := "name01"
	name02 := "name01"
	user, err = model.NewUsersDAO(context.Background(), db).GetFromName(name01)
	fmt.Println("GetFromName..user, err:", user, err)
	users, err = model.NewUsersDAO(context.Background(), db).GetListFromName([]string{name01, name02})
	fmt.Println("GetFromName..user, err:", user, err)
	
	users, err = model.NewUsersDAO(context.Background(), db).GetListFromID([]int64{1, 2})
	for _, i2 := range users {
		fmt.Printf("GetBatchFromID..err:%v, users:%+v \n", err, i2)
	}
	
	user, err = model.NewUsersDAO(context.Background(), db).GetFromCardNo("1")
	fmt.Println("GetFromCardNo..user, err:", user, err)
	
	users, err = model.NewUsersDAO(context.Background(), db).GetListFromCardNo([]string{"1", "2"})
	for _, i2 := range users {
		fmt.Printf("GetBatchFromCardNo..err:%v, users:%+v \n", err, i2)
	}
	
}

// 通过索引获取数据
func TestModelFetch(t *testing.T) {
	db := InitDB()
	
	var user *model.Users
	var users []*model.Users
	var err error
	
	user, err = model.NewUsersDAO(context.Background(), db).FetchByPrimaryKey(1)
	fmt.Println("FetchByPrimaryKey..user, err:", user, err)
	
	user, err = model.NewUsersDAO(context.Background(), db).FetchUniqueByCardNo("1ooo")
	fmt.Println("FetchUniqueByCardNo..user, err:", user, err)
	
	name01 := "name01"
	user, err = model.NewUsersDAO(context.Background(), db).FetchUniqueIndexByUnqNameCard(name01, "1")
	fmt.Println("FetchUniqueIndexByUnqNameCard..user, err:", user, err)
	
	users, err = model.NewUsersDAO(context.Background(), db).FetchIndexByAge(120)
	for _, i2 := range users {
		fmt.Printf("FetchIndexByAge..err:%v, users:%+v \n", err, i2)
	}
}

// 重置连接
func TestModelReset(t *testing.T) {
	db := InitDB()
	
	var user *model.Users
	var err error
	
	userDAO := model.NewUsersDAO(context.Background(), db)
	
	name01 := "name01"
	name02 := "name02"
	user, err = userDAO.GetByOption(
		// userDAO.WithID(1),
		userDAO.WithName(name01),
		userDAO.WithAge(18),
	)
	fmt.Println("userDAO.GetByOption(userDAO.WithID(1)):", user, err)
	
	user, err = userDAO.GetByOption(userDAO.WithName(name02))
	fmt.Println("userDAO.GetByOption(userDAO.WithID(2)):", user, err)
	
}

// 支持事务便捷操作
func TestTracsaction(t *testing.T) {
	db := InitDB()
	
	ctx := context.Background()
	
	var user *model.Users
	var err error
	var rows int64
	
	fmt.Println()
	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxx不重新初始化DB会报错：`sql: transaction has already been committed or rolled back`xxxxxxxxxxxxxxxxxxxxxxxx")
	usersDAO := model.NewUsersDAO(ctx, db)
	// 开启事务，查询并更新，提交或者回滚事务；
	tx := db.Begin()
	fmt.Println(">>>>>>>>>>>>>>>>>开启事务")
	usersDAO = model.NewUsersDAO(ctx, tx)
	user, err = usersDAO.GetFromID(1)
	fmt.Println("GetFromID tx:", user, err)
	
	user.Age = 1
	rows, err = usersDAO.UpdateByOption(user, usersDAO.WithID(1))
	fmt.Println("UpdateByOption:", rows, err, user)
	if err != nil {
		fmt.Println("xxxxxxxxxxxxxx回滚事务1")
		tx.Rollback()
		fmt.Println("Rollback", err)
		return
	}
	
	fmt.Println("-----------------提交事务")
	tx.Commit()
	
	// 再次查询，更新或插入
	user, err = usersDAO.GetFromID(1)
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("-=-=-=-=-=-=-=-err:", err)
	}
	fmt.Println("GetFromID(1):", user, err)
	
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++解决方式1：++++++++++++++++++++++++++++++")
	fmt.Println("在DAO层外开始事务")
	fmt.Println("使用`tx := db.Begin()`开启事务，tx传入DAO层操作DAO层方法：开始事务的DAO服务不与外面公用一个变量")
	usersDAO1Init := model.NewUsersDAO(ctx, db)
	// 开启事务，查询并更新，提交或者回滚事务；
	tx1 := db.Begin() // 开始事务之后，您应该使用 'tx' 而不是 'db'
	fmt.Println(">>>>>>>>>>>>>>>>>开启事务")
	usersDAO1 := model.NewUsersDAO(ctx, tx1) // 开始事务的DAO服务不与外面公用一个变量
	// usersDAO1 = model.NewUsersDAO(ctx, tx1) // 开始事务的DAO服务与外面公用一个变量
	user, err = usersDAO1.GetFromID(1)
	fmt.Println("GetFromID:", user, err)
	
	user.Age = 1
	rows, err = usersDAO1.UpdateByOption(user, usersDAO1.WithID(1))
	fmt.Println("UpdateByOption:", rows, err, user)
	if err != nil {
		fmt.Println("xxxxxxxxxxxxxx回滚事务1")
		tx1.Rollback()
		fmt.Println("Rollback", err)
		return
	}
	
	// // 验证sql正确性：user存在主键ID=2，必定执行错误并执行回滚
	// // 包括测试事务是否回滚成功
	// user.ID = 2
	// rows, err = usersDAO1.UpdateByOption(user, usersDAO1.WithID(1))
	// fmt.Println("UpdateByOption:", rows, err, user)
	// if err != nil {
	//     fmt.Println("xxxxxxxxxxxxxx回滚事务2")
	//     tx1.Rollback()
	//     fmt.Println("Rollback", err)
	//     return
	// }
	
	fmt.Println("-----------------提交事务")
	tx1.Commit()
	
	// 开始事务的DAO服务不与外面公用一个变量，所以可以继续使用事务前初始化的DAO服务
	// 再次查询，更新或插入
	user, err = usersDAO1Init.GetFromID(1)
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("-=-=-=-=-=-=-=-err:", err)
	}
	fmt.Println("GetFromID(1):", user, err)
	
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++解决方式2：++++++++++++++++++++++++++++++")
	fmt.Println("在DAO层外开始事务")
	fmt.Println("使用`tx := db.Begin()`开启事务，tx传入DAO层操作DAO层方法：开始事务的DAO服务与外面公用一个变量")
	usersDAO2Init := model.NewUsersDAO(ctx, db)
	user, err = usersDAO2Init.GetFromID(1)
	fmt.Println("GetFromID:", user, err)
	// 开启事务，查询并更新，提交或者回滚事务；
	tx2 := db.Begin() // 开始事务之后，您应该使用 'tx' 而不是 'db'
	fmt.Println(">>>>>>>>>>>>>>>>>开启事务")
	usersDAO2Init = model.NewUsersDAO(ctx, tx2) // 开始事务的DAO服务与外面公用一个变量
	user, err = usersDAO2Init.GetFromID(1)
	fmt.Println("GetFromID:", user, err)
	
	user.Age = 2
	rows, err = usersDAO2Init.UpdateByOption(user, usersDAO2Init.WithID(1))
	fmt.Println("UpdateByOption:", rows, err, user)
	if err != nil {
		fmt.Println("xxxxxxxxxxxxxx回滚事务1")
		tx2.Rollback()
		fmt.Println("Rollback", err)
		return
	}
	
	// // 验证sql正确性：user存在主键ID=2，必定执行错误并执行回滚
	// // 包括测试事务是否回滚成功
	// user.ID = 2
	// rows, err = usersDAO2Init.UpdateByOption(user, usersDAO2Init.WithID(1))
	// fmt.Println("UpdateByOption:", rows, err, user)
	// if err != nil {
	//     fmt.Println("xxxxxxxxxxxxxx回滚事务2")
	//     tx2.Rollback()
	//     fmt.Println("Rollback", err)
	//     return
	// }
	
	fmt.Println("-----------------提交事务")
	tx2.Commit()
	
	// 开始事务的DAO服务与外面公用一个变量。所以必须重新初始化DAO服务
	usersDAO2Init = model.NewUsersDAO(ctx, db) // 开始事务的DAO服务与外面公用一个变量
	
	// 再次查询，更新或插入
	user, err = usersDAO2Init.GetFromID(1)
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("-=-=-=-=-=-=-=-err:", err)
	}
	fmt.Println("GetFromID(1):", user, err)
	
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++解决方式3：++++++++++++++++++++++++++++++")
	fmt.Println("在DAO层外开始事务")
	fmt.Println("使用DAO层的事务管理。")
	usersDAO3Init := model.NewUsersDAO(ctx, db)
	user, err = usersDAO3Init.GetFromID(1)
	fmt.Println("GetFromID:", user, err)
	// 开启事务，查询并更新，提交或者回滚事务；
	usersDAO3Init.BeginTx() // 开始事务之后，您应该使用 'tx' 而不是 'db'
	fmt.Println(">>>>>>>>>>>>>>>>>开启事务")
	user, err = usersDAO3Init.GetFromID(1)
	fmt.Println("GetFromID:", user, err)
	
	user.Age = 3
	rows, err = usersDAO3Init.UpdateByOption(user, usersDAO3Init.WithID(1))
	fmt.Println("UpdateByOption:", rows, err, user)
	if err != nil {
		fmt.Println("xxxxxxxxxxxxxx回滚事务1")
		usersDAO3Init.RollbackTx()
		fmt.Println("Rollback", err)
		return
	}
	
	// // 验证sql正确性：user存在主键ID=2，必定执行错误并执行回滚
	// // 包括测试事务是否回滚成功
	// user.ID = 2
	// rows, err = usersDAO3Init.UpdateByOption(user, usersDAO3Init.WithID(1))
	// fmt.Println("UpdateByOption:", rows, err, user)
	// if err != nil {
	//     fmt.Println("xxxxxxxxxxxxxx回滚事务2")
	//     usersDAO3Init.RollbackTx()
	//     fmt.Println("Rollback", err)
	//     return
	// }
	
	fmt.Println("-----------------提交事务")
	usersDAO3Init.CommitTx()
	
	// 再次查询，更新或插入
	user, err = usersDAO3Init.GetFromID(1)
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("-=-=-=-=-=-=-=-err:", err)
	}
	fmt.Println("GetFromID(1):", user, err)
	
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++解决方式4：++++++++++++++++++++++++++++++")
	fmt.Println("在DAO层中开启事务")
	fmt.Println("使用`tx := db.Begin()`开启事务。")
	usersDAO4Init := model.NewUsersDAO(ctx, db)
	user, err = usersDAO4Init.GetFromID(1)
	fmt.Println("GetFromID:", user, err)
	
	// DAO服务在外部不是独立的
	fmt.Println("》DAO服务在外部不是独立的")
	usersDAO4Init = model.NewUsersDAO(ctx, db)
	err = usersDAO4Init.DBBeginTest()
	if err != nil {
		fmt.Println("DBBeginTest err:", err)
		// return // 取消注释，测试DAO层事务回滚之后后面的查询是否影响到
	}
	// 再次查询，更新或插入
	user, err = usersDAO4Init.GetFromID(1)
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("-=-=-=-=-=-=-=-err:", err)
	}
	fmt.Println("GetFromID(1):", user, err)
	
	// DAO服务在外部是独立的
	fmt.Println("》DAO服务在外部是独立的")
	usersDAO4_1Init := model.NewUsersDAO(ctx, db)
	err = usersDAO4_1Init.DBBeginTest()
	if err != nil {
		fmt.Println("DBBeginTest err:", err)
		// return // 取消注释，测试DAO层事务回滚之后后面的查询是否影响到
	}
	
	// 再次查询，更新或插入
	user, err = usersDAO4Init.GetFromID(1)
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("-=-=-=-=-=-=-=-err:", err)
	}
	fmt.Println("GetFromID(1):", user, err)
	
	fmt.Println()
	fmt.Println("++++++++++++++++++++++++++++++解决方式5：++++++++++++++++++++++++++++++")
	fmt.Println("在DAO层中开启事务")
	fmt.Println("使用DAO层的事务管理")
	usersDAO5Init := model.NewUsersDAO(ctx, db)
	user, err = usersDAO5Init.GetFromID(1)
	fmt.Println("GetFromID:", user, err)
	
	// DAO服务在外部不是独立的
	fmt.Println("》DAO服务在外部不是独立的")
	usersDAO5Init = model.NewUsersDAO(ctx, db)
	err = usersDAO5Init.ObjBeginTest()
	if err != nil {
		fmt.Println("DBBeginTest err:", err)
		// return // 取消注释，测试DAO层事务回滚之后后面的查询是否影响到
	}
	// 再次查询，更新或插入
	user, err = usersDAO5Init.GetFromID(1)
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("-=-=-=-=-=-=-=-err:", err)
	}
	fmt.Println("GetFromID(1):", user, err)
	
	// DAO服务在外部是独立的
	fmt.Println("》DAO服务在外部是独立的")
	usersDAO5_1Init := model.NewUsersDAO(ctx, db)
	err = usersDAO5_1Init.ObjBeginTest()
	if err != nil {
		fmt.Println("DBBeginTest err:", err)
		// return // 取消注释，测试DAO层事务回滚之后后面的查询是否影响到
	}
	
	// 再次查询，更新或插入
	user, err = usersDAO5Init.GetFromID(1)
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("-=-=-=-=-=-=-=-err:", err)
	}
	fmt.Println("GetFromID(1):", user, err)
}
