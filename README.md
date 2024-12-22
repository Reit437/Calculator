Калькулятор(на go)

------Описание и функция

Этот проект представляет собой простой API калькулятора, написанный на Go. API позволяет выполнять базовые арифметические операции над математическими выражениями.
`/api/v1/calculate`: Принимает POST-запрос с JSON-данными, содержащими математическое выражение. Возвращает JSON-ответ с результатом вычисления или сообщением об ошибке.

------Инструкция по запуску

1. Запустить файл main.go(Важно! Вы должны быть в директории ./Calculator)
2. Ввести в терминал(cmd) команду с вашим выражением: curl -X POST -H "Content-Type: application/json" -d "{\"expression\":\"YOUR EXPRESSION\"}" http://localhost:8080/api/v1/calculate (Важно! Знак "\" после вашего выражения убирать нельзя)
3. Просмотреть ответ и код ошибки, если она есть

ГОТОВО!

Вот эти две команды:

go run ./main.go

curl -X POST -H "Content-Type: application/json" -d "{\"expression\":\"YOUR EXPRESSION\"}" http://localhost:8080/api/v1/calculate

Для запуска тестов:
1. Вы должны находиться в директории ./Calculator/pkg/calc
2. Ввести в терминал команду: go test -v ./Calc_test.go

------Область работы

Калькулятор:
Решает любые выражения с операндами(1+1,1-1,1*2,70/35+1)

Решает любые выражения с одной и несколькими скобками((1+1)*8,(10/5)*(4+5))

Всегда поддерживает порядок действий(2+2*2)

Определяет и выводит ошибки(1+f3,1*(1+1+))

------Примеры

C:\Program Files\Go\projects\Calculator>curl -X POST -H "Content-Type: application/json" -d "{\"expression\":\"1-1\"}" http://localhost:8080/api/v1/calculate          
{"result":0}

C:\Program Files\Go\projects\Calculator>curl -X POST -H "Content-Type: application/json" -d "{\"expression\":\"(10/5)*(4+5)\"}" http://localhost:8080/api/v1/calculate          
{"result":18}

C:\Program Files\Go\projects\Calculator>curl -X POST -H "Content-Type: application/json" -d "{\"expression\":\"2+2*2\"}" http://localhost:8080/api/v1/calculate          
{"result":6}

C:\Program Files\Go\projects\Calculator>curl -X POST -H "Content-Type: application/json" -d "{\"expression\":\"1+f3\"}" http://localhost:8080/api/v1/calculate          
{"error":"Expression is not valid"}

C:\Program Files\Go\projects\Calculator>curl -X POST -H "Content-Type: application/json" -d "{\"expression\":\"1*(1+1+)\"}" http://localhost:8080/api/v1/calculate          
{"error":"error in writing, incorrect symbol after/before sign"}
