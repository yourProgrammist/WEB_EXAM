package data

import (
	"math/rand"
	"time"
)

/*type Route struct {
	Id          int64
	Name        string
	Description string
	MainObject  string
	Coords      [2][2]float64

}*/

type Routes struct {
	Arr   [][]string
	Guids [][]string
}

type Location struct {
	Start string
	Stop  string
}

type Orders struct {
	Arr [][]string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (r *Routes) GetAll() {
	rand.Seed(time.Now().UnixNano())
	r.Arr = [][]string{}
	r.Arr = append(r.Arr, []string{
		"1",
		"Арбат - душа Москвы",
		"Прогулка новому Арбату",
		"Пешая прогулка по новому Арбату",
		"​Улица Воздвиженка, 5/25 ст1​",
		"Смоленская-Сенная площадь",
	})
	r.Arr = append(r.Arr, []string{
		"2",
		"Воробьёвы горы",
		"Пешая прогулка по Воробьёвым горам",
		"Воробьёвы горы - одно из красивейших мест в Москве. С высокого берега открывается изумительный вид на столицу, а также главное здание МГУ.",
		"Москва, Улица Лужники, 24ст1",
		"Москва, ​Ленинские Горы, 1",
	})
	r.Arr = append(r.Arr, []string{
		"3",
		"Нескучный сад",
		"Прогулка от саду",
		"Нескучный сад – старейший парк Москвы. В середине XVIII века здесь, в своём имении, известный московский миллионер Прокопий Демидов построил дворец, получивший название «Демидовский».",
		"Москва, ​Улица Крымский Вал, 9",
		"Москва, Нескучный сад",
	})
	r.Arr = append(r.Arr, []string{
		"4",
		"Для всё семьи: планетарий, зоопарк и музеи",
		"Поход в музеи и зоопарк",
		"Познавательный маршрут для исследователей живой природы в центре Москвы.Три музея, зоопарк и планетарий",
		"Московский зоопарк, Б.Грузинская, 1",
		"Даниловский рынок, Мытная, 74",
	})
	r.Arr = append(r.Arr, []string{
		"5",
		"Исторический маршрут для всей семьи",
		"Пешая прогулка по историческим местам",
		"Маршрут начинается со спокойной прогулки по исторической Усадьбе Кусково, из которой предлагается, передвигаясь по синей ветке, добраться до музеев в самом центре Москвы.",
		"​Москва, ​Усадьба Кусково",
		"Москва, Кремль",
	})
	r.Arr = append(r.Arr, []string{
		"6",
		"Главные пешеходные улицы",
		"Пешая прогулка по главным пешеходным улицам и переулкам столицы",
		"Маршрут идеален именно для вечерней прогулки, когда на этих улицах загорается всё их световое оформление. Проходит он по улице Никольской, Кузнецкому мосту, Столешникову переулку, Камергерскому переулку и улице Большая Дмитровка.",
		"Москва, метро Лубянка",
		"Москва, метро Чеховская",
	})
	r.Arr = append(r.Arr, []string{
		"7",
		"Ахматовские места в Москве",
		"Прогулка по ахматовским местам",
		"Прогуляйтесь по Москве поэтессы и посмотрите, как преображается город в ее стихах.",
		"Москва, ​Улица Большая Ордынка, 31/12 ст1",
		"Москва, Переулок Померанцев, 3",
	})
	r.Arr = append(r.Arr, []string{
		"8",
		"Авторский маршрут Королькова А.Ю. «Правда о Соколе»",
		"Прогулка по Соколу",
		"Обзорная экскурсия по району Сокол с подробными рассказами о храме Всех Святых во Всехсвятском, кооперативном посёлке «Сокол» и Мемориально-парковом комплексе героев Первой мировой войны. Маршрут разработан и подготовлен Корольковым Андреем Юрьевичем.",
		"Москва, Станция метро «Сокол»",
		"Москва, метро Войковская",
	})

	r.Guids = [][]string{}
	r.Guids = append(r.Guids, []string{
		"Игнатов Марат Игоревич",
		"Русский, Английский",
		"7 лет",
		"1000",
		"Арбат - душа Москвы",
	})
	r.Guids = append(r.Guids, []string{
		"Бурунов Сергей Маратович",
		"Русский",
		"10 лет",
		"1500",
		"Исторический маршрут для всей семьи",
	})
	r.Guids = append(r.Guids, []string{
		"Осин Кирилл Андреевич",
		"Испанский",
		"12 лет",
		"2500",
		"Ахматовские места в Москве",
	})	
	r.Guids = append(r.Guids, []string{
		"Репчанский Николай Евгеньевич",
		"Русский",
		"7 лет",
		"1250",
		"Нескучный сад",
	})
	r.Guids = append(r.Guids, []string{
		"Сазонов Иван Максимович",
		"Немецкий, Арабский, Русский",
		"17 лет",
		"3200",
		"Авторский маршрут Королькова А.Ю. «Правда о Соколе»",
	})
}