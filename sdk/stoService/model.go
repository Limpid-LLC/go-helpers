package stoService

type Sto struct {
	Id                     string                 `json:"internal_id"`
	Title                  string                 `json:"title"`
	CategoriesWithServices []CategoryWithServices `json:"services_categories"`
}

type CategoryWithServices struct {
	Id               string        `json:"CategoryID"`
	Alias            string        `json:"CategoryAlias"`
	Name             string        `json:"CategoryName"`
	ExcludedServices []string      `json:"ExcludedServices"`
	IncludedServices []string      `json:"IncludedServices"`
	IsServicesAll    bool          `json:"IsServicesAll"`
	CustomService    CustomService `json:"ServicesCustom"`
}

type CustomService struct {
	Id                 string `json:"ServiceID"`
	Name               string `json:"Name"`
	ReferenceServiceId string `json:"ReferenceServiceId"`
}
