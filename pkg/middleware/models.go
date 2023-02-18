package custommiddleware

type JWTPayload struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}
