package entity

import "time"

type Timestamper interface {
	GetTimestamps() Timestamps
}

type Timestamps struct {
	CreatedAt   int64  `dynamo:"created_at,omitempty"`
	UpdatedAt   int64  `dynamo:"updated_at,omitempty"`
	PublishedAt *int64 `dynamo:"published_at,omitempty"`
	DeletedAt   *int64 `dynamo:"deleted_at,omitempty"`
}

func now() int64 {
	return time.Now().Unix()
}

func (t *Timestamps) CreatedNow() {
	n := now()
	if t.CreatedAt == 0 {
		t.CreatedAt = n
	}
	t.UpdatedAt = n
	t.DeletedAt = nil
}

func (t *Timestamps) UpdatedNow() {
	t.UpdatedAt = now()
}

func (t *Timestamps) PublishNow() {
	n := now()
	t.PublishedAt = &n
	t.UpdatedAt = n
}

func (t *Timestamps) SoftDeleteNow() {
	n := now()
	t.UpdatedAt = n
	t.DeletedAt = &n
}
