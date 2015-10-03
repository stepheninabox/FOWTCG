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

//FindByTitle is to locate rows by the name
func FindByTitle(title string) ([]*Card, error) {
	rows, err := db.Queryx("SELECT * FROM card WHERE title LIKE ?", "%"+title+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cards []*Card
	for rows.Next() {
		c := &Card{}
		if err := rows.StructScan(c); err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	return cards, rows.Err()

}
