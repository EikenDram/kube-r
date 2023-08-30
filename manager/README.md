# Manager service

Go service (API) for:

- backend for web-ui
- creating report requests and sending into message query
- updating report request status and calling application api for notifying about status change

Should have database for storing data about report requests, will use DB2 server for now

## API

| NAME      | DESCRIPTION
|-----------|-----------------------
| create    | Add report to db, send a message to MQ
| update    | Change state of request in database, make API call to application for notifying about status change

## Environment variables

| NAME       | DESCRIPTION
|------------|------------------------

