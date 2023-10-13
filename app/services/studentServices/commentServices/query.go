package commentServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func GetCommmentByUserIDAndCommentID(user_id uint, comment_id uint) (*models.Comment, error) {
	var comment models.Comment
	result := database.DB.Where(&models.Comment{
		UserID:    user_id,
		CommentID: comment_id,
	}).First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comment, nil
}

func GetCommentByUserID(user_id uint) ([]models.Comment, error) {
	var comment_list []models.Comment
	result := database.DB.Where(&models.Comment{
		UserID: user_id,
	}).Find(&comment_list)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment_list, nil
}

func GetCommentByCommentID(comment_id uint) (models.Comment, error) {
	var comment models.Comment
	result := database.DB.Where(&models.Comment{
		CommentID: comment_id,
	}).Find(&comment)
	if result.Error != nil {
		return comment, result.Error
	}
	return comment, nil
}

func GetCommentByArticalID(artical_id uint) ([]models.Comment, error) {
	var comment_list []models.Comment
	result := database.DB.Where(&models.Comment{
		ArticalID: artical_id,
	}).Find(&comment_list)
	if result.Error != nil {
		return comment_list, result.Error
	}
	return comment_list, nil
}
