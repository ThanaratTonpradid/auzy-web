package service

import (
	"fmt"
	"net"
	"strings"

	"github.com/dollarsignteam/go-logger"
	"github.com/imroc/req/v3"

	"auzy-api/internal/api/constant"
	"auzy-api/internal/repository"
)

type GeoLocation struct {
	Country   string  `json:"country"`
	Region    string  `json:"regionName"`
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Status    string  `json:"status"`
	Message   string  `json:"message"`
}

type GeoIPService struct {
	logger     *logger.Logger
	repository *repository.Handler
	httpClient *req.Client
}

func NewGeoIPService(
	logger *logger.Logger,
	repository *repository.Handler,
	httpClient *req.Client,
) GeoIPService {
	return GeoIPService{
		logger:     logger,
		repository: repository,
		httpClient: httpClient,
	}
}

func (svc GeoIPService) Lookup(ip string) (*GeoLocation, error) {
	ip = strings.TrimSpace(ip)
	if ip == "" || isPrivateOrLocalIP(ip) {
		return nil, nil
	}

	cacheKey := repository.GetKeyGeoIP(ip)
	cached := GeoLocation{}
	if err := svc.repository.JSONGet(cacheKey, &cached); err == nil && cached.Status == "success" {
		return &cached, nil
	}

	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,country,regionName,city,lat,lon", ip)
	var result GeoLocation
	resp, err := svc.httpClient.R().
		SetSuccessResult(&result).
		Get(url)
	if err != nil {
		svc.logger.Warnf("geoip lookup failed for %s: %v", ip, err)
		return nil, err
	}
	if !resp.IsSuccessState() || result.Status != "success" {
		msg := result.Message
		if msg == "" {
			msg = "geoip lookup unsuccessful"
		}
		svc.logger.Warnf("geoip lookup unsuccessful for %s: %s", ip, msg)
		return nil, fmt.Errorf("%s", msg)
	}

	_ = svc.repository.JSONSet(cacheKey, result, constant.TTLGeoIPCache)
	return &result, nil
}

func isPrivateOrLocalIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return true
	}
	return ip.IsLoopback() || ip.IsPrivate() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast()
}
