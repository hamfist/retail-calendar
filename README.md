retail-calendar
===============

Library that calculates the current fiscal month (drawn from
http://www.nrf.com/modules.php?name=Pages&sp\_id=391)

```bash
import "github.com/modcloth-labs/retail-calendar"

// returns startTime and endTime of the current fiscal Month
startTime, endTime := retail_calendar.MonthRange(time.Now())
```
