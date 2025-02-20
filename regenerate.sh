#!/bin/sh

wget https://api.sumologic.com/docs/sumologic-api.yaml

rm -rf api docs

openapi-generator-cli generate -c config.yaml

rm sumologic-api.yaml
mv api/docs .

# Workaround for https://github.com/OpenAPITools/openapi-generator/issues/19394
sed -i '' -e 's/var defaultValue \[\]string = \["\(.*\)"]/var defaultValue []string = []string{"\1"}/g' api/api_connection_management.go

# Workaround for https://github.com/OpenAPITools/openapi-generator/issues/19959
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_alerts_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_collector_registration_token_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_monitor_templates_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_monitor_templates_library_monitor_template_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_monitors_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_monitors_library_monitor_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_muting_schedules_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_muting_schedules_library_muting_schedule_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_parsers_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_parsers_library_parser_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_service_now_connection.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_slos_library_folder_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_slos_library_slo_response.go
sed -i '' $'/import (/a\\\n\t"time"\\\n' api/model_webhook_connection.go

# Workaround for incorrect type assignment
# Not sure if there is already an open issue for this or not
sed -i '' -e 's/this.TimeRange = timeRange/this.TimeRange = \&timeRange/' api/model_logs_anomaly_condition.go
sed -i '' -e 's/this.AnomalyDetectorType = anomalyDetectorType/this.AnomalyDetectorType = \&anomalyDetectorType/' api/model_logs_anomaly_condition.go
sed -i '' -e 's/this.Field = field/this.Field = \&field/' api/model_logs_anomaly_condition.go
sed -i '' -e 's/this.TimeRange = timeRange/this.TimeRange = \&timeRange/' api/model_metrics_anomaly_condition.go
sed -i '' -e 's/this.AnomalyDetectorType = anomalyDetectorType/this.AnomalyDetectorType = \&anomalyDetectorType/' api/model_metrics_anomaly_condition.go

# verify that the SKD builds
go build
