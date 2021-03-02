package message

import (
	"github.com/romberli/go-util/config"
)

const (
	// server
	InfoServerStart      = 200001
	InfoServerStop       = 200002
	InfoServerIsRunning  = 200003
	InfoServerNotRunning = 200004

	InfoMetadataGetEnvAll  = 200201
	InfoMetadataGetEnvByID = 200202
	InfoMetadataAddEnv     = 200203
	InfoMetadataUpdateEnv  = 200204

	InfoMetadataGetAppSystemAll  = 200301
	InfoMetadataGetAppSystemByID = 200302
	InfoMetadataAddAppSystem     = 200303
	InfoMetadataUpdateAppSystem  = 200304

	InfoMetadataGetUserAll  = 200901
	InfoMetadataGetUserByID = 200902
	InfoMetadataAddUser     = 200903
	InfoMetadataUpdateUser  = 200904

  InfoMetadataGetDbAll  = 200401
	InfoMetadataGetDbByID = 200402
	InfoMetadataAddDb    = 200403
	InfoMetadataUpdateDb  = 200404

	InfoMetadataGetMSAll  = 200601
	InfoMetadataGetMSByID = 200602
	InfoMetadataAddMS    = 200603
	InfoMetadataUpdateMS  = 200604
)

func initInfoMessage() {
	// server
	Messages[InfoServerStart] = config.NewErrMessage(DefaultMessageHeader, InfoServerStart, "das started successfully. pid: %d, pid file: %s")
	Messages[InfoServerStop] = config.NewErrMessage(DefaultMessageHeader, InfoServerStop, "das stopped successfully. pid: %d, pid file: %s")
	Messages[InfoServerIsRunning] = config.NewErrMessage(DefaultMessageHeader, InfoServerIsRunning, "das is running. pid: %d")
	Messages[InfoServerNotRunning] = config.NewErrMessage(DefaultMessageHeader, InfoServerNotRunning, "das is not running. pid: %d")

	Messages[InfoMetadataGetEnvAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetEnvAll, "metadata: get environment all completed")
	Messages[InfoMetadataGetEnvByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetEnvByID, "metadata: get environment by id completed. id: %s")
	Messages[InfoMetadataAddEnv] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddEnv, "metadata: add new environment completed. env_name: %s")
	Messages[InfoMetadataUpdateEnv] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateEnv, "metadata: update environment completed. id: %s")

	Messages[InfoMetadataGetAppSystemAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetAppSystemAll, "metadata: get appsystem all completed")
	Messages[InfoMetadataGetAppSystemByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetAppSystemByID, "metadata: get appsystem by id completed. id: %s")
	Messages[InfoMetadataAddAppSystem] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddAppSystem, "metadata: add new appsystem completed. system_name: %s")
	Messages[InfoMetadataUpdateAppSystem] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateAppSystem, "metadata: update appsystem completed. id: %s")

	Messages[InfoMetadataGetUserAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetUserAll, "metadata: get user all completed")
	Messages[InfoMetadataGetUserByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetUserByID, "metadata: get user by id completed. id: %s")
	Messages[InfoMetadataAddUser] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddUser, "metadata: add new user completed. env_name: %s")
	Messages[InfoMetadataUpdateUser] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateUser, "metadata: update user completed. id: %s")

  Messages[InfoMetadataGetDbAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetDbAll, "metadata: get database all completed")
	Messages[InfoMetadataGetDbByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetDbByID, "metadata: get database by id completed. id: %s")
	Messages[InfoMetadataAddDb] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddDb, "metadata: add new database completed. db_name and owner_id and env_id: %s")
	Messages[InfoMetadataUpdateDb] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateDb, "metadata: update database completed. id: %s")

	Messages[InfoMetadataGetMSAll] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMSAll, "metadata: get monitor systems all completed")
	Messages[InfoMetadataGetMSByID] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataGetMSByID, "metadata: get monitor system by id completed. id: %s")
	Messages[InfoMetadataAddMS] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataAddMS, "metadata: add new monitor system completed. system_name and system_type and host_ip and port_num and port_num_slow and base_url: %s")
	Messages[InfoMetadataUpdateMS] = config.NewErrMessage(DefaultMessageHeader, InfoMetadataUpdateMS, "metadata: update monitor system completed. id: %s")
}
