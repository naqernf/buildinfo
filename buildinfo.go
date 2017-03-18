package buildinfo

import (
	"fmt"
	"strconv"
	"time"
)

var (
	name       = "{No name provided}"
	githash    = "{No githash provided}"
	buildstamp = "{No buildstamp provided}"
	builder    = "{No builder provided}"
)

func Name() string {
	return name
}

func Githash() string {
	return githash
}

func Buildstamp() string {
	return buildstamp
}

func Builder() string {
	return builder
}

func Age() string {
	i, err := strconv.ParseInt(buildstamp, 10, 64)
	if err != nil {
		return "?"
	}
	t := time.Unix(i, 0)

	return fmt.Sprintf("%s", prettyDuration(time.Now().Sub(t)))
}
