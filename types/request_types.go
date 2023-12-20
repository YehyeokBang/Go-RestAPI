package types

// auth
type RequestLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

/*
user
*/
type RequestCreateUser struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestUpdateUser struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

// post
type RequestCreatePost struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type RequestUpdatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
