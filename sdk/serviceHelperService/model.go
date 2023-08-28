package serviceHelperService

type Names struct {
	En string `json:"en"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Pl string `json:"pl"`
	Ru string `json:"ru"`
	Uk string `json:"uk"`
	Tr string `json:"tr"`
}

type Service struct {
	Id            string         `json:"internal_id"`
	ServiceId     string         `json:"service_id"`
	StoId         string         `json:"sto_id"`
	Names         Names          `json:"names"`
	Type          string         `json:"type"`
	Price         Price          `json:"price"`
	Duration      Duration       `json:"duration"`
	VehicleTypes  []string       `json:"vehicleTypes"`
	CategoryAlias []string       `json:"categoryAlias"`
	ChildServices []ChildService `json:"child_services"`
}

type ChildService struct {
	Id                  string   `json:"internal_id"`
	Names               Names    `json:"names"`
	CategoryAlias       []string `json:"categoryAlias"`
	CustomId            string   `json:"custom_id"`
	Name                string   `json:"name"`
	CustomCategoryAlias string   `json:"category_aliases"`
}

type Price struct {
	IsFixed    bool    `json:"isFixed"`
	FixedPrice float64 `json:"fixedPrice"`
	HourPrice  float64 `json:"hourPrice"`
}

type Duration struct {
	Calendar int64 `json:"calendar"`
	Default  int64 `json:"default"`
}
