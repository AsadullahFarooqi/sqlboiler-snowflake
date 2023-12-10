package main

import (
	"github.com/asadullahfarooqi/sqlboiler-snowflake/driver"
	"github.com/volatiletech/sqlboiler/v4/drivers"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
