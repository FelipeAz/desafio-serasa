package entity

// TokenDetails detalhe dos tokens
type TokenDetails struct {
	AccessToken  string
	AtExpires    int64
	RefreshToken string
	RtExpires    int64
}
