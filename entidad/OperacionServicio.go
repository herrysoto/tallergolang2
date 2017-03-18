package entidad



type OperacionServicio struct {
    NUMNROPIEZA float64 `db:"NUMNROPIEZA" json:"numnropieza"`
    CODIGO string `json:"codigo"`
    VCHNROTRABAJO string `db:"VCHNROTRABAJO" json:"vchnrotrabajo"`
    VCHCODIGOOPERACION string `db:"VCHCODIGOOPERACION" json:"vchcodigooperacion"`
    CHRCODIGOOPERACIONSERVICIO string `db:"CHRCODIGOOPERACIONSERVICIO" json:"chrcodigooperacionservicio"`
	VCHDESCRIPCION   string `db:"VCHDESCRIPCION" json:"vchdescripcion"`
    NUMHORASHOMBRE   float64 `db:"NUMHORASHOMBRE" json:"numhorashombre"`
    NUMPRECIOSUGERIDO float64 `db:"NUMPRECIOSUGERIDO" json:"numpreciosugerido"`
    NUMDESCUENTO      float64 `db:"NUMDESCUENTO" json:"numdescuento"`
    NUMPRECIOESTIMADO float64 `db:"NUMPRECIOESTIMADO" json:"numprecioestimado"`
    NUMPRECIOOFICIAL  float64 `db:"NUMPRECIOOFICIAL" json:"numpreciooficial"`
    NUMMANOOBRA float64 `json:"nummanoobra"`
    NUMREPUESTO float64 `json:"numrepuesto"`
    NUMTOTAL float64 `json:"numtotal"`
}