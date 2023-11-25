package search_projects

type Args struct {
	User        *string `json:"user,omitempty"`
	Hashtag     *string `json:"hashtag,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	Description *string `json:"description,omitempty"`
	Fuzziness   *int    `json:"fuzziness,omitempty"`
}
