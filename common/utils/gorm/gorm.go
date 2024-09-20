package gorm

import "gorm.io/gorm"

func CheckCommit(tx *gorm.DB, err *error) {
	if *err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}
