/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package config

func NewFCTemplateFromFCTemplateJsonCfg(jsnCfg *FcTemplateJsonCfg) *FCTemplate {
	fcTmp := new(FCTemplate)
	if jsnCfg.Id != nil {
		fcTmp.ID = *jsnCfg.Id
	}
	if jsnCfg.Type != nil {
		fcTmp.Type = *jsnCfg.Type
	}
	if jsnCfg.Field_id != nil {
		fcTmp.FieldId = *jsnCfg.Field_id
	}
	if jsnCfg.Filters != nil {
		fcTmp.Filters = make([]string, len(*jsnCfg.Filters))
		for i, fltr := range *jsnCfg.Filters {
			fcTmp.Filters[i] = fltr
		}
	}
	if jsnCfg.Value != nil {
		fcTmp.Value = NewRSRParsersMustCompile(*jsnCfg.Value, true)
	}
	if jsnCfg.Width != nil {
		fcTmp.Width = *jsnCfg.Width
	}
	if jsnCfg.Strip != nil {
		fcTmp.Strip = *jsnCfg.Strip
	}
	if jsnCfg.Padding != nil {
		fcTmp.Padding = *jsnCfg.Padding
	}
	if jsnCfg.Mandatory != nil {
		fcTmp.Mandatory = *jsnCfg.Mandatory
	}
	if jsnCfg.Attribute_id != nil {
		fcTmp.AttributeID = *jsnCfg.Attribute_id
	}
	if jsnCfg.New_branch != nil {
		fcTmp.NewBranch = *jsnCfg.New_branch
	}
	if jsnCfg.Timezone != nil {
		fcTmp.Timezone = *jsnCfg.Timezone
	}
	if jsnCfg.Blocker != nil {
		fcTmp.Blocker = *jsnCfg.Blocker
	}
	if jsnCfg.Break_on_success != nil {
		fcTmp.BreakOnSuccess = *jsnCfg.Break_on_success
	}
	return fcTmp
}

type FCTemplate struct {
	ID             string
	Type           string   // Type of field
	FieldId        string   // Field identifier
	Filters        []string // list of filter profiles
	Value          RSRParsers
	Width          int
	Strip          string
	Padding        string
	Mandatory      bool
	AttributeID    string // Used by NavigableMap when creating CGREvent/XMLElements
	NewBranch      bool   // Used by NavigableMap when creating XMLElements
	Timezone       string
	Blocker        bool
	BreakOnSuccess bool
}

func FCTemplatesFromFCTemapltesJsonCfg(jsnCfgFlds []*FcTemplateJsonCfg) []*FCTemplate {
	retFields := make([]*FCTemplate, len(jsnCfgFlds))
	for i, jsnFld := range jsnCfgFlds {
		retFields[i] = NewFCTemplateFromFCTemplateJsonCfg(jsnFld)
	}
	return retFields
}