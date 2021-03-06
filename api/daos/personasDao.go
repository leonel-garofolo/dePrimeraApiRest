package daos

import (
	"database/sql"
	"log"

	"github.com/leonel-garofolo/dePrimeraApiRest/api/application"
	"github.com/leonel-garofolo/dePrimeraApiRest/api/daos/gorms"
)

type PersonasDaoImpl struct{}

func (ed *PersonasDaoImpl) GetAll() []gorms.PersonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	rows, err := db.Query("select id_persona, nombre, apellido, domicilio, edad, localidad, id_pais, id_provincia, id_tipo_doc, nro_doc from personas")
	if err != nil {
		log.Fatalln("Failed to query")
	}

	apellido := sql.NullString{}
	personas := []gorms.PersonasGorm{}
	for rows.Next() {
		persona := gorms.PersonasGorm{}
		error := rows.Scan(&persona.IDPersona, &persona.Nombre, &apellido, &persona.Domicilio, &persona.Edad, &persona.Localidad, &persona.IDPais, &persona.IDProvincia, &persona.IDTipoDoc, &persona.NroDoc)
		if error != nil {
			if error != sql.ErrNoRows {
				log.Println(error)
				panic(error)
			}
		}
		persona.Apellido = apellido.String
		personas = append(personas, persona)
	}
	return personas
}

func (ed *PersonasDaoImpl) Get(id int) gorms.PersonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select id_persona, nombre, apellido, domicilio, edad, localidad, id_pais, id_provincia, id_tipo_doc, nro_doc from personas where id_persona = ?", id)
	persona := gorms.PersonasGorm{}
	error := row.Scan(&persona.IDPersona, &persona.Nombre, &persona.Apellido, &persona.Domicilio, &persona.Edad, &persona.Localidad, &persona.IDPais, &persona.IDProvincia, &persona.IDTipoDoc, &persona.NroDoc)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return persona
}

func (ed *PersonasDaoImpl) GetPersonasFromUser(idUser string) gorms.PersonasGorm {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	row := db.QueryRow("select id_persona, nombre, apellido, domicilio, edad, localidad, id_pais, id_provincia, id_tipo_doc, nro_doc from personas where id_user = ?", idUser)
	persona := gorms.PersonasGorm{}
	error := row.Scan(&persona.IDPersona, &persona.Nombre, &persona.Apellido, &persona.Domicilio, &persona.Edad, &persona.Localidad, &persona.IDPais, &persona.IDProvincia, &persona.IDTipoDoc, &persona.NroDoc)
	if error != nil {
		if error != sql.ErrNoRows {
			log.Println(error)
			panic(error)
		}
	}
	return persona
}

func (ed *PersonasDaoImpl) Save(e *gorms.PersonasGorm) int64 {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	if e.IDPersona > 0 {
		_, error := db.Exec("update personas"+
			" set nombre=?, apellido=?, domicilio=?, edad=?, localidad=?, id_pais=?, id_provincia=?, id_tipo_doc=?, nro_doc=? "+
			" where id_persona = ?", e.Nombre, e.Apellido, e.Domicilio, e.Edad, e.Localidad, e.IDPais, e.IDProvincia, e.IDTipoDoc, e.NroDoc, e.IDPersona)

		if error != nil {
			log.Println(error)
			panic(error)
		}
	} else {
		res, error := db.Exec("insert into personas"+
			" (nombre, apellido, domicilio, edad, localidad, id_pais, id_provincia, id_tipo_doc, nro_doc) "+
			" values(?,?,?,?,?,?,?,?,?)",
			e.Nombre,
			e.Apellido,
			e.Domicilio,
			e.Edad,
			e.Localidad,
			e.IDPais,
			e.IDProvincia,
			e.IDTipoDoc,
			e.NroDoc)

		if error != nil {
			log.Println(error)
			panic(error)
		} else {
			IDPersona, _ := res.LastInsertId()
			if error != nil {
				log.Println(error)
				panic(error)
			}
			e.IDPersona = IDPersona
		}
	}
	return e.IDPersona
}

func (ed *PersonasDaoImpl) Delete(id int) (bool, error) {
	db, err := application.GetDB()
	defer db.Close()
	if err != nil {
		log.Println(err.Error())
	}

	_, error := db.Exec("delete from personas where id_persona = ?", id)
	if error != nil {
		log.Println(error)
		return false, error
	}
	return true, nil
}
