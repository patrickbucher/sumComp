package sumcomp

import (
	"errors"
	"fmt"
	"math/rand"
)

type SummaryRequest struct {
	DataId int
	Reader chan SummaryResponse
}

type SummaryResponse struct {
	Summary Summary
	Error   error
}

type DataIdsRequest struct {
	Reader chan DataIdResponse
}

type DataIdResponse struct {
	DataIds []int
}

type Cache struct {
	abstractData    map[int]Summary
	publisher       chan Summary
	summaryRequests chan SummaryRequest
	dataIdsRequests chan DataIdsRequest
}

func NewCache(nSummaries int) *Cache {
	cache := Cache{
		abstractData:    make(map[int]Summary, 0),
		publisher:       make(chan Summary),
		summaryRequests: make(chan SummaryRequest),
		dataIdsRequests: make(chan DataIdsRequest),
	}
	for i := 0; i < nSummaries; i++ {
		summary := RandomSummary()
		cache.abstractData[summary.DataId] = summary
	}
	go func() {
		for {
			select {
			case summary := <-cache.publisher:
				cache.abstractData[summary.DataId] = summary
			case request := <-cache.summaryRequests:
				if summary, ok := cache.abstractData[request.DataId]; ok {
					request.Reader <- SummaryResponse{summary, nil}
				} else {
					err := fmt.Errorf("no summary with dataId %d found", request.DataId)
					request.Reader <- SummaryResponse{Summary{}, err}
				}
			case request := <-cache.dataIdsRequests:
				dataIds := make([]int, 0)
				for dataId := range cache.abstractData {
					dataIds = append(dataIds, dataId)
				}
				request.Reader <- DataIdResponse{dataIds}
			}
		}
	}()
	return &cache
}

func (c *Cache) Publish(summary Summary) {
	c.publisher <- summary
}

func (c *Cache) GetDataIds() []int {
	ch := make(chan DataIdResponse)
	c.dataIdsRequests <- DataIdsRequest{ch}
	res := <-ch
	return res.DataIds
}

func (c *Cache) GetSummary(dataId int) (Summary, error) {
	ch := make(chan SummaryResponse)
	c.summaryRequests <- SummaryRequest{dataId, ch}
	res := <-ch
	return res.Summary, res.Error
}

func (c *Cache) GetRandomDataId() (int, error) {
	dataIds := c.GetDataIds()
	if len(dataIds) <= 0 {
		return -1, errors.New("no dataIds available")
	}
	return dataIds[rand.Intn(len(dataIds))], nil
}
