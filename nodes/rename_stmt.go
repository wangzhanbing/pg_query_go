// Auto-generated - DO NOT EDIT

package pg_query

type RenameStmt struct {
	RenameType   ObjectType `json:"renameType"`   /* OBJECT_TABLE, OBJECT_COLUMN, etc */
	RelationType ObjectType `json:"relationType"` /* if column name, associated relation type */
	Relation     *RangeVar  `json:"relation"`     /* in case it's a table */
	Object       []Node     `json:"object"`       /* in case it's some other object */
	Objarg       []Node     `json:"objarg"`       /* argument types, if applicable */
	Subname      *string    `json:"subname"`      /* name of contained object (column, rule,
	 * trigger, etc) */
	Newname   *string      `json:"newname"`    /* the new name */
	Behavior  DropBehavior `json:"behavior"`   /* RESTRICT or CASCADE behavior */
	MissingOk bool         `json:"missing_ok"` /* skip error if missing? */
}
