package cod

import "time"

// StatData stores JSON for player stats
type StatData struct {
	Data struct {
		PlatformInfo struct {
			PlatformSlug           string      `json:"platformSlug"`
			PlatformUserID         interface{} `json:"platformUserId"`
			PlatformUserHandle     string      `json:"platformUserHandle"`
			PlatformUserIdentifier string      `json:"platformUserIdentifier"`
			AvatarURL              string      `json:"avatarUrl"`
			AdditionalParameters   interface{} `json:"additionalParameters"`
		} `json:"platformInfo"`
		UserInfo struct {
			UserID          interface{}   `json:"userId"`
			IsPremium       bool          `json:"isPremium"`
			IsVerified      bool          `json:"isVerified"`
			IsInfluencer    bool          `json:"isInfluencer"`
			CountryCode     interface{}   `json:"countryCode"`
			CustomAvatarURL interface{}   `json:"customAvatarUrl"`
			CustomHeroURL   interface{}   `json:"customHeroUrl"`
			SocialAccounts  []interface{} `json:"socialAccounts"`
		} `json:"userInfo"`
		Metadata struct {
			LastUpdated struct {
				Value        interface{} `json:"value"`
				DisplayValue interface{} `json:"displayValue"`
			} `json:"lastUpdated"`
			HasPlayedModernWarfare bool `json:"hasPlayedModernWarfare"`
		} `json:"metadata"`
		Segments []struct {
			Type       string `json:"type"`
			Attributes struct {
			} `json:"attributes,omitempty"`
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			ExpiryDate time.Time `json:"expiryDate"`
			Stats      struct {
				Kills struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"kills"`
				Deaths struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"deaths"`
				Downs struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"downs"`
				KdRatio struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"kdRatio"`
				Wins struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"wins"`
				Top5 struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"top5"`
				Top10 struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"top10"`
				Top25 struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"top25"`
				GamesPlayed struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"gamesPlayed"`
				TimePlayed struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"timePlayed"`
				WlRatio struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"wlRatio"`
				Score struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"score"`
				ScorePerMinute struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"scorePerMinute"`
				ScorePerGame struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"scorePerGame"`
				Cash struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"cash"`
				Contracts struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"contracts"`
				AverageLife struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"averageLife"`
				Level struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
						IconURL string `json:"iconUrl"`
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"level"`
				LevelXpTotal struct {
					Rank            interface{} `json:"rank"`
					Percentile      float64     `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"levelXpTotal"`
				WeeklyKills struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"weeklyKills"`
				WeeklyKillsPerGame struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"weeklyKillsPerGame"`
				WeeklyScorePerMinute struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"weeklyScorePerMinute"`
				WeeklyScorePerGame struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"weeklyScorePerGame"`
				WeeklyKdRatio struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"weeklyKdRatio"`
				WeeklyHeadshotPct struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"weeklyHeadshotPct"`
				WeeklyMatchesPlayed struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"weeklyMatchesPlayed"`
				WeeklyDamageDone struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"weeklyDamageDone"`
				WeeklyDamagePerMatch struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"weeklyDamagePerMatch"`
				WeeklyDamagePerMinute struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory string      `json:"displayCategory"`
					Category        string      `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"weeklyDamagePerMinute"`
				LevelProgression struct {
					Rank            interface{} `json:"rank"`
					Percentile      interface{} `json:"percentile"`
					DisplayName     string      `json:"displayName"`
					DisplayCategory interface{} `json:"displayCategory"`
					Category        interface{} `json:"category"`
					Metadata        struct {
					} `json:"metadata"`
					Value        float64 `json:"value"`
					DisplayValue string  `json:"displayValue"`
					DisplayType  string  `json:"displayType"`
				} `json:"levelProgression"`
			} `json:"stats,omitempty"`
		}
		AvailableSegments []interface{} `json:"availableSegments"`
		ExpiryDate        time.Time     `json:"expiryDate"`
	} `json:"data"`
}
