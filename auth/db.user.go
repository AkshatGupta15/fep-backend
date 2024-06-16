package auth

import "github.com/gin-gonic/gin"

func firstOrCreateUser(ctx *gin.Context, user *User) (uint, error) {
	tx := db.WithContext(ctx).Create(user)
	if tx.Error != nil {
		tx = db.WithContext(ctx).Where("user_id = ?", user.UserID).Updates(user)
	}
	return user.ID, tx.Error
}

func fetchUser(ctx *gin.Context, user *User, userID string) error {
	tx := db.WithContext(ctx).Where("user_id = ?", userID).First(&user)
	return tx.Error
}