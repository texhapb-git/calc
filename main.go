package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Калькулятор (AVG/SUM/MED)")

	
	for {
		// Первый этап: ввод операции
		fmt.Print("Введите операцию (AVG/SUM/MED): ")
		var operation string
		fmt.Scan(&operation)
		
		operation = strings.TrimSpace(operation)
		operation = strings.ToUpper(operation)

		if operation != "AVG" && operation != "SUM" && operation != "MED" {
			fmt.Println("Ошибка: неподдерживаемая операция. Используйте AVG, SUM или MED")
			continue
		}
		
		// Второй этап: ввод чисел
		fmt.Print("Введите числа через запятую: ")
		var numbersStr string
		fmt.Scan(&numbersStr)
		
		numbersStr = strings.TrimSpace(numbersStr)
		if numbersStr == "exit" {
			break
		}
		
		numbers, err := parseNumbers(numbersStr)
		if err != nil {
			fmt.Printf("Ошибка парсинга чисел: %v\n", err)
			continue
		}
		
		if len(numbers) == 0 {
			fmt.Println("Ошибка: не указаны числа")
			continue
		}
		
		result := calculate(operation, numbers)
		
		fmt.Printf("Результат: %.2f\n", result)
		
		// Запрос на продолжение
		fmt.Println("\nНажмите Enter для новой конвертации или Ctrl+C для выхода...")
		var dummy string
		fmt.Scanln(&dummy) // Ожидание Enter
	}
}

func parseNumbers(input string) ([]float64, error) {
	// Убираем все пробелы и разбиваем по запятым
	cleanInput := strings.ReplaceAll(input, " ", "")
	parts := strings.Split(cleanInput, ",")
	
	var numbers []float64
	for _, part := range parts {
		if part == "" {
			continue
		}
		
		num, err := strconv.ParseFloat(strings.TrimSpace(part), 64)
		if err != nil {
			return nil, fmt.Errorf("неверное число: %s", part)
		}
		numbers = append(numbers, num)
	}
	
	return numbers, nil
}

func calculate(operation string, numbers []float64) float64 {
	switch operation {
	case "SUM":
		return sum(numbers)
	case "AVG":
		return avg(numbers)
	case "MED":
		return median(numbers)
	default:
		return 0
	}
}

func sum(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total
}

func avg(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return sum(numbers) / float64(len(numbers))
}

func median(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	
	// Создаем копию для сортировки
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)
	
	n := len(sorted)
	if n%2 == 0 {
		// Четное количество элементов - среднее двух средних
		return (sorted[n/2-1] + sorted[n/2]) / 2
	} else {
		// Нечетное количество элементов - средний элемент
		return sorted[n/2]
	}
}