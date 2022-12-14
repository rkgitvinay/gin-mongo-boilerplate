package config

import (
	user "gin-mongo-boilerplate/src/user/models"

	. "github.com/shipu/artifact"
)

func Register() {
	Config.AddConfig("App", new(AppConfig))
	Config.AddConfig("DB", new(DatabaseConfig))
	Config.AddConfig("NoSql", new(MongoConfig))
	Config.Load()
}

func Boot() {
	user.UserSetup()
}
