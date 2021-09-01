package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logdb "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

/**
 * @author : Donald Trieu
 * @created : 9/1/21, Wednesday
**/

func NewConnection(conStr string) *gorm.DB {
	logger := logdb.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logdb.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logdb.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conStr,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic(err)
	}
	return db
}
