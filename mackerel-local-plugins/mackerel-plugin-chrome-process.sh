#!/bin/bash

PATH=/bin:/usr/bin:/usr/local/bin

if [[ $MACKEREL_AGENT_PLUGIN_META == "1" ]]; then
    cat <<EOS
# mackerel-agent-plugin
{
  "graphs": {
    "chrome": {
      "label": "Google Chrome",
      "unit": "integer",
      "metrics": [
        {
          "name": "processes",
          "label": "processes",
          "stacked": false
        }
      ]
    }
  }
}
EOS
    exit
fi

now=$(date +%s)
processes=$(ps ax | ps ax | grep 'Google Chrome.app' | grep 'Google Chrome Helper --type=renderer' | wc -l)
printf "chrome.processes\t${processes}\t${now}\n"
