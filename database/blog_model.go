package database

// blog data structure [blogのデータ構造]
type BlogModel struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	PostID int       `json:"post_id"`
	Post   PostModel `json:"detail"`
}

// post data structure [postのデータ構造]
type PostModel struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// shops data structure [shopsのデータ構造]
type ShopsModel struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Product_id int       `json:"product_id"`
	Product    string    `json:"product"`
	Post       PostModel `json:"detail"`
}

// Shopspost data structure [Shopspostのデータ構造]
type ShopsPostModel struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
