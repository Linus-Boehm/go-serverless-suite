package entity

import (
	"embed"
	"html/template"
)

const (
	UserOptedInSubscriptionStatus          SubscriptionStatus = "USER_OPTED_IN"
	UserOptedInRequestedSubscriptionStatus SubscriptionStatus = "USER_OPTED_IN_REQUESTED"
	AdminSubscriptionStatus                SubscriptionStatus = "ADMIN_CREATED"
	UserOptedOutSubscriptionStatus         SubscriptionStatus = "USER_OPTED_OUT"
)

// TODO move out non db related stuff

type CRMOptInMailOptions struct {
	Fullname         string  // Name of the user (optional)
	Subject          *string // Subject (optional)
	EMail            ID
	UserID           ID
	SubID            ID
	ConfirmationPath string
	Template         *TemplateManifest
	FS               *embed.FS
}

type CRMOptInMailTemplateOptions struct {
	//path of the frontend without trailing slash like https://localhost/crm/confirm
	ConfirmURL template.HTML
	Email      string
	FullName   string
}

type CRMContactList struct {
	ID           ID
	DisplayTitle string
	WorkingTitle string
	Slug         string
	Note         string
	ContactCount int
	Timestamps
}

type CRMContactLists []CRMContactList

type CRMEmailListSubscription struct {
	ListID         ID
	EMail          ID
	SubscriptionID ID
	Status         SubscriptionStatus
	Timestamps
}

type SubscriptionStatus string

func (s SubscriptionStatus) String() string {
	return string(s)
}
