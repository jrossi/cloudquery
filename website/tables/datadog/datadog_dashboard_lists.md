# Table: datadog_dashboard_lists

This table shows data for Datadog Dashboard Lists.

The composite primary key for this table is (**account_name**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_name (PK)|`utf8`|
|id (PK)|`int64`|
|author|`json`|
|created|`timestamp[us, tz=UTC]`|
|dashboard_count|`int64`|
|is_favorite|`bool`|
|modified|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|type|`utf8`|
|additional_properties|`json`|