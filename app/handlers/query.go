package handlers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type listQuery struct {
	ID        *int
	Sort      string
	Order     string
	SortValue string
}

func parseListQuery(r *http.Request, defaultSort, defaultOrder string, allowedSorts map[string]string) (listQuery, error) {
	query := r.URL.Query()
	result := listQuery{Sort: defaultSort, Order: defaultOrder}

	if rawID := query.Get("id"); rawID != "" {
		id, err := strconv.Atoi(rawID)
		if err != nil || id < 0 {
			return listQuery{}, fmt.Errorf("id must be a non-negative integer")
		}
		result.ID = &id
	}

	if rawSort := query.Get("sort"); rawSort != "" {
		result.Sort = rawSort
	}
	column, ok := allowedSorts[result.Sort]
	if !ok {
		return listQuery{}, fmt.Errorf("sort must be one of: %s", allowedSortNames(allowedSorts))
	}
	result.SortValue = column

	if rawOrder := query.Get("order"); rawOrder != "" {
		result.Order = rawOrder
	}
	if result.Order != "asc" && result.Order != "desc" {
		return listQuery{}, fmt.Errorf("order must be asc or desc")
	}
	return result, nil
}

func allowedSortNames(allowedSorts map[string]string) string {
	names := make([]string, 0, len(allowedSorts))
	for name := range allowedSorts {
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}

func queryError(w http.ResponseWriter, err error) {
	writeError(w, http.StatusBadRequest, err.Error())
}
