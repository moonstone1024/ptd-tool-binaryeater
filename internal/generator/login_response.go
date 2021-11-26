package generator

type LoginResponse struct {
	NotParamater      int                    `json:"np"`
	ServerTime        string                 `json:"tm"`
	ReturnCode        int                    `json:"rt"`
	ResponseParameter LoginResponseParameter `json:"rp"`
	TransData         TransData              `json:"td"`
}
