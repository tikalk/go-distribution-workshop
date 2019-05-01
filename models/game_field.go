package models

import (
	"fmt"
	"time"
	"golang.org/x/sync/syncmap"

	"encoding/json"
)

type (
	GameField struct {
		Items syncmap.Map
	}
)

func NewGameField() *GameField{
	res := &GameField{Items: syncmap.Map{}}
	return res
}

func (gf *GameField) cleanup(){
	gf.Items.Range(func(key, value interface{}) bool {
		val, ok := value.(*DisplayStatus)
		if !ok {
			// this will break iteration
			return false
		}
		if time.Now().Sub(val.LastUpdated) > 60 * time.Second && val.ItemType != TypeBall{
			gf.Items.Delete(key)
		}
		return true
	})

}

func (gf *GameField) Update(item *DisplayStatus){
	key := "ball"
	if item.ItemType != TypeBall {
		key = fmt.Sprintf("%s|%s|%s", item.ItemType, item.TeamID, item.ItemID)
	}
	gf.Items.Store(key, item)
	//gf.cleanup()
}

func (gf *GameField) MarshalJSON() ([]byte, error){
	res := make(map[string]interface{})
	var items = make(map[string]interface{})
	res["items"] = items

	gf.Items.Range(func(key, value interface{}) bool {
		val, ok := value.(*DisplayStatus)
		if !ok {
			// this will break iteration
			return false
		}

		k, ok := key.(string)
		if !ok {
			// this will break iteration
			return false
		}

		items[k] = val
		return true
	})

	return json.Marshal(res)

}
