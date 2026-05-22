package auth

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type SignupResult struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginResult struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
