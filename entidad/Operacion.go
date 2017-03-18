package entidad


//taller estructura del registro operacion(combo segundo)
type Operacion struct {
	VCHDESCRIPCION   string `db:"VCHDESCRIPCION" json:"vchdescripcion"`
	CHRCODIGOSERVICIO int `db:"CHRCODIGOSERVICIO" json:"chrcodigoservicio"`

}