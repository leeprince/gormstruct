// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"github.com/leeprince/gormstruct/test/gormgentool/dao/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&model.UserInfo{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.UserInfo{}) fail: %s", err)
	}
}


func Test_userInfoSave(t *testing.T) {
	userInfo := newUserInfo(_gen_test_db)
	userInfo = *userInfo.As(userInfo.TableName())
	_do := userInfo.WithContext(context.Background()).Debug()
	
	primaryKey := field.NewString(userInfo.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <user_info> fail:", err)
		return
	}
	
	_, ok := userInfo.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from userInfo success")
	}
	
	err = _do.Save(&model.UserInfo{})
	if err != nil {
		t.Error("save item in table <user_info> fail:", err)
	}
	
	err = _do.Create(&model.UserInfo{})
	if err != nil {
		t.Error("create item in table <user_info> fail:", err)
	}
}

func Test_userInfoQuery(t *testing.T) {
	userInfo := newUserInfo(_gen_test_db)
	userInfo = *userInfo.As(userInfo.TableName())
	_do := userInfo.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(userInfo.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <user_info> fail:", err)
		return
	}

	_, ok := userInfo.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from userInfo success")
	}

	err = _do.Create(&model.UserInfo{})
	if err != nil {
		t.Error("create item in table <user_info> fail:", err)
	}

	err = _do.Save(&model.UserInfo{})
	if err != nil {
		t.Error("create item in table <user_info> fail:", err)
	}

	err = _do.CreateInBatches([]*model.UserInfo{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <user_info> fail:", err)
	}

	_, err = _do.Select(userInfo.ALL).Take()
	if err != nil {
		t.Error("Take() on table <user_info> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <user_info> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <user_info> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <user_info> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.UserInfo{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <user_info> fail:", err)
	}

	_, err = _do.Select(userInfo.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <user_info> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <user_info> fail:", err)
	}

	_, err = _do.Select(userInfo.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <user_info> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <user_info> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <user_info> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <user_info> fail:", err)
	}

	_, err = _do.ScanByPage(&model.UserInfo{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <user_info> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <user_info> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <user_info> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <user_info> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <user_info> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <user_info> fail:", err)
	}
}
