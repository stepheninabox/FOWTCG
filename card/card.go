package card

type Card struct {
	Id      int `db:"_id"`
	Title   string
	Deck    string
	Image   string
	Cat     string
	Subcat  string
	Cost    string
	TCost   string
	Attack  string
	Defense string
	Attr    string
	Rarity  string
	Descr   string
}
