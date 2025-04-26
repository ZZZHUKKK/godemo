package output

import (
	"github.com/fatih/color"
)

func Output(some any) { //func Output(some intrface{}) {}
	// intValue, ok := some.(int)
	// if ok {
	// 	color.Red("Код ошибибки: %d", intValue)
	// 	return
	// }
	// strValue, ok := some.(string)
	// if ok {
	// 	color.Red(strValue)
	// 	return
	// }
	// errorValue, ok := some.(error)
	// if ok {
	// 	color.Red(errorValue.Error())
	// 	return
	// }

	switch t := some.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("Неизвестный тип ошибки")
	}
}
