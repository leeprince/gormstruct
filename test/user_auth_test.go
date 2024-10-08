package test

import (
	"context"
	"fmt"
	"github.com/leeprince/gormstruct/out/model"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/30 17:20
 * @Desc:
 */

func TestInsertData(t *testing.T) {
	db := InitDB()
	userAuthDao := model.NewUserAuthDAO(context.Background(), db)
	userauth := &model.UserAuth{
		ID:         0,
		UserID:     0,
		ExpireTime: time.Time{},
		AccessTime: 0,
		UpdatedAt:  time.Time{},
		CreatedAt:  time.Time{},
		DeletedAt:  nil,
	}
	row, err := userAuthDao.Save(userauth)
	fmt.Println(row, err)
	
}

func TestInsertDataCreateAt(t *testing.T) {
	db := InitDB()
	userAuthDao := model.NewUserAuthDAO(context.Background(), db)
	userauth := &model.UserAuth{
		ID:         0,
		UserID:     0,
		ExpireTime: time.Now(),
		AccessTime: 0,
		UpdatedAt:  time.Time{},
		CreatedAt:  time.Now(),
		DeletedAt:  nil,
	}
	row, err := userAuthDao.Save(userauth)
	fmt.Println(row, err)
}

func TestGetOne(t *testing.T) {
	db := InitDB()
	userAuthDao := model.NewUserAuthDAO(context.Background(), db)
	one, err := userAuthDao.GetByOption(userAuthDao.WithID(1))
	fmt.Println(one, err)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(one.CreatedAt.Format("2006-01-02 15:04:05"))
}

func TestGetList(t *testing.T) {
	db := InitDB()
	userAuthDao := model.NewUserAuthDAO(context.Background(), db)
	userAuthList, err := userAuthDao.GetListByOption(userAuthDao.WithIDs([]int64{1, 2}))
	fmt.Println(userAuthList, err)
	if err != nil {
		fmt.Println("err")
	}
	for _, auth := range userAuthList {
		fmt.Printf("auth: %+v \n", auth)
	}
}

func TestGetByOption(t *testing.T) {
	db := InitDB()
	userAuthDao := model.NewUserAuthDAO(context.Background(), db)
	
	var (
		err          error
		userAuth     *model.UserAuth
		userAuthList []*model.UserAuth
	)
	
	expireTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-08-31 09:55:34", time.Local)
	if err != nil {
		return
	}
	userAuth, err = userAuthDao.GetByOption(
		userAuthDao.WithExpireTime(expireTime),
	)
	fmt.Println("err:", err)
	fmt.Println("userAuth", userAuth)
	
	// updatedAtTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-08-30 23:35:08.000", time.Local)
	updatedAtTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-08-30 23:35:08.001", time.Local)
	if err != nil {
		return
	}
	userAuth, err = userAuthDao.GetByOption(
		userAuthDao.WithUpdatedAt(updatedAtTime),
	)
	fmt.Println("err:", err)
	fmt.Println("userAuth", userAuth)
	
	if err != nil {
		return
	}
	userAuthList, err = userAuthDao.GetListByOption(
		userAuthDao.WithUpdatedAt(updatedAtTime),
	)
	fmt.Println("err:", err)
	fmt.Println("userAuthList", userAuthList)
	for _, auth := range userAuthList {
		fmt.Printf("auth: %+v \n", auth)
	}
}

func TestDeletedAtNotNull(t *testing.T) {
	db := InitDB()
	userAuthDao := model.NewUserAuthDAO(context.Background(), db)
	
	var (
		err error
		// userAuth     *model.UserAuth
		userAuthList []*model.UserAuth
	)
	userAuthList, err = userAuthDao.GetListByOption(
		userAuthDao.WithIDs([]int64{
			1,
			3,
			5,
		}),
	)
	fmt.Println("err:", err)
	fmt.Println("userAuthList", userAuthList)
	for _, auth := range userAuthList {
		fmt.Printf("auth: %+v \n", auth)
	}
}

func TestGetCustomeResultByOption(t *testing.T) {
	db := InitDB()
	userAuthDao := model.NewUserAuthDAO(context.Background(), db)
	
	var (
		err error
	)
	
	type UserAuthSum struct {
		model.UserAuth
		SumId int64 `gorm:"column:sumUserId;" json:"sumUserId"` // 聚合字段:注意必须要添加`gorm:"column:xxxx;"`
	}
	
	const SumUserIdAsField = "sumUserId"
	selectArr := model.UserAuthAllColumns
	sum := fmt.Sprintf("SUM(%s) AS %s", model.UserAuthColumns.UserID, SumUserIdAsField)
	selectArr = append(selectArr, sum)
	
	options := []model.Option{
		userAuthDao.WithIDs([]int64{
			1, 2, 3, 4, 5,
		}),
		userAuthDao.WithGroupBy(model.UserAuthColumns.UserID),
	}
	
	options = append(options, userAuthDao.WithSelect(selectArr))
	
	result := &UserAuthSum{}
	err = userAuthDao.GetCustomeResultByOption(result,
		options...,
	)
	fmt.Println("err:", err)
	fmt.Println("result", *result)
	fmt.Printf("result: %+v \n", result)
	
	var resultList []*UserAuthSum
	err = userAuthDao.GetCustomeResultByOption(&resultList,
		options...,
	)
	fmt.Println("err:", err)
	fmt.Println("resultList", resultList)
	fmt.Printf("resultList: %+v \n", resultList)
	for _, authSum := range resultList {
		fmt.Printf("authSum: %+v \n", authSum)
	}
}
