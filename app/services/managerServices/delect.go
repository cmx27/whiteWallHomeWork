package managerServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func DeleteArticalByArticalID(artical_id uint) error {
	result4 := database.DB.Where(&models.Comment{
		ArticalID: artical_id,
	}).Delete(&models.Comment{})
	if result4.Error != nil {
		return result4.Error
	}
	result := database.DB.Where(&models.Artical{
		ArticalID: artical_id,
	}).Delete(&models.Artical{})
	return result.Error
}

func DeleteCommentByCommentID(comment_id uint) error {
	result := database.DB.Where(&models.Comment{
		CommentID: comment_id,
	}).Delete(&models.Comment{})
	return result.Error
}

func DeleteUserByUserID(user_id uint) error {
	result := database.DB.Unscoped().Where("user_id = ?", user_id).Where("manager_state <> ?", true).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	result3 := database.DB.Where(&models.Student{
		StudentID: user_id,
	}).Delete(&models.Student{})
	if result3.Error != nil {
		return result3.Error
	}
	result4 := database.DB.Where(&models.Comment{
		UserID: user_id,
	}).Delete(&models.Comment{})
	if result4.Error != nil {
		return result4.Error
	}
	result2 := database.DB.Where(&models.Artical{
		UserID: user_id,
	}).Delete(&models.Artical{})

	return result2.Error

}
