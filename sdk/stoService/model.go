package stoService

type Sto struct {
	Id                     string                 `json:"internal_id"`
	Title                  string                 `json:"title"`
	CategoriesWithServices []CategoryWithServices `json:"services_categories"`
}

type CategoryWithServices struct {
	Id               string          `json:"category_id"`
	Alias            string          `json:"category_alias"`
	Name             string          `json:"category_name"`
	ExcludedServices []string        `json:"excluded_services"`
	IncludedServices []string        `json:"included_services"`
	IsServicesAll    bool            `json:"is_services_all"`
	CustomService    []CustomService `json:"services_custom"`
}

type CustomService struct {
	CustomId           string   `json:"custom_id"`
	Name               string   `json:"name"`
	ReferenceServiceId string   `json:"reference_service_id"`
	Id                 string   `json:"category_id"`
	ChildServices      string   `json:"child_services"`
	Price              Price    `json:"price"`
	Duration           Duration `json:"duration"`
}

type Price struct {
	IsFixed    bool    `json:"is_fixed"`
	FixedPrice float64 `json:"fixed_price"`
	HourPrice  float64 `json:"hour_price"`
}

type Duration struct {
	Calendar int64 `json:"calendar"`
	Default  int64 `json:"default"`
}
