package domain

// swagger:model
type Link struct {
	ID string `json:"id" db:"id"`
	// Custom Validation required for Url
	Url  string `json:"url" db:"url" validate:"urlvalid"`
	Name string `json:"name" db:"name" validate:"required"`
}

// swagger:model
type LinkReq struct {
	Url  string `json:"url" db:"url" validate:"urlvalid"`
	Name string `json:"name" db:"name" validate:"required"`
}

type Token struct {
	Token string `json:"token"`
}

// swagger:model
type SimpleResp struct {
	Message    interface{} `json:"message"`
	StatusCode int         `json:"statusCode"`
}

// swagger:model
type TokenResp struct {
	// Example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc1MDI4ODgsImlhdCI6MTY3NzQxNjQ4OCwiaWQiOjE1LCJyb2xlIjoiYWRtaW4ifQ.QIa7EW5jts6QhyjxN7Hxv-NbaaTIE5fbB-TrfZkiwBE
	Token string `json:"token"`
	// Example: 200
	StatusCode int `json:"statusCode"`
}

type ApiError struct {
	Param   string
	Message string
}
