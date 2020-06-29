package daos

import (
	"deprimera/api/application"
	"deprimera/api/models"
	"fmt"
	"log"
)

type EquiposDao struct{}

func (ed *EquiposDao) GetAll() []models.Equipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	equipos := []models.Equipos{}
	db.Find(&equipos)
	return equipos
}

func (ed *EquiposDao) Get(id int) models.Equipos {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	equipo := models.Equipos{}
	db.Find(&equipo, id)
	return equipo
}

func (ed *EquiposDao) Save(e models.Equipos) int {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	equipoDB := db.Find(&e)
	if equipoDB == nil {
		db.Create(&e).Last(&e)
	} else {
		db.Save(&e)
	}
	return e.IDEquipo
}

func (ed *EquiposDao) Delete(id int) bool {
	equipo := models.Equipos{}

	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}
	db.Where("id_equipo = ?", id).First(&equipo)
	if equipo.IDLiga > 0 {
		db.Where("id_equipo=?", id).Delete(&models.Equipos{})
		fmt.Println("delete ID is:", id)
		return true
	} else {
		fmt.Println("no exist ID:", id)
		return false
	}
}

func (ed *EquiposDao) Query(sql string) *models.Equipos {
	equipos := &models.Equipos{}
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	db.First(equipos, 1)
	return equipos
}