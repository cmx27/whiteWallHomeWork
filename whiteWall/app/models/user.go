package models

type User struct {
	UserID   uint   `json:"user_id" gorm:"primary_key;AUTO_INCREMENT"`
	Account  string `json:"account"` //`gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram INVISIBLE"`
	Password string `json:"-"`
	Name     string `json:"name"` //`gorm:"index:,class:FULLTEXT,option:WITH PARSER ngram INVISIBLE"`
	Sex      string `json:"sex"`
	Major    string `json:"major"`
	//Token        string `default:"五花肉串串"`
	ManagerState bool `json:"manager_state"`
}
