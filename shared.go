package validation

type DatabaseFetcher interface {
	// FetchOne gets a record from table and column
	FetchOne(table string, conditions map[string]interface{}) (interface{}, error)
}
