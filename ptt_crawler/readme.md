## go ptt爬蟲
利用goroutine來爬ptt資料
### Usage
爬表特版有`[正妹]`標籤的照片，可指定要存的的頁數以及存檔的位置
```go
go run pttCrawler.go
```
檔案會存在img中.

### Todo
1. 輸入看板網址
2. 輸入要title, 照片, 內文頁數或篇數
3. 決定輸出方式
4. 使用shared map來暫存已爬過的圖 
5. 提供自然語言指令來爬

