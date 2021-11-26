package generator

type TD_UserPresentResponseData struct {
	NotReceived []TD_UserPresent `json:"nr"`
	Received    []TD_UserPresent `json:"rc"`
}
