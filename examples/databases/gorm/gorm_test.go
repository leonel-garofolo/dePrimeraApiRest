package gorm

import (
	"deprimera/api/application"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMensajeAltaPrioridad(t *testing.T) {
	db, err := application.GetDB()
	if err != nil {
		log.Fatalln("fail to database connection")
	}
	defer db.Close()

	//dbSelect(db)
	//dbSelectOne(db)
	//dbInsert(db)
	dbInsertRecord(db)
	//dbUpdate(db)
	//dbDelete(db)

}
