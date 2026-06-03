package models

type Status string

const (
	StatusNew        Status = "NEW"
	StatusInProgress        = "IN_PROGRESS"
	StatusDone              = "DONE"
)
