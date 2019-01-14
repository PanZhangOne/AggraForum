package result

import (
	"forum/conf"
	"github.com/kataras/iris"
)

func Map(data map[string]interface{}) iris.Map {
	var resultMap = iris.Map{}

	for key, value := range data {
		if key == "Title" {
			value = value.(string) + " - " + conf.SystemConfig.AppName
		}
		resultMap[key] = value
	}
	return resultMap
}
