package domain

type User struct {
	Meta
	Id         string  `json:"id"`
	FirstName  string  `json:"firstName"`
	MiddleName *string `json:"middleName"`
	LastName   string  `json:"lastName"`
}

func NewUser() *User {
	var meta = CreateMeta()
	return &User{
		Meta:      meta,
		Id:        "sl1234",
		FirstName: "Simon",
		LastName:  "Lidwell",
	}
}

type UserOption func(*User)

func WithFirstName(v string) UserOption  { return func(u *User) { u.FirstName = v } }
func WithLastName(v string) UserOption   { return func(u *User) { u.LastName = v } }
func WithMiddleName(v string) UserOption { return func(u *User) { u.MiddleName = &v } }
func WithoutMiddleName() UserOption      { return func(u *User) { u.MiddleName = nil } }

func (u *User) Update(opts ...UserOption) {
	for _, opt := range opts {
		opt(u)
	}
	u.Meta.Touch()
}
