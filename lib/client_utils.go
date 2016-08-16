package lib

type Character struct {
	AchievementPoints int    `json:"achievementPoints"`
	Battlegroup       string `json:"battlegroup"`
	CalcClass         string `json:"calcClass"`
	Class             int    `json:"class"`
	Faction           int    `json:"faction"`
	Gender            int    `json:"gender"`
	Items             struct {
		AverageItemLevel         int `json:"averageItemLevel"`
		AverageItemLevelEquipped int `json:"averageItemLevelEquipped"`
		Back                     struct {
			Appearance struct {
				EnchantDisplayInfoID int `json:"enchantDisplayInfoId"`
			} `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []int         `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				Enchant         int `json:"enchant"`
				TimewalkerLevel int `json:"timewalkerLevel"`
				Upgrade         struct {
					Current            int `json:"current"`
					ItemLevelIncrement int `json:"itemLevelIncrement"`
					Total              int `json:"total"`
				} `json:"upgrade"`
			} `json:"tooltipParams"`
		} `json:"back"`
		Chest struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []int         `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
		} `json:"chest"`
		Feet struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []interface{} `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				Set             []int `json:"set"`
				TimewalkerLevel int   `json:"timewalkerLevel"`
				Upgrade         struct {
					Current            int `json:"current"`
					ItemLevelIncrement int `json:"itemLevelIncrement"`
					Total              int `json:"total"`
				} `json:"upgrade"`
			} `json:"tooltipParams"`
		} `json:"feet"`
		Finger1 struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []int         `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
				Upgrade         struct {
					Current            int `json:"current"`
					ItemLevelIncrement int `json:"itemLevelIncrement"`
					Total              int `json:"total"`
				} `json:"upgrade"`
			} `json:"tooltipParams"`
		} `json:"finger1"`
		Finger2 struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []interface{} `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
		} `json:"finger2"`
		Hands struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []interface{} `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
		} `json:"hands"`
		Head struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []interface{} `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
		} `json:"head"`
		Legs struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []interface{} `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
		} `json:"legs"`
		MainHand struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []int         `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
				Upgrade         struct {
					Current            int `json:"current"`
					ItemLevelIncrement int `json:"itemLevelIncrement"`
					Total              int `json:"total"`
				} `json:"upgrade"`
			} `json:"tooltipParams"`
			WeaponInfo struct {
				Damage struct {
					ExactMax int `json:"exactMax"`
					ExactMin int `json:"exactMin"`
					Max      int `json:"max"`
					Min      int `json:"min"`
				} `json:"damage"`
				Dps         float64 `json:"dps"`
				WeaponSpeed float64 `json:"weaponSpeed"`
			} `json:"weaponInfo"`
		} `json:"mainHand"`
		Neck struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []int         `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
				Upgrade         struct {
					Current            int `json:"current"`
					ItemLevelIncrement int `json:"itemLevelIncrement"`
					Total              int `json:"total"`
				} `json:"upgrade"`
			} `json:"tooltipParams"`
		} `json:"neck"`
		Shoulder struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []interface{} `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				Set             []int `json:"set"`
				TimewalkerLevel int   `json:"timewalkerLevel"`
				Upgrade         struct {
					Current            int `json:"current"`
					ItemLevelIncrement int `json:"itemLevelIncrement"`
					Total              int `json:"total"`
				} `json:"upgrade"`
			} `json:"tooltipParams"`
		} `json:"shoulder"`
		Tabard struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []interface{} `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []interface{} `json:"stats"`
			TooltipParams        struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
		} `json:"tabard"`
		Trinket1 struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []interface{} `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
			} `json:"tooltipParams"`
		} `json:"trinket1"`
		Trinket2 struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []int         `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
				Upgrade         struct {
					Current            int `json:"current"`
					ItemLevelIncrement int `json:"itemLevelIncrement"`
					Total              int `json:"total"`
				} `json:"upgrade"`
			} `json:"tooltipParams"`
		} `json:"trinket2"`
		Waist struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []int         `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
				Upgrade         struct {
					Current            int `json:"current"`
					ItemLevelIncrement int `json:"itemLevelIncrement"`
					Total              int `json:"total"`
				} `json:"upgrade"`
			} `json:"tooltipParams"`
		} `json:"waist"`
		Wrist struct {
			Appearance           struct{}      `json:"appearance"`
			Armor                int           `json:"armor"`
			ArtifactAppearanceID int           `json:"artifactAppearanceId"`
			ArtifactID           int           `json:"artifactId"`
			ArtifactTraits       []interface{} `json:"artifactTraits"`
			BonusLists           []int         `json:"bonusLists"`
			Context              string        `json:"context"`
			DisplayInfoID        int           `json:"displayInfoId"`
			Icon                 string        `json:"icon"`
			ID                   int           `json:"id"`
			ItemLevel            int           `json:"itemLevel"`
			Name                 string        `json:"name"`
			Quality              int           `json:"quality"`
			Relics               []interface{} `json:"relics"`
			Stats                []struct {
				Amount int `json:"amount"`
				Stat   int `json:"stat"`
			} `json:"stats"`
			TooltipParams struct {
				TimewalkerLevel int `json:"timewalkerLevel"`
				Upgrade         struct {
					Current            int `json:"current"`
					ItemLevelIncrement int `json:"itemLevelIncrement"`
					Total              int `json:"total"`
				} `json:"upgrade"`
			} `json:"tooltipParams"`
		} `json:"wrist"`
	} `json:"items"`
	LastModified        int    `json:"lastModified"`
	Level               int    `json:"level"`
	Name                string `json:"name"`
	Race                int    `json:"race"`
	Realm               string `json:"realm"`
	Thumbnail           string `json:"thumbnail"`
	TotalHonorableKills int    `json:"totalHonorableKills"`
}
