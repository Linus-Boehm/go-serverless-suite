package mailings

type MailAttributes map[string]string

type SimpleContactFormInput struct {
	SenderMail Mail
	FromMail   Mail
	ToMail     Mail
	Subject    string
	Message    string
	Attributes MailAttributes
}
