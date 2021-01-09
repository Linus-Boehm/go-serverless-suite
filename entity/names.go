package entity

const (
	UserEntityName            Name = "USER"
	CRMSubscriptionEntityName Name = "CRM_SUB"
	CRMEmailListEntityName    Name = "CRM_EMAIL_LIST"
	TenantEntityName          Name = "TENANT"
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
