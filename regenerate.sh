#!/bin/sh

# Install the OpenAPI generator CLI:
#   npm install @openapitools/openapi-generator-cli -g
# Run this script, from within this directory, to regenerate the SDK from the Sumo Logic OpenAPI spec:
#   ./regenerate.sh
# Commit and push:
#   git add .
#   git commit -m "Update SDK to match latest Sumo Logic OpenAPI spec"
#   git push

wget https://api.sumologic.com/docs/sumologic-api.yaml

rm -rf ./*.go docs test

openapi-generator-cli generate -c config.yaml

rm sumologic-api.yaml

# Workaround for https://github.com/OpenAPITools/openapi-generator/issues/19394
sed -i '' -e 's/var defaultValue \[\]string = \["\(.*\)"]/var defaultValue []string = []string{"\1"}/g' api_connection_management.go

# Workaround for https://github.com/OpenAPITools/openapi-generator/issues/19959
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_alerts_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_collector_registration_token_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_monitor_templates_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_monitor_templates_library_monitor_template_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_monitors_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_monitors_library_monitor_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_muting_schedules_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_muting_schedules_library_muting_schedule_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_parsers_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_parsers_library_parser_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_service_now_connection.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_slos_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_slos_library_slo_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' model_webhook_connection.go

# Workaround for incorrect type assignment
# Not sure if there is already an open issue for this or not
sed -i '' -e 's/this.TimeRange = timeRange/this.TimeRange = \&timeRange/' model_logs_anomaly_condition.go
sed -i '' -e 's/this.AnomalyDetectorType = anomalyDetectorType/this.AnomalyDetectorType = \&anomalyDetectorType/' model_logs_anomaly_condition.go
sed -i '' -e 's/this.Field = field/this.Field = \&field/' model_logs_anomaly_condition.go
sed -i '' -e 's/this.TimeRange = timeRange/this.TimeRange = \&timeRange/' model_metrics_anomaly_condition.go
sed -i '' -e 's/this.AnomalyDetectorType = anomalyDetectorType/this.AnomalyDetectorType = \&anomalyDetectorType/' model_metrics_anomaly_condition.go

# verify that the SKD builds
go build

echo ".idea" >> .gitignore
