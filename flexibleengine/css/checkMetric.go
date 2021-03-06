package css

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
	cssId := params[3]
	if cssId == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}

	// Create data for request
	dimension := map[string]interface{}{
		"name":  "cluster_id",
		"value": cssId,
	}
	namespace := "SYS.ES"
	metricsList := []string{metric}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return value[metricsList[0]], nil
}
