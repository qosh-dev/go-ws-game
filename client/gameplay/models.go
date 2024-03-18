package gameplay

type GameStatusPayload struct {
	Id     uint   `json:"id"`
	Login  string `json:"login"`
	Health int8   `json:"health"`
}
