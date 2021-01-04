package entity

const (
	UserOptedInSubscriptionStatus  SubscriptionStatus = "USER_OPTED_IN"
	UserOptedInRequestedSubscriptionStatus  SubscriptionStatus = "USER_OPTED_IN_REQUESTED"
	AdminSubscriptionStatus        SubscriptionStatus = "ADMIN_CREATED"
	UserOptedOutSubscriptionStatus SubscriptionStatus = "USER_OPTED_OUT"
)

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
	ListID ID
	EMail ID
	SubscriptionID ID
	Status SubscriptionStatus
	Timestamps
}

type SubscriptionStatus string

func (s SubscriptionStatus) String() string {
	return string(s)
}
