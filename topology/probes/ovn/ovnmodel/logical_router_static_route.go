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

// LogicalRouterStaticRoute defines the type used by both libovsdb and skydive for table Logical_Router_Static_Route
// easyjson:json
// gendecoder
type LogicalRouterStaticRoute struct {
	UUID        string            `ovsdb:"_uuid" json:",omitempty" `
	ExternalIDs map[string]string `ovsdb:"external_ids" json:",omitempty" `
	IPPrefix    string            `ovsdb:"ip_prefix" json:",omitempty" `
	Nexthop     string            `ovsdb:"nexthop" json:",omitempty" `
	Options     map[string]string `ovsdb:"options" json:",omitempty" `
	OutputPort  []string          `ovsdb:"output_port" json:",omitempty" `
	Policy      []string          `ovsdb:"policy" json:",omitempty" `

	ExternalIDsMeta graph.Metadata `json:",omitempty" field:"Metadata"`
	OptionsMeta     graph.Metadata `json:",omitempty" field:"Metadata"`
}

func (t *LogicalRouterStaticRoute) Metadata() graph.Metadata {
	// Generate Metadata from maps
	t.ExternalIDsMeta = graph.NormalizeValue(t.ExternalIDs).(map[string]interface{})
	t.OptionsMeta = graph.NormalizeValue(t.Options).(map[string]interface{})

	return graph.Metadata{
		"Type":    "LogicalRouterStaticRoute",
		"Manager": "ovn",
		"UUID":    t.GetUUID(),
		"Name":    t.GetName(),
		"OVN":     t,
	}
}

func (t *LogicalRouterStaticRoute) GetUUID() string {
	return t.UUID
}

func (t *LogicalRouterStaticRoute) GetName() string {
	if name := t.UUID; name != "" {
		return name
	}
	return t.GetUUID()
}

// LogicalRouterStaticRouteDecoder implements a json message raw decoder
func LogicalRouterStaticRouteDecoder(raw json.RawMessage) (getter.Getter, error) {
	var t LogicalRouterStaticRoute
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, fmt.Errorf("unable to unmarshal LogicalRouterStaticRoute metadata %s: %s", string(raw), err)
	}
	return &t, nil
}
