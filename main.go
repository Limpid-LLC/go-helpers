package collection

type Collection struct {
	Items []map[string]interface{}
	Pluck []string
	Skip  int
	Limit int
}

func (Collection *Collection) Get() []map[string]interface{} {
	items := Collection.Items

	if Collection.Limit > 0 && Collection.Skip > 0 {
		items = Collection.paginate(items)
	}

	if Collection.Pluck != nil {
		items = Collection.pluck(items)
	}

	return items
}

func (Collection *Collection) paginate(Items []map[string]interface{}) []map[string]interface{} {
	itemsLength := len(Items)

	start := Collection.Skip * Collection.Limit
	if start > itemsLength {
		start = itemsLength
	}

	end := start + Collection.Limit
	if end > itemsLength {
		end = itemsLength
	}

	return Collection.Items[start:end]
}

func (Collection *Collection) pluck(Items []map[string]interface{}) []map[string]interface{} {
	var pluckedItems []map[string]interface{}

	for _, itemValue := range Items {
		newItemValue := make(map[string]interface{})

		for fieldKey, fieldValue := range itemValue {
			set := make(map[string]struct{}, len(Collection.Pluck))

			for _, s := range Collection.Pluck {
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
