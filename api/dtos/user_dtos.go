package dtos

type UserCreateIn struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhotoURL string `json:"photoUrl"`
	HouseID  string `json:"houseId"`
}

type UserChangeHouseIn struct {
	ID      string `json:"id"`
	HouseID string `json:"houseId"`
}
