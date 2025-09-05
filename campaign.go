package patreon

var (
	// CampaignDefaultIncludes specifies default includes for Campaign.
	CampaignDefaultIncludes = []string{"tiers", "creator", "benefits", "goals"}
	// CampaignFields is all fields in the Campaign Attributes struct
	CampaignFields = getObjectFields(CampaignAttributes{})
)

// Campaign is the creator's page, and the top-level object for accessing lists of members, tiers, etc.
type Campaign struct {
	Type          string             `json:"type"`
	ID            string             `json:"id"`
	Attributes    CampaignAttributes `json:"attributes"`
	Relationships struct {
		Benefits              *BenefitsRelationship   `json:"benefits,omitempty"`
		CampaignInstallations *interface{}            `json:"campaign_installations,omitempty"`
		Categories            *CategoriesRelationship `json:"categories,omitempty"`
		Creator               *CreatorRelationship    `json:"creator,omitempty"`
		Goals                 *GoalsRelationship      `json:"goals,omitempty"`
		Tiers                 *TiersRelationship      `json:"tiers,omitempty"`
	} `json:"relationships"`
}

// CampaignV1 is the creator's page, and the top-level object for accessing lists of members, tiers, etc.
type CampaignV1 struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Summary                       string   `json:"summary"`
		Currency                      string   `json:"currency"`
		CreationName                  string   `json:"creation_name"`
		DisplayPatronGoals            bool     `json:"display_patron_goals"`
		PayPerName                    string   `json:"pay_per_name"`
		OneLiner                      string   `json:"one_liner"`
		MainVideoEmbed                string   `json:"main_video_embed"`
		MainVideoURL                  string   `json:"main_video_url"`
		ImageSmallURL                 string   `json:"image_small_url"`
		ImageURL                      string   `json:"image_url"`
		ThanksVideoURL                string   `json:"thanks_video_url"`
		ThanksEmbed                   string   `json:"thanks_embed"`
		ThanksMsg                     string   `json:"thanks_msg"`
		IsChargedImmediately          bool     `json:"is_charged_immediately"`
		IsMonthly                     bool     `json:"is_monthly"`
		IsNsfw                        bool     `json:"is_nsfw"`
		IsPlural                      bool     `json:"is_plural"`
		CreatedAt                     NullTime `json:"created_at"`
		PublishedAt                   NullTime `json:"published_at"`
		PledgeURL                     string   `json:"pledge_url"`
		PledgeSum                     int      `json:"pledge_sum"`
		PatronCount                   int      `json:"patron_count"`
		CreationCount                 int      `json:"creation_count"`
		OutstandingPaymentAmountCents int      `json:"outstanding_payment_amount_cents"`
	} `json:"attributes"`
	Relationships struct {
		Categories      *CategoriesRelationship      `json:"categories,omitempty"`
		Creator         *CreatorRelationship         `json:"creator,omitempty"`
		Rewards         *RewardsRelationship         `json:"rewards,omitempty"`
		Goals           *GoalsRelationship           `json:"goals,omitempty"`
		PostAggregation *PostAggregationRelationship `json:"post_aggregation,omitempty"`
	} `json:"relationships"`
}

// CampaignAttributes is the attributes struct for Campaign
type CampaignAttributes struct {
	CreatedAt            NullTime `json:"created_at"`
	CreationName         string   `json:"creation_name"`
	DiscordServerID      string   `json:"discord_server_id"`
	GoogleAnalyticsID    string   `json:"google_analytics_id"`
	HasRSS               bool     `json:"has_rss"`
	HasSentRSSNotify     bool     `json:"has_sent_rss_notify"`
	ImageSmallURL        string   `json:"image_small_url"`
	ImageURL             string   `json:"image_url"`
	IsChargedImmediately bool     `json:"is_charged_immediately"`
	IsMonthly            bool     `json:"is_monthly"`
	IsNsfw               bool     `json:"is_nsfw"`
	MainVideoEmbed       string   `json:"main_video_embed"`
	MainVideoURL         string   `json:"main_video_url"`
	OneLiner             string   `json:"one_liner"`
	PatronCount          int      `json:"patron_count"`
	PayPerName           string   `json:"pay_per_name"`
	PledgeURL            string   `json:"pledge_url"`
	PublishedAt          NullTime `json:"published_at"`
	RSSArtworkURL        bool     `json:"rss_artwork_url"`
	RSSFeedTitle         string   `json:"rss_feed_title"`
	ShowEarnings         bool     `json:"show_earnings"`
	Summary              string   `json:"summary"`
	ThanksEmbed          string   `json:"thanks_embed"`
	ThanksMsg            string   `json:"thanks_msg"`
	ThanksVideoURL       string   `json:"thanks_video_url"`
	URL                  string   `json:"url"`
	Vanity               string   `json:"vanity"`
}

// CampaignV2Response wraps Patreon's campaign API response
type CampaignResponse struct {
	Data     Campaign `json:"data"`
	Included Includes `json:"included"`
}

// CampaignV1Response wraps Patreon's campaign API response
type CampaignV1Response struct {
	Data     CampaignV1 `json:"data"`
	Included Includes   `json:"included"`
}

// CampaignsV2Response wraps Patreon's campaign API response
type CampaignsResponse struct {
	Data     []Campaign `json:"data"`
	Included Includes   `json:"included"`
}
