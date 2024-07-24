package nav

type Breadcrumb struct {
	Title  string
	URL    string
	Nolink bool `default:"false"`
}
