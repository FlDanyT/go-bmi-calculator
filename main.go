package main // Точка входа в приложение

import (
	"fmt" 
	"math"
	"errors"
)

func main() {

	fmt.Println("__ Калькулятор индекса массы тела __") // Println = \n
	
	for {

	var userHeight float64
	var userKg float64

	fmt.Print("Введите свой рост в сантиметрах: ")

	fmt.Scan(&userHeight) // Получение данных их консоли

	fmt.Print("Введите сврй вес: ")
	fmt.Scan(&userKg)

	IMT, err := calculateIMT(userKg, userHeight)

	if err != nil {
		fmt.Println("Не заданы параметры для расчета")
		continue

		// panic("Не заданы параметры для расчета") // Ввывод ошибки
	}

	outputResult(IMT) // Вызов функции

	isRepeateCalculation := checkRepeatCalculation()

	if !isRepeateCalculation {
			break
	}

	}

}


func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш иднекс массы тела: %.0f", imt) // %.0f - Точность, Sprint - запись в переменую
	fmt.Println(result)

	switch {

	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальная масса тела")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	
	}

	// if IMT < 16 {
	// 	fmt.Println("У вас сильный дефицит массы тела")
	// } else if IMT < 18.5 {
	// 		fmt.Println("У вас дефицит массы тела")
	// } else if IMT < 25 {
	// 		fmt.Println("У вас нормальная масса тела")
	// } else if IMT < 30 {
	// 		fmt.Println("У вас избыточный вес")
	// } else {
	// 	fmt.Println("У вас степень ожирения")
	// }

}

func calculateIMT(userKg float64, userHeight float64) (float64, error){ 

	if userKg <= 0 || userHeight <=0 { // Оброботка ошибки если неверные данные
			return 0, errors.New("NO_PARAMS_ERROR")
	}

	const IMTPower = 2

	IMT := userKg / math.Pow(userHeight / 100, IMTPower)
	return IMT, nil
}

func checkRepeatCalculation() bool {
	fmt.Print("Вы хотите сделать еще расчет: (y/n)\n")


	var userChoise string
	fmt.Scan(&userChoise)

	if userChoise == "y" || userChoise == "Y" {
		return true
	}

	return false
}