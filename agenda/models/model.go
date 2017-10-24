package model

import (
	"container/list"
	"encoding/json"
)

type Mlist struct {
	DataList *list.List
}

//serilize Mlist to json
func (self Mlist) MarshalJSON() ([]byte, error) {
	dl := self.DataList
	length := dl.Len()
	s := make([]interface{}, length)
	n := 0
	for e := dl.Front(); e != nil; e = e.Next() {
		s[n] = e.Value
		n++
	}
	return json.Marshal(s)
}

//decode array to Mlist
func (self *Mlist) UnmarshalJSON(data []byte) error {
	var arr []interface{}
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return err
	}
	self.DataList = list.New()
	for _, e := range arr {
		self.DataList.PushBack(e)
	}
	return err
}

//add model to datalist
func Add(model interface{}, datalist *list.List) {
	datalist.PushBack(model)
}

//del model from storage
func Del(filter func(interface{}) bool, datalist *list.List) int {
	var n *list.Element
	count := 0
	for i := datalist.Front(); i != nil; i = n {
		n = i.Next()
		if filter(i.Value) {
			datalist.Remove(i)
			count++
		}
	}
	return count
}

//query model from storage
func Query(filter func(interface{}) bool, datalist *list.List) *list.List {
	l := list.New()
	for i := datalist.Front(); i != nil; i = i.Next() {
		if filter(i.Value) {
			l.PushBack(i.Value)
		}
	}
	return l
}
