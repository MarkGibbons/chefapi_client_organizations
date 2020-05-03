# chefapi_client_organizations

This code provides a go client interface to the chefapi to interact with organizations. In the chefapi demonstration this endpoint is used to populate the pull down list of organizations. See the chefapi_demo_server repository to see how this code was installed and started.

## Endpoints
-----------

## GET /orgs
===========================

### Request
No body is passed

### Return
The body returned looks like this:
````json
[
  "org1",
  "org2"
]
````
Values
* 200 - A list of organizations was returned
* 400 - Invalid request was made
* 401 - The requester was not logged in
* 403 - The requester was not authorized
