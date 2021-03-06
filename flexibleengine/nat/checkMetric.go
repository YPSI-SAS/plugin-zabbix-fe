package nat

import (
	"errors"
	"fmt"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

// CheckMetric verify params, set dimension and namespace values
func CheckMetric(params []string, metric string) (result interface{}, err error) {
	// Verify params
	if len(params) != 8 {
		return nil, errors.New("Wrong parameters.")
	}
	natID := params[3]
	if natID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	// Create data for request
	dimension := map[string]interface{}{
		"name":  "nat_gateway_id",
		"value": natID,
	}
	namespace := "SYS.NAT"
	metricsList := []string{metric}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return value[metricsList[0]], nil
}
