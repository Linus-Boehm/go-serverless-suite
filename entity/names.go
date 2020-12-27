package entity

const (
	UserEntityName   Name = "USER"
	TenantEntityName Name = "TENANT"
)

type Name string

func (e Name) String() string {
	return string(e)
}

type TableIndex struct {
	Name string
	PK   string
	SK   string
}
