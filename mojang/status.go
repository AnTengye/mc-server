package mojang

import "encoding/json"

// These constants represent the possible states of the Mojang services
const (
	StatusGreen  = "green"
	StatusYellow = "yellow"
	StatusRed    = "red"
)

// Status contains all states of the Mojang services
type Status struct {
	MinecraftWebsite string
	MojangWebsite    string
	Session          string
	SessionServer    string
	AuthServer       string
	Account          string
	Textures         string
	API              string
}

// parseStatusFromBody parses a status object from the response body of the API
func parseStatusFromBody(body []byte) (*Status, error) {
	// Parse multiple single states out of the response body
	var rawStates []map[string]string

	err := json.Unmarshal(body, &rawStates)
	if err != nil {
		return nil, err
	}

	// Create the status object and put the corresponding values in it
	status := new(Status)
	for _, stateMap := range rawStates {
		for key, state := range stateMap {
			switch key {
			case "minecraft.net":
				status.MinecraftWebsite = state
				break
			case "mojang.com":
				status.MojangWebsite = state
				break
			case "session.minecraft.net":
				status.Session = state
				break
			case "sessionserver.mojang.com":
				status.SessionServer = state
				break
			case "authserver.mojang.com":
				status.AuthServer = state
				break
			case "account.mojang.com":
				status.Account = state
				break
			case "textures.minecraft.net":
				status.Textures = state
				break
			case "api.mojang.com":
				status.API = state
				break
			}
		}
	}

	// Return the status object
	return status, nil
}
