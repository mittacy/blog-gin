package models

type Category struct {
	ID           int
	Title        string `gorm:"unique;not ull;size:50"`
	ArticleCount int    `gorm:"default:0"`
	Articles     []Article
}

func CreateCate(cate *Category) (string, error) {
	result := db.Create(&cate)
	defer result.Close()
	if result.Error != nil {
		return CATE_EXIST, result.Error
	}
	return "", nil
	// stmt, err := sqlDb.Prepare("INSERT INTO category(title) values (?)")
	// if err != nil {
	// 	return SQL_ERROR, err
	// }
	// defer stmt.Close()

	// result, err := stmt.Exec(cate.Title)
	// if err != nil {
	// 	return SQL_ERROR, err
	// }
	// id, _ := result.LastInsertId()
	// cate.ID = int(id)
	// return "", nil
}
