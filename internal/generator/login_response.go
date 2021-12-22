package generator

type LoginResponse struct {
	NotParamater      *int                   `json:"np,omitempty"`
	ServerTime        *string                `json:"tm,omitempty"`
	ReturnCode        *int                   `json:"rt,omitempty"`
	ResponseParameter LoginResponseParameter `json:"rp"`
	TransData         TransData              `json:"td"`
}
