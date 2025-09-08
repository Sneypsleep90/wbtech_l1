package main

import (
	"context"
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

var stopFlag int32

func main() {

	//1.Метод с сигналом
	//fmt.Println("Начало показа метода остановки с помощью флага сигнала")
	//go stopByFlag()
	//time.Sleep(5 * time.Second)
	//atomic.StoreInt32(&stopFlag, 1)
	//time.Sleep(2 * time.Second)
	//fmt.Println("Завершение показа метода с флагом!!!")

	//2.Остановка горутины с помощью отправки сигнала в канал
	//fmt.Println("Начало показа метода работы горутины с отправкой сигнала в канал")
	//ch := make(chan bool)
	//go stopByChannel(ch)
	//time.Sleep(5 * time.Second)
	//ch <- true
	//time.Sleep(2 * time.Second)
	//fmt.Println("остановка показа метода с отправкой в канал!!!")

	//3. Метод остановка с помощью контекста
	//ctx, cancel := context.WithCancel(context.Background())
	//fmt.Println("Начало показа метода остановки с помощью context!!!")
	//go stopByContext(ctx)
	//time.Sleep(5 * time.Second)
	//cancel()
	//time.Sleep(2 * time.Second)
	//fmt.Println("Оставнока показа метода с конекстом!!!")

	//4.runtime метод
	//fmt.Println("начало метода с рантайм")
	//go stopByGoExit()
	//time.Sleep(5 * time.Second)

	//5. метод стоп с помощью таймера
	fmt.Println("начало метода остановки таймер")
	go stopByTimer(5 * time.Second)
	time.Sleep(50 * time.Second)
	fmt.Println("конец")
}

func stopByFlag() {
	var counter int
	fmt.Println("Горутина начала работу!!!")
	for {
		counter++
		if atomic.LoadInt32(&stopFlag) == 1 {
			fmt.Println("Горутина остновилась!Получен стоп сигнал!")
			break
		}
		fmt.Printf("Горутина %d работает!\n", counter)
		time.Sleep(534 * time.Millisecond)
	}
}

func stopByChannel(done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Остановка горутины!")
			return
		default:
			fmt.Println("Горутина воркает!!")
			time.Sleep(545 * time.Millisecond)
		}
	}

}
func stopByContext(ctx context.Context) {
	fmt.Println("копим деньги на самсу!!")
	var moneyFoSamsa int = 10
	for {
		moneyFoSamsa++
		select {

		case <-ctx.Done():
			fmt.Printf("Закончил копить деньги на самсу, итоговая сумма: %d\n", moneyFoSamsa)
			return
		default:
			fmt.Printf("Копим деньги: %d рублей!\n", moneyFoSamsa)
			time.Sleep(645 * time.Millisecond)
		}
	}

}
func stopByGoExit() {
	var dollar int = 60
	dollar++
	fmt.Printf("Цена доллара: %d\n", dollar)
	time.Sleep(650 * time.Millisecond)
	defer fmt.Println("завершаем работу горутины")
	runtime.Goexit()
}
func stopByTimer(timeout time.Duration) {
	var bitcoin int = 1000
	timer := time.After(timeout)
	fmt.Printf("начало роста биткоина,начальная цена: %d\n", bitcoin)
	for {
		select {
		case <-timer:
			fmt.Printf("закончился рост биткоина,итоговая цена: %d\n", bitcoin)
			return
		default:
			bitcoin = bitcoin + 100
			fmt.Printf("биткоин: %d\n", bitcoin)
			time.Sleep(545 * time.Millisecond)
		}
	}
}
