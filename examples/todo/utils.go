package main

func getPrefix(item *Item) string {
	prefix := separator
	if item.focused {
		prefix = ">"
	}

	return prefix
}
