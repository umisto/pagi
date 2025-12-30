package pagi

import (
	"net/http"
	"strconv"
	"strings"
)

func GetLimit(
	r *http.Request,
	max uint,
) uint {
	limit := uint(20)

	if s := r.URL.Query().Get("limit"); s != "" {
		parsed, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			limit = uint(parsed)
		}
	}

	if limit == 0 {
		limit = 20
	}
	if limit > max {
		limit = max
	}

	return limit
}

func GetSort(r *http.Request) *SortField {
	sortStr := strings.TrimSpace(r.URL.Query().Get("sort"))
	if sortStr == "" {
		return nil
	}

	sort := &SortField{
		Ascend: true,
	}

	if strings.HasPrefix(sortStr, "-") {
		sort.Ascend = false
		sort.Field = strings.TrimPrefix(sortStr, "-")
	} else {
		sort.Field = sortStr
	}

	sort.Field = strings.TrimSpace(sort.Field)
	if sort.Field == "" {
		return nil
	}

	return sort
}
