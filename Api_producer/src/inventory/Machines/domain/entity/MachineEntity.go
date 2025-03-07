package entities

type machine struct {
	id int32
	cname string
	ctype string
	cstatus string
}

func NewMachine(cname string, ctype string, cstatus string) *machine {
	return &machine{cname: cname, ctype: ctype, cstatus: cstatus}
}