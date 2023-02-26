package api

import "sort"

func init() {
	// If any method is able to change this list, remeber to call once sort.
	sort.Strings(AiModelList)
}
