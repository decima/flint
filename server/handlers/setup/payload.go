package setup

type SetupPayload struct {
	Username string `json:"username" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6"`
}

type Response struct {
	Username string `json:"username"`
}
