package dtos

type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	Code int `json:"code"`
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type ErrResponseDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RegisterRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponseDTO struct {
	Code int `json:"code"`
}

type GetTendencyDTO struct {
	PatientName   string `json:"patient_name"`
	IndicatorName int    `json:"indicator_name"`
}

type PostRecommendationRequestDTO struct {
	PatientName   string    `json:"patient_name"`
	IndicatorName int       `json:"indicator_name"`
	Indicators    []float64 `json:"indicators"`
}

type PostRecommendationResponseDTO struct {
	Code   int  `json:"code"`
	Answer bool `json:"answer"`
}
