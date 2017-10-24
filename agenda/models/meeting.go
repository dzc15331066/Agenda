package model

import (
	"container/list"
)

type Meeting struct {
	title         string
	start         string
	end           string
	sponsor       string
	participators *list.List
}
