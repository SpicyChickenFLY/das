package metadata

import (
	"encoding/json"
	"fmt"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddlewareServerServiceAll(t *testing.T) {
	TestMiddlewareClusterService_GetEntities(t)
	TestMiddlewareClusterService_GetAll(t)
	TestMiddlewareClusterService_GetByID(t)
	TestMiddlewareClusterService_Create(t)
	TestMiddlewareClusterService_Update(t)
	TestMiddlewareClusterService_Delete(t)
	TestMiddlewareClusterService_Marshal(t)
	TestMiddlewareClusterService_MarshalWithFields(t)
}

func TestMiddlewareServerService_GetEntities(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMiddlewareServerService_GetAll(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetAll()
	asst.Nil(err, "test GetEntities() failed")
	entities := s.GetEntities()
	asst.Greater(len(entities), constant.ZeroInt, "test GetEntities() failed")
}

func TestMiddlewareServerService_GetByID(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetByID("1")
	asst.Nil(err, "test GetByID() failed")
	id := s.Entities[constant.ZeroInt].Identity()
	asst.Equal("1", id, "test GetByID() failed")
}

func TestMiddlewareServerService_Create(t *testing.T) {
	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.Create(map[string]interface{}{
		middlewareServerClusterIDStruct:      defaultMiddlewareServerInfoClusterID,
		middlewareServerNameStruct:           defaultMiddlewareServerInfoServerName,
		middlewareServerMiddlewareRoleStruct: defaultMiddlewareServerInfoMiddlewareRole,
		middlewareServerHostIPStruct:         defaultMiddlewareServerInfoSHostIP,
		middlewareServerPortNumStruct:        defaultMiddlewareServerInfoPortNum,
	})
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
	// delete
	err = deleteMiddlewareClusterByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Create() failed", err))
}

func TestMiddlewareServerService_Update(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	s := NewMiddlewareServerService(middlewareClusterRepo)
	err = s.Update(entity.Identity(), map[string]interface{}{middlewareServerNameStruct: newMiddlewareServerName})
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	err = s.GetByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	middlewareServerName, err := s.GetEntities()[constant.ZeroInt].Get(middlewareServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
	asst.Equal(newMiddlewareServerName, middlewareServerName)
	// delete
	err = deleteMiddlewareServerByID(s.Entities[0].Identity())
	asst.Nil(err, common.CombineMessageWithError("test Update() failed", err))
}

func TestMiddlewareServerService_Delete(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	s := NewMiddlewareServerService(middlewareServerRepo)
	err = s.Delete(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
	// delete
	err = deleteMiddlewareServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}

func TestMiddlewareServerService_Marshal(t *testing.T) {
	var entitiesUnmarshal []*MiddlewareServerInfo

	asst := assert.New(t)

	s := NewMiddlewareServerService(middlewareServerRepo)
	err := s.GetAll()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	data, err := s.Marshal()
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	err = json.Unmarshal(data, &entitiesUnmarshal)
	asst.Nil(err, common.CombineMessageWithError("test Marshal() failed", err))
	entities := s.GetEntities()
	for i := 0; i < len(entities); i++ {
		entity := entities[i]
		entityUnmarshal := entitiesUnmarshal[i]
		asst.True(middlewareServerStuctEqual(entity.(*MiddlewareServerInfo), entityUnmarshal), common.CombineMessageWithError("test Marshal() failed", err))
	}
}

func TestMiddlewareServerService_MarshalWithFields(t *testing.T) {
	asst := assert.New(t)

	entity, err := createMiddlewareServer()
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	s := NewMiddlewareServerService(middlewareServerRepo)
	err = s.GetByID(entity.Identity())
	dataService, err := s.MarshalWithFields(middlewareServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	dataEntity, err := entity.MarshalJSONWithFields(middlewareServerNameStruct)
	asst.Nil(err, common.CombineMessageWithError("test MarshalWithFields() failed", err))
	asst.Equal(string(dataService), fmt.Sprintf("[%s]", string(dataEntity)))
	// delete
	err = deleteMiddlewareServerByID(entity.Identity())
	asst.Nil(err, common.CombineMessageWithError("test Delete() failed", err))
}