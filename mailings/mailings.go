package mailings

type MailingsService struct {
	provider IMailProvider
}

func NewMailingsService(provider IMailProvider) *MailingsService {
	return &MailingsService{provider: provider}
}


func SendSimpleContactForm(input SimpleContactFormInput) error {

	return nil
}