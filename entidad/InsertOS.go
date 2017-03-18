
package entidad

import (
    "time"
    )

type InsertOS struct {
    CHRCODIGOOPERACIONSERVICIO string `db:"CHRCODIGOOPERACIONSERVICIO" json:"chrcodigooperacionservicio"`
    VCHCODIGOOPERACION string `db:"VCHCODIGOOPERACION" json:"vchcodigooperacion"`
	VCHDESCRIPCION   string `db:"VCHDESCRIPCION" json:"vchdescripcion"`
    NUMHORASHOMBRE   float64 `db:"NUMHORASHOMBRE" json:"numhorashombre"`
    NUMPRECIOSUGERIDO float64 `db:"NUMPRECIOSUGERIDO" json:"numpreciosugerido"`
    NUMPRECIOOFICIAL  float64 `db:"NUMPRECIOOFICIAL" json:"numpreciooficial"`
    NUMDESCUENTO      float64 `db:"NUMDESCUENTO" json:"numdescuento"`
    CHRESTADO string `db:"CHRESTADO" json:"chrestado"`
    DTECREACION time.Time `db:"DTECREACION" json:"dtecreacion"`
    NUMTIPOSERVICIO float64 `db:"NUMTIPOSERVICIO" json:"numtiposervicio"`   
    DTEMODIFICACION      time.Time   `db:"DTEMODIFICACION" json:"dtemodificacion"`              
    NUMPRECIOESTIMADO    float64    `db:"NUMPRECIOESTIMADO" json:"numprecioestimado"`            
    NUMNROPIEZA          float64     `db:"NUMNROPIEZA" json:"numnropieza"`      
    VCHNROTRABAJO         string    `db:"VCHNROTRABAJO" json:"vchnrotrabajo"`  
    NUMCODIGO             float64   `db:"NUMCODIGO" json:"numcodigo"`                 
    NUMMANOOBRA           float64   `db:"NUMMANOOBRA" json:"nummanoobra"`               
    NUMREPUESTO           float64   `db:"NUMREPUESTO" json:"numrepuesto"`                
    NUMTOTAL       float64        `db:"NUMTOTAL" json:"numtotal"`
}