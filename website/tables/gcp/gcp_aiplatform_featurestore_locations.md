# Table: gcp_aiplatform_featurestore_locations

This table shows data for GCP AI Platform Feature Store Locations.

https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_aiplatform_featurestore_locations:
  - [gcp_aiplatform_featurestores](gcp_aiplatform_featurestores)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|location_id|`utf8`|
|display_name|`utf8`|
|labels|`json`|
|metadata|`json`|