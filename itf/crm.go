package itf

type NewsWriter interface {
}

type CRMServicer interface {
	GetMailer() Mailer
	GetNewsWriter() NewsWriter
}
