package models

import (
	"encoding/json"
	"bytes"
	"net/http"
	"URLServer/logs"
	"fmt"
)

type EVENTINFO struct{
	EVENT     string `json:"event"`
	PARAMS    RESULTJSON `json:"params"`
}
type RESULTJSON struct{
	ID       string `json:"id"`
}
var Eventip="http://127.0.0.1:10678/"

func (ei *EVENTINFO) Init(){
	ei.EVENT=""
	ei.PARAMS=RESULTJSON{}
}

func (ei *EVENTINFO) InitAndPost(event string ,params string){
	ei.EVENT=event
	ei.PARAMS.ID=params
	ei.SendEventserver()
}
func (ei *EVENTINFO) Check()bool{
	if ei.EVENT=="" {
		return false
	}
	return true
}
func (ei *EVENTINFO) SendEventserver(){
	if ei.Check(){
		jsonbyte,err :=json.Marshal(ei)
		if err !=nil {
			logger.Log.Error(err.Error())
		}
		fmt.Println(Eventip)
		http.Post(Eventip,"application/json;charset=utf-8", bytes.NewBuffer(jsonbyte))
		ei.Init()
	}
}
