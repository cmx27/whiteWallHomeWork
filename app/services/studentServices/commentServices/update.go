package commentServices

import (
	"whiteWall/app/models"
	"whiteWall/config/database"
)

func UpdateComment(comment string, comment_id uint) error {
	result := database.DB.Model(&models.Comment{
		CommentID: comment_id,
	}).Updates(models.Comment{
		Comment: comment})
	return result.Error
}

func UpdateImage(artical_id uint, imageAddr string) error {

	result := database.DB.Where(&models.Artical{
		ArticalID: artical_id,
	}).Updates(models.Artical{
		ImageAddr: imageAddr,
	})
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}

}
