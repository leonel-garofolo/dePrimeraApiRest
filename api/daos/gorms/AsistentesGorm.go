package gorms

type AsistentesGorm struct {
	IDAsistente  int64 `gorm:"column:id_asistente;primary_key"`
	IDPersona    int64 `gorm:"column:id_persona"`
	IDCampeonato int64 `gorm:"column:id_campeonato"`
}

// TableName sets the insert table name for this struct type
func (a *AsistentesGorm) TableName() string {
	return "asistentes"
}
