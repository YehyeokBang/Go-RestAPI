package types

/*
auth
*/
type ResponseToken struct {
	Token string `json:"token"`
}

/*
user
*/
type ResponseUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

/*
post
*/
type ResponsePost struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserID   uint   `json:"user_id"`
	Nickname string `json:"nickname"`
}
