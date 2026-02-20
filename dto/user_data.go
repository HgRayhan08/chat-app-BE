package dto

// type PhoneNumberData struct {
// 	PhoneNumber string `json:"phone_number"`
// 	UserId      string `json:"user_id"`
// }

type SearchPhoneNumberRequest struct {
	PhoneNumber []string `json:"phone_number"`
}
