package reforger

import (
	"encoding/json"
	"fmt"
	"os"
)

// ConfigTemplate returns a minimal config template with proper structure
const ConfigTemplate = `{
    "bindPort": 2001,
    "publicPort": 2001,
    "a2s": {
        "port": 17777
    },
    "game": {
        "name": "My Arma Reforger Server",
        "password": "",
        "passwordAdmin": "",
        "scenarioId": "{ECC61978EDCC2B5A}Missions/23_Campaign.conf",
        "maxPlayers": 32,
        "visible": true,
        "gameProperties": {
            "serverMaxViewDistance": 2500,
            "serverMinGrassDistance": 50,
            "networkViewDistance": 1500,
            "disableThirdPerson": false,
            "fastValidation": true,
            "battlEye": true
        },
        "mods": []
    }
}`

// ConfigTemplateCommented returns a heavily commented template for user guidance
const ConfigTemplateCommented = `{
    // Optional: IP address to bind to (leave empty for 0.0.0.0)
    // "bindAddress": "",

    // Port for game server (required)
    "bindPort": 2001,

    // Optional: Public IP address (leave empty for auto-detection)
    // "publicAddress": "",

    // Public port (typically same as bindPort)
    "publicPort": 2001,

    // Steam Query (A2S) settings - allows server browsers to query server info
    "a2s": {
        // "address": "", // Optional bind address
        "port": 17777    // A2S query port
    },

    // Optional: RCON (Remote Console) settings
    // "rcon": {
    //     "port": 19999,
    //     "password": "secure_password",
    //     "permission": "admin"  // or "monitor"
    // },

    "game": {
        // Server name displayed in browser
        "name": "My Arma Reforger Server",

        // Password to join (empty = public server)
        "password": "",

        // Admin password for server management
        "passwordAdmin": "",

        // Optional: Array of admin Steam IDs
        // "admins": ["76561198012345678"],

        // Mission/scenario to load
        // Conflict: {ECC61978EDCC2B5A}Missions/23_Campaign.conf
        // Combat Ops: {28E344039A5D6FDC}Missions/21_Combat_Ops.conf
        // Game Master: {59AD59368755F41A}Missions/21_GM_Eden.conf
        "scenarioId": "{ECC61978EDCC2B5A}Missions/23_Campaign.conf",

        // Maximum players (1-128)
        "maxPlayers": 32,

        // Visible in server browser
        "visible": true,

        // Optional: Cross-platform support
        // "crossPlatform": true,
        // "supportedPlatforms": ["PLATFORM_PC", "PLATFORM_XBL", "PLATFORM_PSN"],

        "gameProperties": {
            // Maximum view distance in meters (500-10000)
            "serverMaxViewDistance": 2500,

            // Minimum grass rendering distance
            "serverMinGrassDistance": 50,

            // Network streaming distance in meters (500-5000)
            "networkViewDistance": 1500,

            // Disable third-person camera
            "disableThirdPerson": false,

            // Enable fast validation (faster loading)
            "fastValidation": true,

            // Enable BattlEye anti-cheat
            "battlEye": true

            // Optional VON (Voice Over Network) settings:
            // "VONDisableUI": false,
            // "VONDisableDirectSpeechUI": false,
            // "VONCanTransmitCrossFaction": false
        },

        // Mods (Steam Workshop)
        // Example: [{"modId": "5994AD5A9F33BE57", "name": "Game Master FX"}]
        "mods": []
    }

    // Optional: Server operating settings
    // "operating": {
    //     "lobbyPlayerSynchronise": true,
    //     "aiLimit": 50,
    //     "playerSaveTime": 60
    // }
}`

// GenerateConfigTemplate creates a config template file
func GenerateConfigTemplate(path string, commented bool) error {
	var content string
	if commented {
		content = ConfigTemplateCommented
	} else {
		content = ConfigTemplate
	}

	// For commented version, write as-is (JSON comments for documentation)
	if commented {
		return os.WriteFile(path, []byte(content), 0644)
	}

	// For non-commented, validate it's proper JSON and pretty-print
	var test interface{}
	if err := json.Unmarshal([]byte(content), &test); err != nil {
		return fmt.Errorf("template contains invalid JSON: %w", err)
	}

	// Pretty-print JSON
	formatted, err := json.MarshalIndent(test, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}

	return os.WriteFile(path, formatted, 0644)
}
