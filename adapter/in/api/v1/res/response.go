package res


// Name string `json:"name" time_format:"2006-01-02"`
// 
type User struct {
	Id   uint  `json:"id"`
	Name string `json:"name" `
}

func NewUser(Id uint, name string) *User {
	return &User{Id: Id, Name: name}
}