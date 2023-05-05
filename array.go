package helpers

type Array struct {
	Items  []map[string]interface{}
	fields []string
}

func (Array *Array) SetFields(fields []string) {
	Array.fields = fields
}

func (Array *Array) Get() []map[string]interface{} {
	items := Array.Items

	if Array.fields != nil {
		items = Array.pluck(items)
	}

	return items
}

func (Array *Array) pluck(Items []map[string]interface{}) []map[string]interface{} {
	var pluckedItems []map[string]interface{}

	for _, itemValue := range Items {
		newItemValue := make(map[string]interface{})

		for fieldKey, fieldValue := range itemValue {
			set := make(map[string]struct{}, len(Array.fields))

			for _, s := range Array.fields {
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
