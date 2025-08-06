package domain

type RouteResponseHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RouteItem struct {
	ID          int    `json:"id"`
	Path        string `json:"path"`
	Description string `json:"description"`

	ResponseCode   int                   `json:"responseCode"`
	ResponseHeader []RouteResponseHeader `json:"responseHeader"`
	ResponseBody   string                `json:"responseBody"`
}

type RouteItemRequestData struct {
	Path        string `json:"path"`
	Description string `json:"description"`

	ResponseCode    int                   `json:"responseCode"`
	ResponseHeaders []RouteResponseHeader `json:"responseHeaders"`
	ResponseBody    string                `json:"responseBody"`
}
