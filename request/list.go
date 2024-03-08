package request

import "github.com/jihanlugas/pandora/config"

type Listing struct {
	Limit int `json:"limit,omitempty" form:"limit" query:"limit"`
}

func (p *Listing) GetLimit() int {
	if p.Limit >= config.MaxDataPerList {
		return config.MaxDataPerList
	} else {
		return p.Limit
	}
}

func (p *Listing) SetLimit(lim int) {
	p.Limit = lim
}
