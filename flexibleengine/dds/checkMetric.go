package dds

import (
	"errors"
	"fmt"
	"strings"

	akskrequest "zabbix.com/plugins/flexibleengine/akskRequest"
)

// CheckMetric verify params, set dimension and namespace values
func CheckMetric(params []string, metric string) (result interface{}, err error) {
	// Verify params
	if len(params) != 9 {
		return nil, errors.New("Wrong parameters.")
	}
	ddsID := params[3]
	if ddsID == "" {
		return nil, fmt.Errorf("Need to specify $INSTANCE_ID option.")
	}
	ddsRole := strings.ToLower(params[8])
	if ddsRole == "" {
		return nil, fmt.Errorf("Need to specify $ROLE option (primary or secondary).")
	}

	// Create data for request
	dimension := map[string]interface{}{
		"name":  "mongod_" + ddsRole + "_instance_id",
		"value": ddsID,
	}
	namespace := "SYS.DDS"
	metricsList := []string{metric}

	value, err := akskrequest.ExecuteProcess(params, dimension, namespace, metricsList)
	if err != nil {
		return nil, err
	}

	return value[metricsList[0]], nil
}
