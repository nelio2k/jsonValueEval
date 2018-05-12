package jsonValueEval

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/buger/jsonparser"
)

type FlexibleMatcher struct{}

func NewFlexibleMatcher() *FlexibleMatcher {
	return &FlexibleMatcher{}
}

// Returns true if it's a match. Error if something went wrong.
func (m *FlexibleMatcher) Match(data []byte, expression *govaluate.EvaluableExpression, parameters map[string]interface{}) (bool, error) {

	// name["first"]="Neil" OR (age<50 AND isActive=True)
	firstName, err := jsonparser.GetString(data, "name", "first")
	if err != nil {
		fmt.Printf("GetString Error: %v\n", err.Error())
		return false, err
	}

	age, err := jsonparser.GetInt(data, "age")
	if err != nil {
		fmt.Printf("GetInt Error: %v\n", err.Error())
		return false, err
	}

	isActive, err := jsonparser.GetBoolean(data, "isActive")
	if err != nil {
		fmt.Printf("GetBoolean Error: %v\n", err.Error())
		return false, err
	}

	parameters["firstName"] = firstName
	parameters["age"] = age
	parameters["isActive"] = isActive

	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Printf("Evaluate Error: %v\n", err.Error())
		return false, err
	}
	return result.(bool), err
}
