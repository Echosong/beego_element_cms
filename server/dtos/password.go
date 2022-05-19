package dtos

type PasswordDto struct {
	NewPassword string `json:"newPassword"`
	Password    string `json:"password"`
	RePassword  string `json:"rePassword"`
	UserId      int    `json:"user_id"`
}
