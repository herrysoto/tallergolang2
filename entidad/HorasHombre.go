
package entidad

import (
    "time"
    )

type HorasHombre struct {
    NUMHORASHOMBRE float64 `db:"NUMHORASHOMBRE" json:"numhorashombre"`
    DTEFECHA time.Time `db:"DTEFECHA" json:"dtefecha"`
    NUMCODIGO float64 `db:"NUMCODIGO" json:"numcodigo"`
}