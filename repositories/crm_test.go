package repositories

import (
	"strconv"
	"testing"
	"time"

	"github.com/Linus-Boehm/go-serverless-suite/entity"
	"github.com/Linus-Boehm/go-serverless-suite/infra/persistence"
	"github.com/stretchr/testify/assert"
)

func TestCrmRepo_PutSubscriptions(t *testing.T) {
	now := time.Now().Unix()

	tests := []struct {
		name      string
		inputID   entity.ID
		putSubs   func(id entity.ID) []entity.CRMEmailListSubscription
		expectErr bool
	}{
		{
			name:    "happy",
			inputID: entity.IDFromStringOrNil("email@example.org"),
			putSubs: func(id entity.ID) []entity.CRMEmailListSubscription {
				subs := []entity.CRMEmailListSubscription{}
				for i := 0; i < 5; i++ {
					subs = append(subs, entity.CRMEmailListSubscription{
						ListID:         entity.IDFromStringOrNil(strconv.Itoa(i + 1)),
						EMail:          id,
						SubscriptionID: entity.IDFromStringOrNil(strconv.Itoa(i + 1)),
						Status:         entity.UserOptedInRequestedSubscriptionStatus,
						Timestamps: entity.Timestamps{
							CreatedAt: now,
						},
					})
				}
				return subs
			},
			expectErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := persistence.NewTestProvider(SubscriptionEntity{})
			assert.NoError(t, err)
			defer b.DeleteTable()

			repo := NewCRMRepository(b)
			putSubs := test.putSubs(test.inputID)
			err = repo.PutSubscriptions(putSubs)

			assert.Equal(t, test.expectErr, err != nil)
			if !test.expectErr {
				subs, err := repo.GetSubscriptionsOfEmail(test.inputID)
				assert.NoError(t, err)
				assert.Equal(t, putSubs, subs)
			}

		})
	}
}
