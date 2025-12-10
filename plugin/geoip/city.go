package geoip

import (
	"context"
	"strconv"

	"github.com/coredns/coredns/plugin/metadata"

	"github.com/oschwald/geoip2-golang/v2"
)

func (g GeoIP) setCityMetadata(ctx context.Context, data *geoip2.City) {
	// Set labels for city, country and continent names.
	// In v2, Names is a struct with language fields instead of a map.
	cityName := data.City.Names.English
	metadata.SetValueFunc(ctx, pluginName+"/city/name", func() string {
		return cityName
	})
	countryName := data.Country.Names.English
	metadata.SetValueFunc(ctx, pluginName+"/country/name", func() string {
		return countryName
	})
	continentName := data.Continent.Names.English
	metadata.SetValueFunc(ctx, pluginName+"/continent/name", func() string {
		return continentName
	})

	// In v2, IsoCode is renamed to ISOCode.
	countryCode := data.Country.ISOCode
	metadata.SetValueFunc(ctx, pluginName+"/country/code", func() string {
		return countryCode
	})
	isInEurope := strconv.FormatBool(data.Country.IsInEuropeanUnion)
	metadata.SetValueFunc(ctx, pluginName+"/country/is_in_european_union", func() string {
		return isInEurope
	})
	continentCode := data.Continent.Code
	metadata.SetValueFunc(ctx, pluginName+"/continent/code", func() string {
		return continentCode
	})

	// In v2, Latitude and Longitude are pointers to properly distinguish
	// between missing coordinates and the valid location (0, 0).
	var latitude string
	if data.Location.Latitude != nil {
		latitude = strconv.FormatFloat(*data.Location.Latitude, 'f', -1, 64)
	}
	metadata.SetValueFunc(ctx, pluginName+"/latitude", func() string {
		return latitude
	})
	var longitude string
	if data.Location.Longitude != nil {
		longitude = strconv.FormatFloat(*data.Location.Longitude, 'f', -1, 64)
	}
	metadata.SetValueFunc(ctx, pluginName+"/longitude", func() string {
		return longitude
	})
	timeZone := data.Location.TimeZone
	metadata.SetValueFunc(ctx, pluginName+"/timezone", func() string {
		return timeZone
	})
	postalCode := data.Postal.Code
	metadata.SetValueFunc(ctx, pluginName+"/postalcode", func() string {
		return postalCode
	})
}
