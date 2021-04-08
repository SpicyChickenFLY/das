package metadata

import (
	"testing"

	"github.com/romberli/das/internal/dependency/metadata"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/stretchr/testify/assert"
)

const (
	onlineUserName = "online"
	newUserName    = "nun"
)

var userRepo = initUserRepo()

func initUserRepo() *UserRepo {
	pool, err := mysql.NewMySQLPoolWithDefault(addr, dbName, dbUser, dbPass)
	if err != nil {
		log.Error(common.CombineMessageWithError("initUserRepo() failed", err))
		return nil
	}

	return NewUserRepo(pool)
}

func createUser() (metadata.User, error) {
	userInfo := NewUserInfoWithDefault(
		defaultUserInfoUserName,
		defaultUserInfoDepartmentName,
		defaultUserInfoEmployeeID,
		defaultUserInfoAccountName,
		defaultUserInfoEmail,
		defaultUserInfoTelephone,
		defaultUserInfoMobile,
		defaultUserInfoRole,
	)
	entity, err := userRepo.Create(userInfo)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func deleteUserByID(id int) error {
	sql := `delete from t_meta_user_info where id = ?`
	_, err := userRepo.Execute(sql, id)
	return err
}

func TestUserRepoAll(t *testing.T) {
	TestUserRepo_Execute(t)
	TestUserRepo_GetAll(t)
	TestUserRepo_GetByID(t)
	TestUserRepo_Create(t)
	TestUserRepo_Update(t)
	TestUserRepo_Delete(t)
}

func TestUserRepo_Execute(t *testing.T) {
	asst := assert.New(t)

	sql := `select 1;`
	result, err := userRepo.Execute(sql)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	r, err := result.GetInt(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Execute() failed", err))
	asst.Equal(1, int(r), "test Execute() failed")
}

func TestUserRepo_Transaction(t *testing.T) {
	asst := assert.New(t)

	sql := `insert into t_meta_user_info(user_name,department_name,employee_id,account_name,email,telephone,mobile,role) values(?,?,?,?,?,?,?,?);`
	tx, err := userRepo.Transaction()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	err = tx.Begin()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	_, err = tx.Execute(sql,
		defaultUserInfoUserName,
		defaultUserInfoDepartmentName,
		defaultUserInfoEmployeeID,
		defaultUserInfoAccountName,
		defaultUserInfoEmail,
		defaultUserInfoTelephone,
		defaultUserInfoMobile,
		defaultUserInfoRole)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if inserted
	sql = `select user_name from t_meta_user_info where user_name=?`
	result, err := tx.Execute(sql,
		defaultUserInfoUserName,
	)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	userName, err := result.GetString(0, 0)
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	if userName != defaultUserInfoUserName {
		asst.Fail("test Transaction() failed")
	}
	err = tx.Rollback()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	// check if rollbacked
	entities, err := userRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
	for _, entity := range entities {
		userName := entity.GetUserName()
		asst.Nil(err, common.CombineMessageWithError("test Transaction() failed", err))
		if userName == defaultUserInfoUserName {
			asst.Fail("test Transaction() failed")
			break
		}
	}
}

func TestUserRepo_GetAll(t *testing.T) {
	asst := assert.New(t)

	entities, err := userRepo.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	userName := entities[0].GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetAll() failed", err))
	asst.Equal(onlineUserName, userName, "test GetAll() failed")
}

func TestUserRepo_GetByID(t *testing.T) {
	asst := assert.New(t)

	entity, err := userRepo.GetByID(2)
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	userName := entity.GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test GetByID() failed", err))
	asst.Equal(onlineUserName, userName, "test GetByID() failed")
}

func TestUserRepo_Create(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()

	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	id := entity.Identity()
	err = deleteUserByID(id)
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestUserRepo_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = entity.Set(map[string]interface{}{userNameStruct: newUserName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = userRepo.Update(entity)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	entity, err = userRepo.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	userName := entity.GetUserName()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newUserName, userName, "test Update() failed")
	// delete
	err = deleteUserByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestUserRepo_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createUser()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteUserByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}
