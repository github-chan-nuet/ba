package model

type Quote struct {
	Quote string `db:"QUOTE" fieldtag:"pk" json:"quote"`
	Song  string `db:"SONG" json:"song"`
	Album string `db:"ALBUM" json:"album"`
}
