package employeeService

type Employee struct {
	Id                     string                   `json:"internal_id"`
	StoId                  string                   `json:"sto_id"`
	UserId                 any                      `json:"user_id"`
	Name                   string                   `json:"name"`
	Phone                  string                   `json:"phone"`
	RoleId                 string                   `json:"role_id"`
	HireDate               string                   `json:"hire_date"`
	Schedule               []ScheduleDay            `json:"schedule"`
	CategoriesWithServices []CategoriesWithServices `json:"categories_with_services"`
}

type ScheduleDay struct {
	PositionOfWeek string `json:"position_of_week"`
	TimeFrom       string `json:"time_from"`
	TimeTo         string `json:"time_to"`
}

type CategoriesWithServices struct {
	CategoryID           string          `json:"category_id"`
	CategoryAlias        string          `json:"category_alias"`
	CategoryName         string          `json:"category_name"`
	IsServicesAll        bool            `json:"is_services_all"`
	ServicesCustom       []CustomService `json:"services_custom"`
	IncludedServiceIds   []string        `json:"included_services_ids"`
	IncludedServiceNames []string        `json:"included_services_names"`
	ExcludedServiceIds   []string        `json:"excluded_services_ids"`
	ExcludedServiceNames []string        `json:"excluded_services_names"`
}

type CustomService struct {
	CategoryId string `json:"category_id"`
	Name       string `json:"name"`
}
