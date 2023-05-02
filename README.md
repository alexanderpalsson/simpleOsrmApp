## How to use

# Run server
`Make server` will start a simple rest server on port :8080

### API spec </br> 
* `GET /routes`  returns the fastest duration with car between a source point and multiple destinations. 

    Query Params:
  - `src` - source location (customer's home), only one can be provided.
  - `dst` - destination location (pickup point), multiple can be provided.

Example query: `http://localhost:8080/routes?src=13.388860,52.517037&dst=13.428555,52.523219&dst=13.397634,52.529407`


# Run tests
`Make test-unit` will run all unit-tests
