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

// DHCPOptions defines the type used by both libovsdb and skydive for table DHCP_Options
// easyjson:json
// gendecoder
type DHCPOptions struct {
	UUID        string            `ovsdb:"_uuid" json:",omitempty" `
	Cidr        string            `ovsdb:"cidr" json:",omitempty" `
	ExternalIDs map[string]string `ovsdb:"external_ids" json:",omitempty" `
	Options     map[string]string `ovsdb:"options" json:",omitempty" `

	ExternalIDsMeta graph.Metadata `json:",omitempty" field:"Metadata"`
	OptionsMeta     graph.Metadata `json:",omitempty" field:"Metadata"`
}

func (t *DHCPOptions) Metadata() graph.Metadata {
	// Generate Metadata from maps
	t.ExternalIDsMeta = graph.NormalizeValue(t.ExternalIDs).(map[string]interface{})
	t.OptionsMeta = graph.NormalizeValue(t.Options).(map[string]interface{})

	return graph.Metadata{
		"Type":    "DHCPOptions",
		"Manager": "ovn",
		"UUID":    t.GetUUID(),
		"Name":    t.GetName(),
		"OVN":     t,
	}
}

func (t *DHCPOptions) GetUUID() string {
	return t.UUID
}

func (t *DHCPOptions) GetName() string {
	if name := t.UUID; name != "" {
		return name
	}
	return t.GetUUID()
}

// DHCPOptionsDecoder implements a json message raw decoder
func DHCPOptionsDecoder(raw json.RawMessage) (getter.Getter, error) {
	var t DHCPOptions
	if err := json.Unmarshal(raw, &t); err != nil {
		return nil, fmt.Errorf("unable to unmarshal DHCPOptions metadata %s: %s", string(raw), err)
	}
	return &t, nil
}
