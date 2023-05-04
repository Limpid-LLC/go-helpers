package collection

type Array struct {
	Items []map[string]interface{}
	Pluck []string
	Skip  int
	Limit int
}

func (Array *Array) Get() []map[string]interface{} {
	items := Array.Items

	if Array.Limit > 0 && Array.Skip > 0 {
		items = Array.paginate(items)
	}

	if Array.Pluck != nil {
		items = Array.pluck(items)
	}

	return items
}

func (Array *Array) paginate(Items []map[string]interface{}) []map[string]interface{} {
	itemsLength := len(Items)

	start := Array.Skip * Array.Limit
	if start > itemsLength {
		start = itemsLength
	}

	end := start + Array.Limit
	if end > itemsLength {
		end = itemsLength
	}

	return Array.Items[start:end]
}

func (Array *Array) pluck(Items []map[string]interface{}) []map[string]interface{} {
	var pluckedItems []map[string]interface{}

	for _, itemValue := range Items {
		newItemValue := make(map[string]interface{})

		for fieldKey, fieldValue := range itemValue {
			set := make(map[string]struct{}, len(Array.Pluck))

			for _, s := range Array.Pluck {
				set[s] = struct{}{}
			}

			_, ok := set[fieldKey]

			if ok {
				newItemValue[fieldKey] = fieldValue
			}
		}

		pluckedItems = append(pluckedItems, newItemValue)
	}

	return pluckedItems
}
