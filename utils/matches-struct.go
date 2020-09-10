package utils

import (
	"time"
)

// MatchData stores the JSON response for match data
type MatchData struct {
	Data struct {
		Matches []struct {
			Attributes struct {
				ID     string `json:"id"`
				MapID  string `json:"mapId"`
				ModeID string `json:"modeId"`
			} `json:"attributes"`
			Metadata struct {
				Duration struct {
					Value        int    `json:"value"`
					DisplayValue string `json:"displayValue"`
					DisplayType  string `json:"displayType"`
				} `json:"duration"`
				Timestamp   time.Time `json:"timestamp"`
				PlayerCount int       `json:"playerCount"`
				TeamCount   int       `json:"teamCount"`
				MapName     string    `json:"mapName"`
				MapImageURL string    `json:"mapImageUrl"`
				ModeName    string    `json:"modeName"`
			} `json:"metadata"`
			Segments []struct {
				Type       string `json:"type"`
				Attributes struct {
					PlatformUserIdentifier string      `json:"platformUserIdentifier"`
					PlatformSlug           interface{} `json:"platformSlug"`
					Team                   string      `json:"team"`
				} `json:"attributes"`
				Metadata struct {
					PlatformUserHandle string `json:"platformUserHandle"`
					ClanTag            string `json:"clanTag"`
					Placement          int    `json:"placement"`
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
						Value        float64 `json:"value"`
						DisplayValue string  `json:"displayValue"`
						DisplayType  string  `json:"displayType"`
					} `json:"kills"`
					MedalXp struct {
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
					} `json:"medalXp"`
					MatchXp struct {
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
					} `json:"matchXp"`
					ScoreXp struct {
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
					} `json:"scoreXp"`
					WallBangs struct {
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
					} `json:"wallBangs"`
					Score struct {
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
					} `json:"score"`
					TotalXp struct {
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
					} `json:"totalXp"`
					Headshots struct {
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
					} `json:"headshots"`
					Assists struct {
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
					} `json:"assists"`
					ChallengeXp struct {
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
					} `json:"challengeXp"`
					ScorePerMinute struct {
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
					} `json:"scorePerMinute"`
					DistanceTraveled struct {
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
					} `json:"distanceTraveled"`
					TeamSurvivalTime struct {
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
					} `json:"teamSurvivalTime"`
					Deaths struct {
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
					} `json:"deaths"`
					KdRatio struct {
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
					} `json:"kdRatio"`
					BonusXp struct {
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
					} `json:"bonusXp"`
					GulagDeaths struct {
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
					} `json:"gulagDeaths"`
					TimePlayed struct {
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
					} `json:"timePlayed"`
					Executions struct {
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
					} `json:"executions"`
					GulagKills struct {
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
					} `json:"gulagKills"`
					Nearmisses struct {
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
					} `json:"nearmisses"`
					ObjectiveBrCacheOpen struct {
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
					} `json:"objectiveBrCacheOpen"`
					PercentTimeMoving struct {
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
					} `json:"percentTimeMoving"`
					MiscXp struct {
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
					} `json:"miscXp"`
					LongestStreak struct {
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
					} `json:"longestStreak"`
					TeamPlacement struct {
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
					} `json:"teamPlacement"`
					DamageDone struct {
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
					} `json:"damageDone"`
					DamageTaken struct {
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
					} `json:"damageTaken"`
					DamageDonePerMinute struct {
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
					} `json:"damageDonePerMinute"`
					Placement struct {
						Rank            interface{} `json:"rank"`
						Percentile      interface{} `json:"percentile"`
						DisplayName     string      `json:"displayName"`
						DisplayCategory interface{} `json:"displayCategory"`
						Category        interface{} `json:"category"`
						Metadata        struct {
						} `json:"metadata"`
						Value        int    `json:"value"`
						DisplayValue string `json:"displayValue"`
						DisplayType  string `json:"displayType"`
					} `json:"placement"`
				} `json:"stats"`
			} `json:"segments"`
		} `json:"matches"`
		Metadata struct {
			Next int64 `json:"next"`
		} `json:"metadata"`
		PaginationType             int `json:"paginationType"`
		RequestingPlayerAttributes struct {
			PlatformUserIdentifier string `json:"platformUserIdentifier"`
		} `json:"requestingPlayerAttributes"`
	} `json:"data"`
}
