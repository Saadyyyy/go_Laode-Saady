package request

type LooginRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
