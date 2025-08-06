package domain

type RouteResponseHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RouteItem struct {
	ID          int    `json:"id"`
	Method      string `json:"method"`
	Path        string `json:"path"`
	Description string `json:"description"`

	ResponseCode    int                   `json:"responseCode"`
	ResponseHeaders []RouteResponseHeader `json:"responseHeaders"`
	ResponseBody    string                `json:"responseBody"`
}

type RouteItemRequestData struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Description string `json:"description"`

	ResponseCode    int                   `json:"responseCode"`
	ResponseHeaders []RouteResponseHeader `json:"responseHeaders"`
	ResponseBody    string                `json:"responseBody"`
}
