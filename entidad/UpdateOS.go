package entidad

// import (
//     "time"
//     )

// type UpdateOS struct {
//     VCHDESCRIPCION string `db:"VCHDESCRIPCION" json:"vchdescripcion"`
//     NUMPRECIOSUGERIDO float64 `db:"NUMPRECIOSUGERIDO" json:"numpreciosugerido"`
//     NUMHORASHOMBRE float64 `db:"NUMHORASHOMBRE" json:"numhorashombre"`
//     NUMDESCUENTO float64 `db:"NUMDESCUENTO" json:"numdescuento"`
//     DTEMODIFICACION time.Time `db:"DTEMODIFICACION" json:"dtemodificacion"`
//     NUMPRECIOESTIMADO float64 `db:"NUMPRECIOESTIMADO" json:"numprecioestimado"`
//     CHRCODIGOOPERACIONSERVICIO string `db:"CHRCODIGOOPERACIONSERVICIO" json:"chrcodigooperacionservicio"`
//     VCHCODIGOOPERACION string `db:"VCHCODIGOOPERACION" json:"vchcodigooperacion"`
// }

type UpdateOS struct {
	VCHNROTRABAJO string `db:"VCHNROTRABAJO" json:"vchnrotrabajo"`
	VCHDESCRIPCION    string  `db:"VCHDESCRIPCION" json:"vchdescripcion"`
	NUMPRECIOSUGERIDO float64 `db:"NUMPRECIOSUGERIDO" json:"numpreciosugerido"`
	NUMHORASHOMBRE    float64 `db:"NUMHORASHOMBRE" json:"numhorashombre"`
	NUMDESCUENTO      float64 `db:"NUMDESCUENTO" json:"numdescuento"`
	// DTEMODIFICACION time.Time `db:"DTEMODIFICACION" json:"dtemodificacion"`
	NUMPRECIOESTIMADO          float64 `db:"NUMPRECIOESTIMADO" json:"numprecioestimado"`
	CHRCODIGOOPERACIONSERVICIO string  `db:"CHRCODIGOOPERACIONSERVICIO" json:"chrcodigooperacionservicio"`
	VCHCODIGOOPERACION         string  `db:"VCHCODIGOOPERACION" json:"vchcodigooperacion"`
	NUMTOTAL       float64        `db:"NUMTOTAL" json:"numtotal"`
}