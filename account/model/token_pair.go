package model

// TokenPair used for returning pair of id and refresh token
type TokenPair struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}
