package ledg_types

import (
	"database/sql/driver"
	"github.com/ipfs/go-cid"
)

type Cid struct{ cid.Cid }
func CidFromString(s string)(Cid,error){
	r,err:=cid.Parse(s)
	return Cid{r},err
}

func (c Cid) Value() (driver.Value,error) {
	return c.String(),nil
}
func (c *Cid) Scan(value interface{}) (error) {
	v:=value.(string)
	parsed,err:=cid.Parse(v)
	if (err!=nil){
		c=&Cid{parsed}
		return nil
	}
	return err
}
