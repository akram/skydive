// Code generated - DO NOT EDIT.

package ovnmodel

import (
	"github.com/skydive-project/skydive/graffiti/getter"
	"strings"
)

func (obj *DNS) GetFieldBool(key string) (bool, error) {
	return false, getter.ErrFieldNotFound
}

func (obj *DNS) GetFieldInt64(key string) (int64, error) {
	return 0, getter.ErrFieldNotFound
}

func (obj *DNS) GetFieldString(key string) (string, error) {
	switch key {
	case "UUID":
		return string(obj.UUID), nil
	}
	return "", getter.ErrFieldNotFound
}

func (obj *DNS) GetFieldKeys() []string {
	return []string{
		"UUID",
		"ExternalIDs",
		"Records",
		"ExternalIDsMeta",
		"RecordsMeta",
	}
}

func (obj *DNS) MatchBool(key string, predicate getter.BoolPredicate) bool {
	first := key
	index := strings.Index(key, ".")
	if index != -1 {
		first = key[:index]
	}

	switch first {
	case "ExternalIDsMeta":
		if index != -1 && obj.ExternalIDsMeta != nil {
			return obj.ExternalIDsMeta.MatchBool(key[index+1:], predicate)
		}
	case "RecordsMeta":
		if index != -1 && obj.RecordsMeta != nil {
			return obj.RecordsMeta.MatchBool(key[index+1:], predicate)
		}
	}
	return false
}

func (obj *DNS) MatchInt64(key string, predicate getter.Int64Predicate) bool {
	first := key
	index := strings.Index(key, ".")
	if index != -1 {
		first = key[:index]
	}

	switch first {

	case "ExternalIDsMeta":
		if index != -1 && obj.ExternalIDsMeta != nil {
			return obj.ExternalIDsMeta.MatchInt64(key[index+1:], predicate)
		}
	case "RecordsMeta":
		if index != -1 && obj.RecordsMeta != nil {
			return obj.RecordsMeta.MatchInt64(key[index+1:], predicate)
		}
	}
	return false
}

func (obj *DNS) MatchString(key string, predicate getter.StringPredicate) bool {
	if b, err := obj.GetFieldString(key); err == nil {
		return predicate(b)
	}

	first := key
	index := strings.Index(key, ".")
	if index != -1 {
		first = key[:index]
	}

	switch first {

	case "ExternalIDs":
		if obj.ExternalIDs != nil {
			if index == -1 {
				for _, v := range obj.ExternalIDs {
					if predicate(v) {
						return true
					}
				}
			} else if v, found := obj.ExternalIDs[key[index+1:]]; found {
				return predicate(v)
			}
		}
	case "Records":
		if obj.Records != nil {
			if index == -1 {
				for _, v := range obj.Records {
					if predicate(v) {
						return true
					}
				}
			} else if v, found := obj.Records[key[index+1:]]; found {
				return predicate(v)
			}
		}
	case "ExternalIDsMeta":
		if index != -1 && obj.ExternalIDsMeta != nil {
			return obj.ExternalIDsMeta.MatchString(key[index+1:], predicate)
		}
	case "RecordsMeta":
		if index != -1 && obj.RecordsMeta != nil {
			return obj.RecordsMeta.MatchString(key[index+1:], predicate)
		}
	}
	return false
}

func (obj *DNS) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}

	first := key
	index := strings.Index(key, ".")
	if index != -1 {
		first = key[:index]
	}

	switch first {
	case "ExternalIDsMeta":
		if obj.ExternalIDsMeta != nil {
			if index != -1 {
				return obj.ExternalIDsMeta.GetField(key[index+1:])
			} else {
				return obj.ExternalIDsMeta, nil
			}
		}
	case "RecordsMeta":
		if obj.RecordsMeta != nil {
			if index != -1 {
				return obj.RecordsMeta.GetField(key[index+1:])
			} else {
				return obj.RecordsMeta, nil
			}
		}

	}
	return nil, getter.ErrFieldNotFound
}

func init() {
	strings.Index("", ".")
}
