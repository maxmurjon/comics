package handler

import (
	"comics/config"
	"comics/storage"
)

type Handler struct {
	cfg  *config.Config
	strg storage.StorageRepoI
}

type Response struct {
	Status      int
	Description string
	Data        interface{}
}

func NewHandler(cfg *config.Config, storage storage.StorageRepoI) *Handler {
	return &Handler{
		cfg:  cfg,
		strg: storage,
	}
}

// func (h *Handler) handlerResponse(c *gin.Context, path string, code int, message interface{}) {
// 	response := Response{
// 		Status:      code,
// 		Description: path,
// 		Data:        message,
// 	}

// 	switch {
// 	case code < 300:
// 		h.logger.Info(path, logger.Any("info", response.Description))
// 	case code >= 400:
// 		h.logger.Error(path, logger.Any("info", response))
// 	}

// 	c.JSON(code, response)
// }

// func (h *Handler) getOffsetQuery(offset string) (int, error) {
// 	if len(offset) <= 0 {
// 		return h.cfg.DefaultOffset, nil
// 	}

// 	return strconv.Atoi(offset)
// }

// func (h *Handler) getLimitQuery(limit string) (int, error) {

// 	if len(limit) <= 0 {
// 		return h.cfg.DefaultLimit, nil
// 	}

// 	return strconv.Atoi(limit)
// }
