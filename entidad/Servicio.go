package entidad


//taller estructura del registro servicio(combo primero)
type Servicio struct {
	VCHDESCRIPCION   string `db:"VCHDESCRIPCION" json:"vchdescripcion"`
	CHRCODIGOSERVICIO string `db:"CHRCODIGOSERVICIO" json:"chrcodigoservicio"`
	VCHCODIGOOPERACION string `db:"VCHCODIGOOPERACION" json:"vchcodigooperacion"`
}