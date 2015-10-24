// Auto-generated - DO NOT EDIT

package pg_query

type TypeName struct {
	Names       []Node `json:"names"`       /* qualified name (list of Value strings) */
	TypeOid     Oid    `json:"typeOid"`     /* type identified by OID */
	Setof       bool   `json:"setof"`       /* is a set? */
	PctType     bool   `json:"pct_type"`    /* %TYPE specified? */
	Typmods     []Node `json:"typmods"`     /* type modifier expression(s) */
	Typemod     int32  `json:"typemod"`     /* prespecified type modifier */
	ArrayBounds []Node `json:"arrayBounds"` /* array bounds */
	Location    int    `json:"location"`    /* token location, or -1 if unknown */
}
