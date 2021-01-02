package repositories

import (
	"fmt"
	"github.com/Linus-Boehm/go-serverless-suite/common"
	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/infra/persistence"
	"github.com/Linus-Boehm/go-serverless-suite/itf"
)

type ContactListEntity struct {
	BaseEntity
	Name string
	Slug string
}

type crmRepo struct {
	table itf.BaseTableProvider
}

func NewCRMRepository(table itf.BaseTableProvider) itf.CRMProvider {
	return &crmRepo{table: table}
}

func (c crmRepo) GetSubscriptionsOfEmail(email entity.ID) ([]entity.CRMEmailListSubscription, error) {
	pk := common.NewJoinedStringDBKey(entity.UserEntityName.String(), email.String())
	var rows []SubscriptionEntity
	err := c.table.ReadAllWithPK(pk, nil, entity.CRMSubscriptionEntityName, &rows)
	if err != nil {
		return nil, err
	}

	return mapSubscriptionsToEntity(rows)
}

func (c crmRepo) PutSubscription(subscription entity.CRMEmailListSubscription) error {
	row := NewSubscriptionEntity(subscription)
	return c.table.PutItem(row)
}

func (c crmRepo) PutSubscriptions(subs []entity.CRMEmailListSubscription) error {
	var rows []itf.DBKeyer
	for _, sub := range subs {
		rows = append(rows, NewSubscriptionEntity(sub))
	}
	return c.table.BatchWriteItems(rows)
}

func (c crmRepo) GetSubscriptionsOfList(listID entity.ID) ([]entity.CRMEmailListSubscription, error) {
	pk := common.NewJoinedStringDBKey(entity.CRMEmailListEntityName.String(), listID.String())
	var rows []SubscriptionEntity
	err := c.table.ReadAllWithPK(pk, &persistence.DynamoReverseIndex, entity.CRMSubscriptionEntityName, &rows)
	if err != nil {
		return nil, err
	}

	return mapSubscriptionsToEntity(rows)
}

func mapSubscriptionsToEntity(rows []SubscriptionEntity) ([]entity.CRMEmailListSubscription,error) {
	var subs []entity.CRMEmailListSubscription
	for _, row := range rows {
		sub, err := row.GetSubscription()
		if err != nil {
			return nil, err
		}
		subs = append(subs, sub)
	}
	return subs, nil
}

type SubscriptionEntity struct {
	BaseEntity
	SubscriptionStatus string `dynamo:"subscription_status,omitempty"`
	SubscriptionID string `dynamo:"subscription_id,omitempty"`
}

func NewSubscriptionEntity(s entity.CRMEmailListSubscription) itf.DBKeyer {
	return &SubscriptionEntity{
		BaseEntity: BaseEntity{
			PK:     common.JoinStringerDBKey(entity.UserEntityName, s.EMail),
			SK:     common.JoinStringerDBKey(entity.CRMEmailListEntityName, s.ListID),
			Entity: entity.CRMSubscriptionEntityName.String(),
			Slug:   fmt.Sprintf("crm-sub-%s-%s", s.EMail.String(), s.ListID.String()),
			Timestamps: s.Timestamps,
		},
		SubscriptionStatus: s.Status.String(),
		SubscriptionID: s.SubscriptionID.String(),
	}
}

func (e *SubscriptionEntity) GetSubscription() (entity.CRMEmailListSubscription, error) {
	email, err := entity.IDFromDBString(e.PK)
	if err != nil {
		return entity.CRMEmailListSubscription{}, err
	}
	listID, err := entity.IDFromDBString(e.SK)
	if err != nil {
		return entity.CRMEmailListSubscription{}, err
	}
	s := entity.CRMEmailListSubscription{
		ListID:     listID,
		EMail:      email,
		Status:     entity.SubscriptionStatus(e.SubscriptionStatus),
		Timestamps: e.Timestamps,
		SubscriptionID: entity.IDFromStringOrNil(e.SubscriptionID),
	}

	return s, nil
}