package info

import (
	"frog-panel/internal/client"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

type VersionResponse struct {
	Version string `json:"version"`
	Type    string `json:"type"`
}

type FabricVersion struct {
	Stable  bool   `json:"stable"`
	Version string `json:"version"`
}

type VanillaVersion struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

func Register(r *gin.RouterGroup) {
	r.GET("/versions-types", func(ctx *gin.Context) {
		versions := []string{"fabric", "vanilla"}
		ctx.JSON(http.StatusOK, versions)
	})

	r.GET("/versions/fabric", func(ctx *gin.Context) {
		typeQueries := ctx.QueryArray("type")
		showStable := len(typeQueries) == 0
		showUnstable := len(typeQueries) == 0
		for _, typeQuery := range typeQueries {
			switch typeQuery {
			case "stable":
				showStable = true
			case "unstable":
				showUnstable = true
			}
		}
		httpClient := client.New()
		var versions []FabricVersion
		if err := httpClient.GetJSON("https://meta.fabricmc.net/v2/versions/game", &versions); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Fabric versions"})
			return
		}
		versionResponses := make([]VersionResponse, 0, len(versions))
		for _, version := range versions {
			if (showStable && version.Stable) || (showUnstable && !version.Stable) {
				versionResponse := VersionResponse{
					Version: version.Version,
				}
				if version.Stable {
					versionResponse.Type = "stable"
				} else {
					versionResponse.Type = "unstable"
				}
				versionResponses = append(versionResponses, versionResponse)
			}
		}
		ctx.JSON(http.StatusOK, versionResponses)
	})

	r.GET("/versions/vanilla", func(ctx *gin.Context) {
		typeQueries := ctx.QueryArray("type")
		httpClient := client.New()
		var vanillaVersions struct {
			Versions []VanillaVersion `json:"versions"`
		}
		if err := httpClient.GetJSON("https://launchermeta.mojang.com/mc/game/version_manifest.json", &vanillaVersions); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Vanilla versions"})
			return
		}
		versionResponses := make([]VersionResponse, 0, len(vanillaVersions.Versions))
		for _, version := range vanillaVersions.Versions {
			shouldResponse := len(typeQueries) == 0 || slices.Contains(typeQueries, version.Type)
			if shouldResponse {
				versionResponses = append(versionResponses, VersionResponse{
					Version: version.Id,
					Type:    version.Type,
				})
			}
		}
		ctx.JSON(http.StatusOK, versionResponses)
	})
}
