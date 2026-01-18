package helper

import (
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB, err *error) {
	if r := recover(); r != nil {
		tx.Rollback()
		panic(r)
	} else if *err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
