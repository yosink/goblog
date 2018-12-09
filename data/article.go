package data

type Article struct {
	ID              int `json:"id"`
	CategoryID      int `json:"category_id"`
	UserID          int `json:"author"`
	LastUserID      int `json:"-"`
	Slug            string
	Title           string `json:"title"`
	SubTitle        string `gorm:"column:subtitle"`
	Content         string `json:"content"`
	PageImage       string `json:"image"`
	MetaDescription string `json:"-"`
	IsOriginal      bool
	IsDraft         bool `gorm:"-" json:"-"`
	ViewCount       int
	PublishedAt     LocalTime
	CreatedAt       LocalTime
}

func GetArticleByID(id int) (art Article, err error) {
	defer DB.Close()
	err = DB.First(&art, id).Error
	return
}
func GetArticleList() (arts []Article, err error) {
	defer DB.Close()
	where := Article{IsDraft: false}
	limit := 10
	order := "created_at desc"
	DB.Order(order).Limit(limit).Find(&arts, where)
	return arts, nil
}
