package transparents

import (
	"context"
	"net/url"
	"strconv"
)

const routeListQRCodePIX = "/transparents/list"

func (t *TransparentManager) List(ctx context.Context, query ListQRCodePIXQuery) (*ListQRCodePIXData, error) {
	var data ListQRCodePIXData
	endpoint := routeListQRCodePIX
	if values := query.values(); len(values) > 0 {
		endpoint += "?" + values.Encode()
	}

	if err := t.http.Get(ctx, endpoint, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (t *TransparentManager) Get(ctx context.Context, id string) (*ListQRCodePIXData, error) {
	return t.List(ctx, ListQRCodePIXQuery{ID: id, Limit: 1})
}

func (q ListQRCodePIXQuery) values() url.Values {
	values := url.Values{}
	set := func(key, value string) {
		if value != "" {
			values.Set(key, value)
		}
	}
	set("id", q.ID)
	set("externalId", q.ExternalID)
	set("status", string(q.Status))
	set("method", q.Method)
	set("after", q.After)
	set("before", q.Before)
	set("startDate", q.StartDate)
	set("endDate", q.EndDate)
	if q.Limit > 0 {
		values.Set("limit", strconv.Itoa(q.Limit))
	}
	return values
}
