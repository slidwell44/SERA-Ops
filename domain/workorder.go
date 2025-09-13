package domain

type WorkOrder struct {
	Meta
	Id         int64
	PartNumber string
	Operations []Operation
}
