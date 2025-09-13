package domain

import "time"

var NowUTC = func() time.Time { return time.Now().UTC().Round(0) }

type Meta struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	IsVoid    *time.Time `json:"isVoid"`
}

func CreateMeta() Meta {
	t := NowUTC()
	return Meta{
		CreatedAt: t,
		UpdatedAt: nil,
		IsVoid:    nil,
	}
}

func (m *Meta) Update() {
	t := NowUTC()
	m.UpdatedAt = &t
}

func (m *Meta) Delete() {
	t := NowUTC()
	m.IsVoid = &t
}
