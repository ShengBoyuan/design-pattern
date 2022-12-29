package proxy

import "fmt"

// Service Interface
type ThirdPartyYTBLib interface {
	ListVideos() []int
	GetVideoInfo(int) string
	DownloadVideo(int) bool
}

// Service
type ThirdPartyYTBClass struct{}

var _ ThirdPartyYTBLib = ThirdPartyYTBClass{}

func (ThirdPartyYTBClass) ListVideos() []int {
	videoIds := make([]int, 10)
	for i := 0; i < 10; i++ {
		videoIds[i] = i
	}
	return videoIds
}

func (ThirdPartyYTBClass) GetVideoInfo(id int) string {
	return fmt.Sprint("videoId:", id)
}

func (ThirdPartyYTBClass) DownloadVideo(id int) bool {
	return id < 10
}

// Proxy
type CachedYTBClass struct {
	ListCache  []int
	VideoCache map[int]string
	service    ThirdPartyYTBLib
}

var _ ThirdPartyYTBLib = &CachedYTBClass{}

func (c *CachedYTBClass) Init(service ThirdPartyYTBClass) {
	c.ListCache = make([]int, 10)
	c.VideoCache = make(map[int]string)
	c.service = service
}

func (c *CachedYTBClass) ListVideos() []int {
	if len(c.ListCache) == 0 {
		c.ListCache = c.service.ListVideos()
	}
	return c.ListCache
}

func (c *CachedYTBClass) GetVideoInfo(id int) string {
	if info, ok := c.VideoCache[id]; ok {
		return info
	} else {
		info := c.service.GetVideoInfo(id)
		c.VideoCache[id] = info
		return info
	}
}

func (c *CachedYTBClass) DownloadVideo(id int) bool {
	return c.service.DownloadVideo(id)
}

// Client
type YTBManager struct {
	service ThirdPartyYTBLib
}

func (y *YTBManager) Init(service ThirdPartyYTBLib) {
	y.service = service
}

func (y *YTBManager) RenderListPanel() {
	list := y.service.ListVideos()
	fmt.Println(list)
}

func (y *YTBManager) RenderVideoPage(id int) {
	fmt.Println(y.service.GetVideoInfo(id))
}

func (y *YTBManager) ReactOnUserInput() {
	y.RenderListPanel()
	y.RenderVideoPage(5)
}

// Main Processing Flow
func SeeSomethigFunOnYTB() {
	actualService := ThirdPartyYTBClass{}
	// init proxy
	ytbProxy := &CachedYTBClass{}
	ytbProxy.Init(actualService)

	// init client
	ytbManager := &YTBManager{}
	ytbManager.Init(ytbProxy)

	// do
	ytbManager.ReactOnUserInput()
}
