package helper

import "database/sql"

func CommietOrRolleback(tx *sql.Tx){
	err := recover()
	if err != nil {
		errolRolback := tx.Rollback()
		PanicError(errolRolback)
		panic( err)
	}else{
		errorCommit := tx.Commit()
		PanicError(errorCommit)
	}
}