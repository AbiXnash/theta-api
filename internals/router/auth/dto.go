package auth

type RegisterRequest struct {
	FullName        string `form:"fullName" binding:"required,min=3"`
	Password        string `form:"password" binding:"required,min=8"`
	Email           string `form:"email" binding:"required,email"`
	PhoneNo         string `form:"phoneNo" binding:"required,len=10,numeric"`
	RegNo           string `form:"regNo" binding:"required"`
	IsSastraStudent bool   `form:"isSastraStudent"`
	Department      string `form:"department" binding:"required"`
	Year            string `form:"year" binding:"required"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=8"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
