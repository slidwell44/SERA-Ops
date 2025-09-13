package domain

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

var NowUTC = func() time.Time { return time.Now().UTC().Round(0) }

func NewHex() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b[:])
}

type Meta struct {
	Uid       string     `json:"uid"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	IsVoid    *time.Time `json:"isVoid"`
}

func CreateMeta() Meta {
	t := NowUTC()
	return Meta{
		Uid:       NewHex(),
		CreatedAt: t,
		UpdatedAt: nil,
		IsVoid:    nil,
	}
}

func (m *Meta) Touch() {
	t := NowUTC()
	m.UpdatedAt = &t
}

func (m *Meta) Delete() {
	t := NowUTC()
	m.IsVoid = &t
}
