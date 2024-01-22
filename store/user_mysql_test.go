package store

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"testing"
)

// GetNewDbMock 删除、创建、更新，都会开启事务，所以需要 mock 事务
// regexp.QuoteMeta 将字符串中的特殊字符转义
func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	gormDB, err := gorm.Open(mysql.Dialector{
		Config: &mysql.Config{
			DriverName:                "mysql",
			Conn:                      db,
			SkipInitializeWithVersion: true,
		},
	}, &gorm.Config{})
	if err != nil {
		return gormDB, mock, err
	}

	return gormDB, mock, err
}

// TestUserModel_List https://github.com/Watson-Sei/go-sqlmock-gorm/blob/main/main_test.go
func TestUserModel_List(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	userModel := NewUserModel(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age"}).
			AddRow(1, "test1", 18))
	users, err := userModel.List()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
}

func TestUserModel_GetUserById(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	userModel := NewUserModel(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE id = ?")).
		WithArgs(2).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "age"}).
			AddRow(2, "test1", 20))
	user, err := userModel.GetUserById(2)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(user, &User{Id: 2, Name: "test1", Age: 20}))
}

func TestUserModel_UpdateUser(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	userModel := NewUserModel(db)

	mock.ExpectBegin()
	// 只要前面 sql 语句匹配了，就会通过
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `users`")).WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()

	err = userModel.UpdateUser(&User{Id: 2, Name: "test1", Age: 30})
	assert.NoError(t, err)
}

func TestUserModel_CreateUser(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	userModel := NewUserModel(db)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`name`,`age`) VALUES (?,?)")).WithArgs("test1", 30).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = userModel.CreateUser(&User{Name: "test1", Age: 30})
	assert.NoError(t, err)
}

func TestUserModel_DeleteUser(t *testing.T) {
	db, mock, err := GetNewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	userModel := NewUserModel(db)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `users` WHERE id = ?")).WithArgs(2).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = userModel.DeleteUser(2)
	assert.NoError(t, err)
}
