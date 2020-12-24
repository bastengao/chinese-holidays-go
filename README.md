# chinese-holidays-go

![badge](https://github.com/bastengao/chinese-holidays-go/workflows/Go/badge.svg)

提供具有中国特色的休假安排或者工作日查询。

## Install

    go get github.com/bastengao/chinese-holidays-go

## Usage

```go
import (
    "github.com/bastengao/chinese-holidays-go/holidays"
)

d := time.Date(2019, 10, 1, 0, 0, 0, 0, china)
holidays.isHoliday(d)    // true
holidays.isWorkingday(d) // false
```

## Features

- [x] bundled data
  - support [2021](http://www.gov.cn/zhengce/content/2020-11/25/content_5564127.htm)
  - support [2020](http://www.gov.cn/zhengce/content/2019-11/21/content_5454164.htm)
  - support [2019](http://www.gov.cn/zhengce/content/2018-12/06/content_5346276.htm) and 5.1 [changes](http://www.gov.cn/zhengce/content/2019-03/22/content_5375877.htm)
  - support [2018](http://www.gov.cn/zhengce/content/2017-11/30/content_5243579.htm)
  - support [2017](http://www.gov.cn/zhengce/content/2016-12/01/content_5141603.htm)
  - support 2016
- [ ] online data
