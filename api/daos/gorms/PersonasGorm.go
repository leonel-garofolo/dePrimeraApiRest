package gorms

import "database/sql"

type PersonasGorm struct {
	ApellidoNombre string         `gorm:"column:apellido_nombre"`
	Domicilio      sql.NullString `gorm:"column:domicilio"`
	Edad           sql.NullInt64  `gorm:"column:edad"`
	IDLiga         int64            `gorm:"column:id_liga"`
	IDLocalidad    sql.NullInt64  `gorm:"column:id_localidad"`
	IDPais         sql.NullInt64  `gorm:"column:id_pais"`
	IDPersona      int64            `gorm:"column:id_persona;primary_key"`
	IDProvincia    sql.NullInt64  `gorm:"column:id_provincia"`
	IDTipoDoc      int            `gorm:"column:id_tipo_doc"`
	NroDoc         int            `gorm:"column:nro_doc"`
}

// TableName sets the insert table name for this struct type
func (p *PersonasGorm) TableName() string {
	return "personas"
}