# chinese-holidays-go

提供具有中国特色的休假安排或者工作日查询。

### Install

    go get github.com/bastengao/chinese-holidays-go


### Usage

```go
import (
    "github.com/bastengao/chinese-holidays-go/holidays"
)

holidays.isHoliday(d)
holidays.isWorkingday(d)
```