package service

import (
	"strconv"

	"github.com/dollarsignteam/go-logger"

	"auzy-api/internal/api/dto"
	"auzy-api/internal/repository"
	"auzy-api/model"
)

type VisitorService struct {
	logger       *logger.Logger
	repository   *repository.Handler
	geoIPService GeoIPService
}

func NewVisitorService(
	logger *logger.Logger,
	repository *repository.Handler,
	geoIPService GeoIPService,
) VisitorService {
	return VisitorService{
		logger:       logger,
		repository:   repository,
		geoIPService: geoIPService,
	}
}

func (svc VisitorService) RecordVisit(ip, userAgent string, req *dto.RecordVisitRequest) error {
	entity := model.VisitorLog{
		IP:        ip,
		CreatedAt: GetUnixTimestamp(),
	}
	if userAgent != "" {
		ua := truncate(userAgent, 512)
		entity.UserAgent = &ua
	}
	if req != nil {
		if req.Path != "" {
			path := truncate(req.Path, 255)
			entity.Path = &path
		}
		if req.Referer != "" {
			referer := truncate(req.Referer, 512)
			entity.Referer = &referer
		}
	}

	if geo, err := svc.geoIPService.Lookup(ip); err == nil && geo != nil {
		meta := &model.LocationMetadata{
			Country: geo.Country,
			Region:  geo.Region,
			City:    geo.City,
			Source:  "ip-api",
		}
		if geo.Latitude != 0 || geo.Longitude != 0 {
			lat := geo.Latitude
			lon := geo.Longitude
			meta.Latitude = &lat
			meta.Longitude = &lon
		}
		entity.Metadata = meta
	}

	if err := svc.repository.CreateVisitorLog(&entity); err != nil {
		svc.logger.Error(err)
		return err
	}
	return nil
}

func (svc VisitorService) ListVisits(page, limit int) (dto.VisitorLogListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	offset := (page - 1) * limit

	logs, err := svc.repository.ListVisitorLogs(offset, limit)
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.VisitorLogListResponse{}, err)
	}
	total, err := svc.repository.CountVisitorLogs()
	if err != nil {
		svc.logger.Error(err)
		return NewCommonErrorSomethingWentWrong(dto.VisitorLogListResponse{}, err)
	}

	items := make([]dto.VisitorLogItem, 0, len(logs))
	for _, log := range logs {
		items = append(items, dto.VisitorLogItem{
			ID:        log.ID,
			IP:        log.IP,
			Metadata:  log.Metadata,
			UserAgent: log.UserAgent,
			Path:      log.Path,
			Referer:   log.Referer,
			CreatedAt: log.CreatedAt,
		})
	}
	return dto.VisitorLogListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func truncate(value string, max int) string {
	if len(value) <= max {
		return value
	}
	return value[:max]
}

func ParsePositiveInt(raw string, fallback int) int {
	if raw == "" {
		return fallback
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n < 1 {
		return fallback
	}
	return n
}
