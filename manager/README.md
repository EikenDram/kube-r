# Manager service

Go service (API) for:

- backend for web-ui
- creating requests into message query
- monitoring message query about results and sending messages to applications
- etc

Should have database for storing data about report requests

## API

| NAME      | DESCRIPTION
|-----------|-----------------------
| create    | Sends a message to MQ for creating report
| state     | Change state of request in database, makes API call to sender if state is final

## Environment variables

| NAME       | DESCRIPTION
|------------|------------------------

