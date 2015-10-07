package game

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

type Cards struct {
	cards []*Card
}

func (c *Cards) Len() int { return len(c.cards) }

func (c *Cards) Get(pos int) *Card { return c.cards[pos] }

//FindByTitle is to locate rows by the name
func FindByTitle(title string) (*Cards, error) {
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
	return &Cards{cards}, rows.Err()
}
