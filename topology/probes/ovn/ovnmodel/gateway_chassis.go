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

// GatewayChassis defines the type used by both libovsdb and skydive for table Gateway_Chassis
// easyjson:json
// gendecoder
type GatewayChassis struct {
	UUID        string            `ovsdb:"_uuid" json:",omitempty" `
	ChassisName string            `ovsdb:"chassis_name" json:",omitempty" `
	ExternalIDs map[string]string `ovsdb:"external_ids" json:",omitempty" `
	Name        string            `ovsdb:"name" json:",omitempty" `
	Options     map[string]string `ovsdb:"options" json:",omitempty" `
	Priority    int               `ovsdb:"priority" json:",omitempty" `

	ExternalIDsMeta graph.Metadata `json:",omitempty" field:"Metadata"`
	OptionsMeta     graph.Metadata `json:",omitempty" field:"Metadata"`
}

func (t *GatewayChassis) Metadata() graph.Metadata {
	// Generate Metadata from maps
	t.ExternalIDsMeta = graph.NormalizeValue(t.ExternalIDs).(map[string]interface{})
	t.OptionsMeta = graph.NormalizeValue(t.Options).(map[string]interface{})

	return graph.Metadata{
		"Type":    "GatewayChassis",
		"Manager": "ovn",
		"UUID":    t.GetUUID(),
		"Name":    t.GetName(),
		"OVN":     t,
	}
}

func (t *GatewayChassis) GetUUID() string {
	return t.UUID
}

func (t *GatewayChassis) GetName() string {
	if name := t.Name; name != "" {
		return name
	}
	return t.GetUUID()
}

// GatewayChassisDecoder implements a json message raw decoder
func GatewayChassisDecoder(raw json.RawMessage) (getter.Getter, error) {
	var t GatewayChassis
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, fmt.Errorf("unable to unmarshal GatewayChassis metadata %s: %s", string(raw), err)
	}
	return &t, nil
}
