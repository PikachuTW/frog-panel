package routes

import (
	"frog-panel/internal/services"
	"frog-panel/internal/utils"
	"log"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

type Version struct {
	Version string `json:"version"`
	Type    string `json:"type"`
}

type VanillaVersion struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

const (
	FabricAPIURL = "https://meta.fabricmc.net/v2/versions/game"
	StableType   = "stable"
	UnstableType = "unstable"
)

func fetchVersionTypes(c fuego.ContextNoBody) ([]string, error) {
	return services.SupportLoaders(), nil
}

func fetchVanillaVersion(c fuego.ContextNoBody) ([]Version, error) {
	return []Version{}
}

func fetchFabricVersions(c fuego.ContextNoBody) ([]Version, error) {
	stableQuery := c.QueryParamBool("stable")
	hasStableQuery := c.HasQueryParam("stable")
	includeStable := !hasStableQuery || stableQuery
	includeUnstable := !hasStableQuery || !stableQuery

	type FabricVersion struct {
		Stable  bool   `json:"stable"`
		Version string `json:"version"`
	}

	fabricVersions, err := utils.FetchData[[]FabricVersion](FabricAPIURL)

	if err != nil {
		log.Printf("Failed to fetch Fabric versions: %v", err)
		return []Version{}, fuego.InternalServerError{Err: err}
	}

	if fabricVersions == nil {
		log.Printf("Received nil response from Fabric API")
		return []Version{}, fuego.InternalServerError{Err: err}
	}

	response := make([]Version, 0, len(fabricVersions))

	for _, version := range fabricVersions {
		if (version.Stable && includeStable) || (!version.Stable && includeUnstable) {
			var versionType string
			if version.Stable {
				versionType = StableType
			} else {
				versionType = UnstableType
			}
			response = append(response, Version{
				Version: version.Version,
				Type:    versionType,
			})
		}
	}

	return response, nil
}

func SetupInfoHandlers(s *fuego.Server) {
	fuego.Get(s,
		"/version-types",
		fetchVersionTypes,
		option.Summary("Fetch version types"),
	)

	fuego.Get(
		s,
		"/version/fabric",
		fetchFabricVersions,
		option.Summary("Fetch fabric versions"),
		option.QueryBool("stable", "Filtered by stable"),
	)

}
