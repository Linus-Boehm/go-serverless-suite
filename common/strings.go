package common

import (
	"fmt"
	"strings"
)

func StringValue(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func StringPtr(val string) *string {
	if val == "" {
		return nil
	}
	return &val
}

type String struct {
	Val string
}

func NewString(s string) String {
	return String{Val: s}
}

func (s String) String() string {
	return s.Val
}

func (s String) StringPtr() *string {
	return &s.Val
}

func NewJoinedStringDBKey(s ...string) String {
	return String{Val: JoinDBKey(s...)}
}

func JoinDBKey(s ...string) string {
	return strings.Join(s, "#")
}

func JoinStringerDBKey(stringer ...fmt.Stringer) string {
	var strs []string
	for _, s := range stringer {
		strs = append(strs, s.String())
	}
	return strings.Join(strs, "#")
}

type StringKey struct {
	PK  String
	SK  String
	Ent fmt.Stringer
}

func NewStringKey(pk string, sk string, entity fmt.Stringer) *StringKey {
	return &StringKey{
		PK:  NewJoinedStringDBKey(entity.String(), pk),
		SK:  NewJoinedStringDBKey(entity.String(), sk),
		Ent: entity,
	}
}
func (s StringKey) GetPK() fmt.Stringer {
	return s.PK
}
func (s StringKey) GetSK() fmt.Stringer {
	return s.SK
}
func (s StringKey) GetEntity() fmt.Stringer {
	return s.Ent
}
