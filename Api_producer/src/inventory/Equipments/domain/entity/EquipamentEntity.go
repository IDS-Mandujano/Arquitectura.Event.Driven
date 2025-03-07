package entities

type equipament struct {
	id int32
	cname string
	category string
	ccondition string
}

func NewEquipament(cname string, category string, ccondition string) *equipament {
	return &equipament{cname: cname, category: category, ccondition: ccondition}
}