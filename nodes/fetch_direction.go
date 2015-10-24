// Auto-generated - DO NOT EDIT

package pg_query

type FetchDirection uint

const (
	/* for these, howMany is how many rows to fetch; FETCH_ALL means ALL */
	FETCH_FORWARD = iota
	FETCH_BACKWARD
	/* for these, howMany indicates a position; only one row is fetched */
	FETCH_ABSOLUTE
	FETCH_RELATIVE
)
