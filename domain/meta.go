package domain

import "time"

type Meta struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	IsVoid    time.Time
}

func (m *Meta) Create() {
	m.CreatedAt = time.Now()
}

func (m *Meta) Update() {
	m.UpdatedAt = time.Now()
}

func (m *Meta) Delete() {
	m.IsVoid = time.Now()
}
