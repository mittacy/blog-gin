package models

type Category struct {
	ID           int
	title        string `gorm:"unique;not ull;size:50"`
	ArticleCount int    `gorm:"default:0"`
	Articles     []Article
}

func CreateCate(title string) (string, error) {
	stmt, err := sqlDb.Prepare("INSERT INTO category(title) values (?)")
	if err != nil {
		return SQL_ERROR, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title)
	if err != nil {
		return SQL_ERROR, err
	}
	return "", nil
}
