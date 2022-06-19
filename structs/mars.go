package structs

import (
	"encoding/json"
	"fmt"
	"time"
)

type AtmosphericTemperature struct {
	Av float64 `json:"av"`
	Ct int     `json:"ct"`
	Mn float64 `json:"mn"`
	Mx float64 `json:"mx"`
}

type HWS struct {
	Av float64 `json:"av"`
	Ct int     `json:"ct"`
	Mn float64 `json:"mn"`
	Mx float64 `json:"mx"`
}

type PRE struct {
	Av float64 `json:"av"`
	Ct int     `json:"ct"`
	Mn float64 `json:"mn"`
	Mx float64 `json:"mx"`
}

type Compass struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}

type Sol struct {
	AT       AtmosphericTemperature `json:"AT"`
	FirstUTC time.Time              `json:"First_UTC"`
	HWS      HWS                    `json:"HWS"`
	LastUTC  time.Time              `json:"Last_UTC"`
	PRE      PRE                    `json:"PRE"`
	Season   string                 `json:"Season"`
	WD       map[string]Compass     `json:"WD"`
}
type FriendlyResponse struct {
	Sols        map[string]Sol `json:"sols"`
	HottestSol  Sol            `json:"hottest_sol"`
	WindiestSol Sol            `json:"windiest_sol"`
	HeaviestSol Sol            `json:"heaviest_sol"`
}

func (swr SolWeatherResponse) ToFriendlyResponse() FriendlyResponse {
	j, err := json.Marshal(swr)
	if err != nil {
		fmt.Printf("error serializing SolWeatherResponse %+v", err)
		return FriendlyResponse{}
	}

	var solsInfo map[string]Sol
	err = json.Unmarshal(j, &solsInfo)
	if err != nil {
		fmt.Printf("error Deserializing SolWeatherResponse Data %+v", err)
		return FriendlyResponse{}
	}

	h, w, heaviest := getTheMostSols(solsInfo)

	return FriendlyResponse{
		Sols:        solsInfo,
		HottestSol:  h,
		WindiestSol: w,
		HeaviestSol: heaviest,
	}
}
func getTheMostSols(
	solsInfo map[string]Sol,
) (hottest Sol, windiest Sol, heaviest Sol) {
	for _, v := range solsInfo {
		if hottest.AT.Av == 0 {
			hottest = v
		} else if v.AT.Av > hottest.AT.Av {
			hottest = v
		}
		if v.HWS.Av > windiest.HWS.Av {
			windiest = v
		}
		if v.PRE.Av > heaviest.PRE.Av {
			heaviest = v
		}
	}
	return hottest, windiest, heaviest
}

type SolWeatherResponse struct {
	NUM001 *Sol `json:"1,omitempty"`
	NUM002 *Sol `json:"2,omitempty"`
	// Shorten…
	NUM687 *Sol `json:"687,omitempty"`
}
