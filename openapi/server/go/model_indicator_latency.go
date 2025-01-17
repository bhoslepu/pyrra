/*
 * Pyrra
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.3.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type IndicatorLatency struct {
	Success Query `json:"success"`

	Total Query `json:"total"`

	Grouping []string `json:"grouping,omitempty"`
}

// AssertIndicatorLatencyRequired checks if the required fields are not zero-ed
func AssertIndicatorLatencyRequired(obj IndicatorLatency) error {
	elements := map[string]interface{}{
		"success": obj.Success,
		"total":   obj.Total,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertQueryRequired(obj.Success); err != nil {
		return err
	}
	if err := AssertQueryRequired(obj.Total); err != nil {
		return err
	}
	return nil
}

// AssertRecurseIndicatorLatencyRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of IndicatorLatency (e.g. [][]IndicatorLatency), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseIndicatorLatencyRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aIndicatorLatency, ok := obj.(IndicatorLatency)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertIndicatorLatencyRequired(aIndicatorLatency)
	})
}
