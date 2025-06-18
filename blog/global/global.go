package global

import (
	"blog/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config *config.Config
	DB     *gorm.DB
	//Redis    *redis.Client
	Log      *logrus.Logger
	MysqlLog logger.Interface
)
