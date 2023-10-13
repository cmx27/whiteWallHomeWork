package models

type Comment struct {
	CommentID uint   `json:"comment_id" gorm:"primary_key;AUTO_INCREMENT"`
	ArticalID uint   `json:"artical_id"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
}
