package creative

// CarouselMaterial 图文信息
type CarouselMaterial struct {
	// ID 素材唯一标识
	ID uint64 `json:"id,omitempty"`
	// AwemeCarouselID 抖音图文id
	AwemeCarouselID uint64 `json:"aweme_carousel_id,omitempty"`
	// CarouselID 图文ID
	CarouselID uint64 `json:"carousel_id,omitempty"`
	// IsAutoGenerated 是否为派生创意标识，1：是，0：不是
	IsAutoGenerated int `json:"is_auto_generated,omitempty"`
	// Images 图片信息
	Images []struct {
		// URL 图片url
		URL string `json:"url,omitempty"`
	} `json:"iamges,omitempty"`
	// Audio 音频信息
	Audio []struct {
		// URL 音频url
		URL string `json:"url,omitempty"`
		// Description 图文描述信息
		Description string `json:"description,omitempty"`
	}
}