package entity

type Mail struct {
	Name string
	Mail string
}

type ContactList struct {
	ID             ID
	Name           string
	RecipientCount int
}

type MinimalMail struct {
	FromMail     Mail
	ToMail       Mail
	Subject      *string
	HTMLTemplate HTMLTemplate
}

func (m *MinimalMail) GetPlainText() string {
	return m.HTMLTemplate.GetPlainText()
}

func (m *MinimalMail) GetHTMLTemplate() string {
	return m.HTMLTemplate.GetHTML()
}

type ContactForm struct {
	SenderMail Mail
	FromMail   Mail
	ToMail     Mail
	Subject    string
	Message    string
	Attributes map[string]string
}
