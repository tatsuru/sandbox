package mposxbattery

import (
	"fmt"
	"os/exec"
	"strings"
	"regexp"
	"strconv"
	mp "github.com/mackerelio/go-mackerel-plugin"
)

type OSXBatteryPlugin struct {
	Prefix string
}

func (o OSXBatteryPlugin) MetricKeyPrefix() string {
	if o.Prefix == "" {
		o.Prefix = "battery"
	}
	return o.Prefix
}

func (o OSXBatteryPlugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(o.MetricKeyPrefix())
	return map[string]mp.Graphs{
		"": {
			Label: labelPrefix,
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "percentage", Label: "%"},
			},
		},
	}
}

func (o OSXBatteryPlugin) FetchMetrics() (map[string]float64, error) {
	percentage, err := getBatteryPercentage()
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch battery metrics: %s", err)
	}
	return map[string]float64{"percentage": percentage}, nil
}

func getBatteryPercentage() (float64, error) {
	output, err := exec.Command("pmset", "-g", "ps").Output()
	if err != nil {
		return 0.0, fmt.Errorf("`pmset -g ps` command failed: %s", err)
	}
	//  -InternalBattery-0     28%; discharging; 1:31 remaining present: true
	re := regexp.MustCompile("\\s+[^\\s]+\\s+([0-9\\.]+)%;.+")
	percentage := 0.0
	for _, line := range strings.Split(string(output), "\n") {
		match := re.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		percentage, err = strconv.ParseFloat(match[1], 64)
	}
	return percentage, nil
}

// Do the plugin
func Do() {
	o := OSXBatteryPlugin{}
	plugin := mp.NewMackerelPlugin(o)
	plugin.Run()
}
