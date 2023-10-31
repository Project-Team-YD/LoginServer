package response_packet

type Login struct {
	MessageCode uint   `json:"messageCode"`
	Message     string `json:"message"`
	UUID        string `json:"UUID"`
}
