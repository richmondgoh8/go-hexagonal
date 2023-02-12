package domain

type Link struct {
	ID string `json:"id" db:"id"`
	// Custom Validation required for Url
	Url  string `json:"url" db:"url" validate:"urlvalid"`
	Name string `json:"name" db:"name" validate:"required"`
}

type SimpleResp struct {
	Message    interface{} `json:"message"`
	StatusCode int         `json:"statusCode"`
}

type ApiError struct {
	Param   string
	Message string
}
