package entity

type Tenant struct {
	ID       ID
	ParentID ID
	Name     string
	Slug     string
	Timestamps
}

func NewTenant(name, slug string, parent ID) Tenant {
	id := NewEntityIDV4()
	if parent.IsNil() {
		parent = id
	}
	t := Tenant{
		ID:       id,
		ParentID: parent,
		Name:     name,
		Slug:     slug,
	}

	t.CreatedNow()
	return t
}
