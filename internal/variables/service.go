package variables

import (
	"net/http"
)

type ByHub map[string]map[string]string

var (
	ErrMissingQueryParams     = HTTPError{"missing hub or variable name", http.StatusBadRequest}
	ErrNoManualVariablesFound = HTTPError{"no hubs with manual variables found", http.StatusNotFound}
	ErrHubNotFound            = HTTPError{"hub not found", http.StatusNotFound}
	ErrVariableNotFound       = HTTPError{"variable not found", http.StatusNotFound}
)

func GetVarValue(hubName string, varName string, hubsVariables ByHub) (string, error) {
	if len(hubsVariables) == 0 {
		return "", ErrNoManualVariablesFound
	}

	if hubName == "" || varName == "" {
		return "", ErrMissingQueryParams
	}

	hubFound := hubsVariables[hubName]
	if hubFound == nil {
		return "", ErrHubNotFound
	}

	varValue, ok := hubFound[varName]
	if !ok {
		return "", ErrVariableNotFound
	}

	return varValue, nil
}
