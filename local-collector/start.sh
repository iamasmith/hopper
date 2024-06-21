#!/bin/bash
docker run -it --name otel -v $(pwd):/conf -p 4317:4317 -p 4318:4318 otel/opentelemetry-collector:0.49.0 /otelcol --config=/conf/config.yaml