package entidad



type OperacionServiciosContenidos struct {
    NUMNROPIEZA float64 `db:"NUMNROPIEZA" json:"numnropieza"`
    CODIGO string `json:"codigo"`
    VCHNROTRABAJO string `db:"VCHNROTRABAJO" json:"vchnrotrabajo"`
    NUMMANOOBRA float64 `json:"nummanoobra"`
    NUMREPUESTO float64 `json:"numrepuesto"`
    NUMTOTAL float64 `json:"numtotal"`
    VCHCODIGOOPERACION string `db:"VCHCODIGOOPERACION" json:"vchcodigooperacion"`
    CHRCODIGOOPERACIONSERVICIO string `db:"CHRCODIGOOPERACIONSERVICIO" json:"chrcodigooperacionservicio"`
    VCHDESCRIPCION   string `db:"VCHDESCRIPCION" json:"vchdescripcion"`
    NUMPRECIOSUGERIDO float64 `db:"NUMPRECIOSUGERIDO" json:"numpreciosugerido"`
    NUMHORASHOMBRE   float64 `db:"NUMHORASHOMBRE" json:"numhorashombre"`
    NUMDESCUENTO      float64 `db:"NUMDESCUENTO" json:"numdescuento"`
}