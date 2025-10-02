package services

import (
	"frog-panel/config"
	"frog-panel/internal/utils"
	"slices"
	"strings"
	"time"
)

func ModpackCategories() ([]string, error) {
	type CategoryItem struct {
		Icon        string `json:"icon"`
		Name        string `json:"name"`
		ProjectType string `json:"project_type"`
		Header      string `json:"header"`
	}

	categoryItems, err := utils.FetchData[[]CategoryItem](config.ModrinthApiUrl + "/tag/category")

	if err != nil {
		return []string{}, err
	}

	var categories []string

	for _, categoryItem := range categoryItems {
		if categoryItem.ProjectType == "modpack" {
			categories = append(categories, categoryItem.Name)
		}
	}

	return categories, nil
}

type Modpack struct {
	ProjectId         string    `json:"project_id"`
	ProjectType       string    `json:"project_type"`
	Slug              string    `json:"slug"`
	Author            string    `json:"author"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Categories        []string  `json:"categories"`
	DisplayCategories []string  `json:"display_categories"`
	Versions          []string  `json:"versions"`
	Downloads         uint      `json:"downloads"`
	Follows           uint      `json:"follows"`
	IconUrl           string    `json:"icon_url"`
	DateCreated       time.Time `json:"date_created"`
	DateModified      time.Time `json:"date_modified"`
	LatestVersion     string    `json:"latest_version"`
	License           string    `json:"license"`
	ClientSide        string    `json:"client_side"`
	ServerSide        string    `json:"server_side"`
	Gallary           []string  `json:"gallary"`
	FeaturedGallary   string    `json:"featured_gallary"`
	Color             int       `json:"color"`
}

type ModpacksParams struct {
	Categories []string
	Versions   []string
	Loaders    []string
	Offset     uint
	Limit      uint
	Query      string
}

func buildFacets(categories []string, loaders []string, versions []string) [][]string {
	var facets [][]string

	if len(categories) > 0 {
		categoryFacets := make([]string, len(categories))
		for i, cat := range categories {
			categoryFacets[i] = "categories:" + cat
		}
		facets = append(facets, categoryFacets)
	}

	if len(loaders) > 0 {
		loaderFacets := make([]string, len(loaders))
		for i, loader := range loaders {
			loaderFacets[i] = "categories:" + loader
		}
		facets = append(facets, loaderFacets)
	}

	if len(versions) > 0 {
		versionFacets := make([]string, len(versions))
		for i, ver := range versions {
			versionFacets[i] = "versions:" + ver
		}
		facets = append(facets, versionFacets)
	}

	return facets
}

func Modpacks(params ModpacksParams) ([]Modpack, error) {
	i := 0
	for _, loader := range params.Loaders {
		if slices.Contains(SupportLoaders(), strings.ToLower(loader)) {
			params.Loaders[i] = loader
			i++
		}
	}
	params.Loaders = params.Loaders[:i]

	queryParams := make(map[string]string)

	if params.Query != "" {
		queryParams["query"] = params.Query
	}

	searchResults, err := utils.FetchData[[]Modpack](config.ModrinthApiUrl+"/search", queryParams)

	if err != nil {
		return []Modpack{}, err
	}

	return searchResults, nil
}
