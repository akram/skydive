// Code generated - DO NOT EDIT.

package ovnmodel

import (
	"github.com/skydive-project/skydive/graffiti/getter"
	"strings"
)

func (obj *QoS) GetFieldBool(key string) (bool, error) {
	return false, getter.ErrFieldNotFound
}

func (obj *QoS) GetFieldInt64(key string) (int64, error) {
	switch key {
	case "Priority":
		return int64(obj.Priority), nil
	}
	return 0, getter.ErrFieldNotFound
}

func (obj *QoS) GetFieldString(key string) (string, error) {
	switch key {
	case "UUID":
		return string(obj.UUID), nil
	case "Direction":
		return string(obj.Direction), nil
	case "Match":
		return string(obj.Match), nil
	}
	return "", getter.ErrFieldNotFound
}

func (obj *QoS) GetFieldKeys() []string {
	return []string{
		"UUID",
		"Action",
		"Bandwidth",
		"Direction",
		"ExternalIDs",
		"Match",
		"Priority",
		"ActionMeta",
		"BandwidthMeta",
		"ExternalIDsMeta",
	}
}

func (obj *QoS) MatchBool(key string, predicate getter.BoolPredicate) bool {
	first := key
	index := strings.Index(key, ".")
	if index != -1 {
		first = key[:index]
	}

	switch first {
	case "ActionMeta":
		if index != -1 && obj.ActionMeta != nil {
			return obj.ActionMeta.MatchBool(key[index+1:], predicate)
		}
	case "BandwidthMeta":
		if index != -1 && obj.BandwidthMeta != nil {
			return obj.BandwidthMeta.MatchBool(key[index+1:], predicate)
		}
	case "ExternalIDsMeta":
		if index != -1 && obj.ExternalIDsMeta != nil {
			return obj.ExternalIDsMeta.MatchBool(key[index+1:], predicate)
		}
	}
	return false
}

func (obj *QoS) MatchInt64(key string, predicate getter.Int64Predicate) bool {
	if b, err := obj.GetFieldInt64(key); err == nil {
		return predicate(b)
	}

	first := key
	index := strings.Index(key, ".")
	if index != -1 {
		first = key[:index]
	}

	switch first {

	case "ActionMeta":
		if index != -1 && obj.ActionMeta != nil {
			return obj.ActionMeta.MatchInt64(key[index+1:], predicate)
		}
	case "BandwidthMeta":
		if index != -1 && obj.BandwidthMeta != nil {
			return obj.BandwidthMeta.MatchInt64(key[index+1:], predicate)
		}
	case "ExternalIDsMeta":
		if index != -1 && obj.ExternalIDsMeta != nil {
			return obj.ExternalIDsMeta.MatchInt64(key[index+1:], predicate)
		}
	}
	return false
}

func (obj *QoS) MatchString(key string, predicate getter.StringPredicate) bool {
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
	case "ActionMeta":
		if index != -1 && obj.ActionMeta != nil {
			return obj.ActionMeta.MatchString(key[index+1:], predicate)
		}
	case "BandwidthMeta":
		if index != -1 && obj.BandwidthMeta != nil {
			return obj.BandwidthMeta.MatchString(key[index+1:], predicate)
		}
	case "ExternalIDsMeta":
		if index != -1 && obj.ExternalIDsMeta != nil {
			return obj.ExternalIDsMeta.MatchString(key[index+1:], predicate)
		}
	}
	return false
}

func (obj *QoS) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}

	if i, err := obj.GetFieldInt64(key); err == nil {
		return i, nil
	}

	first := key
	index := strings.Index(key, ".")
	if index != -1 {
		first = key[:index]
	}

	switch first {
	case "ActionMeta":
		if obj.ActionMeta != nil {
			if index != -1 {
				return obj.ActionMeta.GetField(key[index+1:])
			} else {
				return obj.ActionMeta, nil
			}
		}
	case "BandwidthMeta":
		if obj.BandwidthMeta != nil {
			if index != -1 {
				return obj.BandwidthMeta.GetField(key[index+1:])
			} else {
				return obj.BandwidthMeta, nil
			}
		}
	case "ExternalIDsMeta":
		if obj.ExternalIDsMeta != nil {
			if index != -1 {
				return obj.ExternalIDsMeta.GetField(key[index+1:])
			} else {
				return obj.ExternalIDsMeta, nil
			}
		}

	}
	return nil, getter.ErrFieldNotFound
}

func init() {
	strings.Index("", ".")
}
