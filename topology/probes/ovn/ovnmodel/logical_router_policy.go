//go:generate go run github.com/skydive-project/skydive/graffiti/gendecoder -package github.com/skydive-project/skydive/topology/probes/ovn/ovnmodel
//go:generate go run github.com/mailru/easyjson/easyjson $GOFILE

// Code generated by ovnmetagen
// DO NOT EDIT.

package ovnmodel

import (
	"encoding/json"
	"fmt"

	"github.com/skydive-project/skydive/graffiti/getter"
	"github.com/skydive-project/skydive/graffiti/graph"
)

// LogicalRouterPolicy defines the type used by both libovsdb and skydive for table Logical_Router_Policy
// easyjson:json
// gendecoder
type LogicalRouterPolicy struct {
	UUID        string            `ovsdb:"_uuid" json:",omitempty" `
	Action      string            `ovsdb:"action" json:",omitempty" `
	ExternalIDs map[string]string `ovsdb:"external_ids" json:",omitempty" `
	Match       string            `ovsdb:"match" json:",omitempty" `
	Nexthop     []string          `ovsdb:"nexthop" json:",omitempty" `
	Options     map[string]string `ovsdb:"options" json:",omitempty" `
	Priority    int               `ovsdb:"priority" json:",omitempty" `

	ExternalIDsMeta graph.Metadata `json:",omitempty" field:"Metadata"`
	OptionsMeta     graph.Metadata `json:",omitempty" field:"Metadata"`
}

func (t *LogicalRouterPolicy) Metadata() graph.Metadata {
	// Generate Metadata from maps
	t.ExternalIDsMeta = graph.NormalizeValue(t.ExternalIDs).(map[string]interface{})
	t.OptionsMeta = graph.NormalizeValue(t.Options).(map[string]interface{})

	return graph.Metadata{
		"Type":    "LogicalRouterPolicy",
		"Manager": "ovn",
		"UUID":    t.GetUUID(),
		"Name":    t.GetName(),
		"OVN":     t,
	}
}

func (t *LogicalRouterPolicy) GetUUID() string {
	return t.UUID
}

func (t *LogicalRouterPolicy) GetName() string {
	if name := t.UUID; name != "" {
		return name
	}
	return t.GetUUID()
}

// LogicalRouterPolicyDecoder implements a json message raw decoder
func LogicalRouterPolicyDecoder(raw json.RawMessage) (getter.Getter, error) {
	var t LogicalRouterPolicy
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, fmt.Errorf("unable to unmarshal LogicalRouterPolicy metadata %s: %s", string(raw), err)
	}
	return &t, nil
}
