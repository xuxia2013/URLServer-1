package server

import (
	"database/sql"
	_ "github.com/lib/pq"
)
type Postgresql struct{
	DB      *sql.DB
}
func (p *Postgresql)Create()error{
	var err error
	p.DB,err=sql.Open("postgres",Settings.DB["url"][0])
	return err
}
func (p *Postgresql)Close()error{
	return p.DB.Close()
}