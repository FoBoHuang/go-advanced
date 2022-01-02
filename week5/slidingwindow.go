package main

import (
	"log"
	"sync"
	"time"
)

type counterSlidingWindow struct {
	windowSize    int64 // 滑动窗口大小 /s
	splitNum      int64 // 切分窗口的数目大小，每个窗口对应一个桶存储数据
	currentBucket int   // 当前的桶
	limit         int   // 滑动窗口内限流大小
	Bucket        []int //存放每个窗口内的计数
	startTime     int64 // 滑动窗口开始时间
}

func NewWindow(windowSize int64, limit int, spiltNum int64) *counterSlidingWindow {
	return &counterSlidingWindow{
		windowSize:    windowSize,
		splitNum:      spiltNum,
		currentBucket: 0,
		limit:         limit,
		Bucket:        make([]int, spiltNum),
		startTime:     time.Now().Unix(),
	}
}

func (c *counterSlidingWindow) tryAcquire() bool {
	currentTime := time.Now().Unix()
	t := currentTime - c.windowSize - c.startTime
	if t < 0 {
		t = 0
	}
	windowsNum := t / (c.windowSize / c.splitNum)
	c.slideWindow(windowsNum)
	count := 0
	for i := 0; i < int(c.splitNum); i++ {
		count += c.Bucket[i]
	}
	log.Printf("当前滑动窗口总数为: %d", count)
	if count > c.limit {
		log.Println("开始限流")
		return false
	}
	index := c.getCurrentBucket()
	log.Println("当前的bucket: ", index)
	c.Bucket[index]++
	return true
}

func (c *counterSlidingWindow) getCurrentBucket() int {
	currentTime := time.Now().Unix()
	t := int(currentTime / (c.windowSize / c.splitNum) % c.splitNum)
	return t
}

func (c *counterSlidingWindow) slideWindow(windowsNum int64) {
	// 滑动窗口默认设置为 splitNum
	var slideNum int64 = c.splitNum
	if windowsNum == 0 {
		return
	}
	log.Println("当前要滑动的窗口数: ", slideNum)
	for i := 0; i < int(slideNum); i++ {
		// 根据splitNum取余，获取当前的bucket
		c.currentBucket = (c.currentBucket + 1) % int(c.splitNum)
		log.Printf("当前清空的位置: %d, 当前的大小: %d", c.currentBucket, c.Bucket[c.currentBucket])
		c.Bucket[c.currentBucket] = 0
	}
	c.startTime = c.startTime + windowsNum*(c.windowSize/c.splitNum)
}

func main() {
	c := NewWindow(10, 5, 5)
	c.tryAcquire()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i <= 100; i++ {
			for j := 0; j < 1; j++ {
				if c.tryAcquire() {
				}
			}
			time.Sleep(time.Second * 6)
		}
	}()
	wg.Wait()
}
