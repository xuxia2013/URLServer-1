package models
import(
	
)
type URLLISTS struct{
	LIST []URLLIST  `json:"list"`
}
type URLLIST struct{
	ID  string  `json:"id"`
	URL string  `json:"url"`
}

func (url * URLLIST) Init() {
	url.ID=""
	url.URL=""
}
func (url * URLLIST) Check() bool {
	if url.ID==""{
		return false
	}else if url.URL==""{
		return false
	}
	return true
}
func (url * URLLISTS) Get() URLLIST {
	if len(url.LIST)<1 {
		return URLLIST{}
	}else{
		for i:=range url.LIST{
			return url.LIST[i]
		}
	}
	return URLLIST{}
}

func (url * URLLISTS) GetId (id string) URLLIST{
	if  len(url.LIST)<1{
		return URLLIST{}
	}else {
		for i:=range url.LIST{
			if url.LIST[i].ID==id {
				return url.LIST[i]
			}else {
				return URLLIST{}
			}
		}
	}
	return URLLIST{}
}
func (url * URLLISTS) Insert(u URLLIST) {
	url.LIST=append(url.LIST,u)
}
func (url * URLLISTS) Delete(id string) bool {
	var	newlist []URLLIST 
	if len(url.LIST)<1{
		return false
	}else{
		for i:=range url.LIST{
			if url.LIST[i].ID!=id{
				newlist=append(newlist,url.LIST[i])
			}
		}
		if len(newlist)>0{
			url.LIST=newlist
		}else{
			return false
		}
		
	}
	return true
}
func (url * URLLISTS) Clean() {
	url.LIST=[]URLLIST{}
}
