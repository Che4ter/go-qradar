package qradar

import (
	"context"
	"net/http"
)

// RuleWithDataService handles methods related to RuleWithData of the QRadar API.
type RuleWithDataService service

const ruleWithDataAPIPrefix = "api/analytics/rules_with_data"

// RuleWithData represents QRadar's RuleWithData. Undocumented.
type RuleWithData struct {
	Rule
	CRE             *int    `json:"cre,omitempty"`
	TypeID          *int    `json:"type_id,omitempty"`
	IsBuildingBlock *bool   `json:"is_building_block,omitempty"`
	RuleXML         *string `json:"rule_xml,omitempty"`
}

// Get returns RulesWithData of the current QRadar installation. Undocumented.
func (c *RuleWithDataService) Get(ctx context.Context, fields, filter string, from, to int) ([]RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleWithDataAPIPrefix, fields, filter, from, to, nil, nil)
	if err != nil {
		return nil, err
	}
	var result []RuleWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Create creates RulesWithData in the current QRadar installation. Undocumented.
func (c *RuleWithDataService) Create(ctx context.Context, fields string, data interface{}) (*RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodPost, ruleWithDataAPIPrefix, fields, "", 0, 0, nil, data)
	if err != nil {
		return nil, err
	}
	var result RuleWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID returns RulesWithData of the current QRadar installation by ID. Undocumented.
func (c *RuleWithDataService) GetByID(ctx context.Context, fields string, id int) (*RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodGet, ruleWithDataAPIPrefix, fields, "", 0, 0, &id, nil)
	if err != nil {
		return nil, err
	}
	var result RuleWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByID updates RulesWithData of the current QRadar installation by ID. Undocumented.
func (c *RuleWithDataService) UpdateByID(ctx context.Context, fields string, id int, data interface{}) (*RuleWithData, error) {
	req, err := c.client.requestHelp(http.MethodPost, ruleWithDataAPIPrefix, fields, "", 0, 0, &id, data)
	if err != nil {
		return nil, err
	}
	var result RuleWithData
	_, err = c.client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
