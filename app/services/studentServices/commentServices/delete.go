package commentServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func DeleteComment(comment_id uint) error {
	result := database.DB.Where(&models.Comment{
		CommentID: comment_id,
	}).Delete(&models.Comment{})
	return result.Error
}
