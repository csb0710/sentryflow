package api

import (
	"sync"
)

// aa Local reference for API analyzer
var aa *APIAnalyzer

// init function
func init() {
	aa = NewAPIAnalyzer()
}

// APIAnalyzer Structure
type APIAnalyzer struct {
	perAPICount     map[string]uint64
	perAPICountLock sync.Mutex // @todo perhaps combine those two?

	stopChan chan struct{}
	apiJob   chan string
}

// NewAPIAnalyzer Function
func NewAPIAnalyzer() *APIAnalyzer {
	ret := &APIAnalyzer{
		perAPICount: make(map[string]uint64),
	}

	return ret
}

// StartAPIAnalyzer Function
func StartAPIAnalyzer(wg *sync.WaitGroup) {
	go apiAnalyzerRoutine(wg)
}

// StopAPIAnalyzer Function
func StopAPIAnalyzer() {
	aa.stopChan <- struct{}{}
}

// apiAnalyzerRoutine Function
func apiAnalyzerRoutine(wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		select {
		case job, ok := <-aa.apiJob:
			if !ok {
				// @todo perhaps error message here?
				continue
			}
			analyzeAPI(job)

		case <-aa.stopChan:
			wg.Done()
			break
		}
	}
}

// analyzeAPI Function
func analyzeAPI(api string) {
	// @todo implement this
	classifyAPI(api)
}

// GetPerAPICount Function
func GetPerAPICount() map[string]uint64 {
	aa.perAPICountLock.Lock()
	ret := aa.perAPICount
	aa.perAPICountLock.Unlock()

	return ret
}

// UpdatePerAPICount Function
func UpdatePerAPICount(nm map[string]uint64) {
	aa.perAPICountLock.Lock()
	aa.perAPICount = nm
	aa.perAPICountLock.Unlock()
}

// InsertAnalyzeJob Function
func InsertAnalyzeJob(api string) {
	aa.apiJob <- api
}