package guild

type Guild struct {
	LastModified      int64  `json:"lastModified"`
	Name              string `json:"name"`
	Realm             string `json:"realm"`
	Battlegroup       string `json:"battlegroup"`
	Level             int    `json:"level"`
	Side              int    `json:"side"`
	AchievementPoints int    `json:"achievementPoints"`
	Members           []struct {
		Character struct {
			Name              string `json:"name"`
			Realm             string `json:"realm"`
			Battlegroup       string `json:"battlegroup"`
			Class             int    `json:"class"`
			Race              int    `json:"race"`
			Gender            int    `json:"gender"`
			Level             int    `json:"level"`
			AchievementPoints int    `json:"achievementPoints"`
			Thumbnail         string `json:"thumbnail"`
			Spec              struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
			Guild        string `json:"guild"`
			GuildRealm   string `json:"guildRealm"`
			LastModified int    `json:"lastModified"`
		} `json:"character"`
		Rank int `json:"rank"`
	} `json:"members"`
	Emblem struct {
		Icon              int    `json:"icon"`
		IconColor         string `json:"iconColor"`
		IconColorID       int    `json:"iconColorId"`
		Border            int    `json:"border"`
		BorderColor       string `json:"borderColor"`
		BorderColorID     int    `json:"borderColorId"`
		BackgroundColor   string `json:"backgroundColor"`
		BackgroundColorID int    `json:"backgroundColorId"`
	} `json:"emblem"`
}
