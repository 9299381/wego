package jwt

type Claims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	Iat  int64  `json:"iat"`
	Exp  int64  `json:"exp"`
}
