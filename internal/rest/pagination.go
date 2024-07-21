package rest

// Pagination represents a set of pagination links for a list of resources.
type Pagination struct {
	Previous *string `json:"previous"`
	Next     *string `json:"next"`
}
