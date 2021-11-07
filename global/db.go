package global

import (
	"Blog/pkg/logger"
	"github.com/jinzhu/gorm"
)

var (
	DBEngine *gorm.DB
	Logger   *logger.Logger
)
