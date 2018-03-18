package types

import "time"

type jsonBase struct {
	Valid bool
}

type JsonNullString struct {
	jsonBase
	String string
}

type JsonNullInt struct {
	jsonBase
	Int int64
}

type JsonNullTime struct {
	jsonBase
	Time time.Time
}

type JsonNullFloat struct {
	jsonBase
	Float float64
}
