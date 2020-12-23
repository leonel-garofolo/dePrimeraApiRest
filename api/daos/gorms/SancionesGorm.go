package gorms

import "database/sql"

type SancionesGorm struct {
	Descripcion   sql.NullString `gorm:"column:descripcion"`
	IDLigas       int64            `gorm:"column:id_ligas"`
	IDSanciones   int64            `gorm:"column:id_sanciones;primary_key"`
	Observaciones sql.NullString `gorm:"column:observaciones"`
}

// TableName sets the insert table name for this struct type
func (s *SancionesGorm) TableName() string {
	return "sanciones"
}
