
### Какой самый эффективный способ конкатенации строк?
Ответ: Использование strings.Builder. Он использует динамический буфер для построения строки, что позволяет избежать создания большого количества промежуточных строк в памяти, как это происходит при использовании оператора +
___
### Что такое интерфейсы, как они применяются в Go?
Ответ: интерфейс — это некий контракт, согласно которому компоненты системы ожидают друг от друга определенного поведения, например в части обмена информацией
___
### Чем отличаются RWMutex от Mutex?
Ответ: Mutex предназначен для блокировки доступа к общим данным для одной горутины в любое время, а RWMutex позволяет множеству горутин читать данные одновременно, но блокирует доступ на запись для других горутин.
___
### Чем отличаются буферизированные и не буферизированные каналы?
Ответ: Буферизированные имеют фиксированный размер буфера, что позволяет отправлять данные без блокировки до заполнения буфера. Небуферизированные каналы блокируются при отправке, пока данные не будут приняты.
___
### Какой размер у структуры struct{}{}?
Ответ: Пустая структура не занимает места в памяти 
___
### Есть ли в Go перегрузка методов или операторов?
Ответ: В Go нет поддержки перегрузки методов или операторов
___
### В какой последовательности будут выведены элементы map[int]int?
```go
m[0]=1
m[1]=124
m[2]=281
```
Ответ: Порядок не гарантирован.
___
### В чем разница make и new?
Ответ: new используется для выделения памяти и возвращает указатель на ноль заданного типа данных.
make используется для инициализации слайсов, карт и каналов, возвращая их инициализированные значения.
___
### Сколько существует способов задать переменную типа slice или map?
Ответ: Литерал, make, var.
___
### Что выведет данная программа и почему?
```go
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```
Ответ: 1 и 1. Потому что функция update не изменяет значение переменной, на которую указывает указатель p
___
### Что выведет данная программа и почему?
```go
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```
Ответ: Числа от 0 до 4, затем мы получим deadlock, так как надо передать wg по указателю.
___
### Что выведет данная программа и почему?
```go
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
Ответ: Выведется 0, потому что локальная область видимости if'a никак не влияет на переменную в области видимости функции.
___
### Что выведет данная программа и почему?
```go
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
Ответ: В функции someAction, значение элемента v[0] изменяется на 100.
Затем выполняется операция append, которая добавляет новый элемент b в слайс v. Однако, это изменяет копию v, а не оригинальный слайс, переданный из функции main.
Поэтому изменение внутри функции применяется только к оригинальному слайсу, но добавление нового элемента никуда не записывается. Вывод: [100 2 3 4 5]
___
### Что выведет данная программа и почему?
```go
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
Ответ: В данном случае вывод будет `[b b a][a a]`, так как после вызова append создаётся новый слайс и исходный не изменяется
___