# custom-lambda
go-api which works as a trigger between two different services

It is useful if your one application is running and you want to trigger an event which would run on some other application

From your First application you can send a curl like this with required parameters

curl -s -k -X POST -H "Content-Type: application/json" \
-d '{"mode":"orders","job_type":"test","query_mode":"","gcs_bucket": "some_bucket",
"gcs_file_name":"file_name.gz",
"is_remove_lock":"False",
"use_redis_lock":"True",
"etl_id":"noawttphdy-20240118-195316-etl-tool"}' \
https://your-host/etl-app/txn

