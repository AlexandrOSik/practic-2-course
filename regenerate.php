<?php

    $COOKIE_SECRET = hash("md5", random_bytes(100));
	$FLAG = hash("md5", random_bytes(100));
	$FLAG_CHECK = hash("md5", random_bytes(100));

    $PAGE_COUNT = 10000;

    $names = [];
    for ($i = 0; $i < $PAGE_COUNT; $i++) {
        $names[] = hash('md5', random_bytes(256));
    }

    $titles_and_places = [
        ["Кузница", "в кузнице"],
        ["Поляна", "на поляну"],
        ["Лес", "в лес"],
        ["Дупло", "к чьему-то дуплу"],
        ["Озеро", "к озеру"],
        ["Дно", "на дно"],
        ["Корабль", "к кораблю"],
        ["Площадь", "на площадь"],
        ["Дом", "в дом"],
        ["Подвал", "в подвал"],
    ];
    $tellers = [
        "Кот", "Дед Мороз", "Пчела", 
        "Медведь", "Единорог", "Гриб", 
        "Кролик", "Капибара", "Капибарёнок",
        "Паладин 80 уровня", "Шляпа волшебника", "Абсолютный ноль",
        "Маг", "Пингвин", "Целитель",
        "Заяц по имени Волк", "Котокошка", "Бедствие",
        "Торговец", "Ваш двойник", "Некто",
        "Тот самый Он", "Корова", "Принцесса",
        "Пельмень Ли", "Айболит", "Бревно",
        "Полено", "Зимбобо"
    ];
    $words = [
        "У меня теперь есть новая тройка! Теперь могу путешествовать по московскому метро!", 
        "Что у тебя для меня сегодня есть?", 
        "Не смотри, я стесняюсь!", 
        "Зачем я здесь?",
        "Продаю мёд! Недорого!",
        "Эй, псс... Есть что?",
        "Вспомнишь про... А... Это всего лишь ты...",
        "Вы потерялись? Вас найти?",
        "Не хочешь есть? Вот и я хочу.",
        "Не видел тут говорящие Бревно или Полено? Я их потерял...",
    ];
    $items = [
        "Палка-копалка",
        "Ключ к сердцу",
        "Меч леденец",
        "Перчатка резиновая 1 штука",
        "Огурец",
        "Банан",
        "Сеньор Помидор",
        "Баклажан",
        "Лада Приора",
        "Палка-стрелялка",
        "Плеть",
        "Безгаечный ключ",
        "Пылесос",
        "Вареник",
    ];

    $PAGES = [];
    $title_indexes = [];
    for ($i = 0; $i < count($titles_and_places); $i++) {
        $title_count[] = 0;
    }

    for ($i = 0; $i < $PAGE_COUNT; $i++) {

        $title_index = array_rand($titles_and_places);
        $title_count[$title_index] += 1;
        $title_and_place = $titles_and_places[$title_index];
        
        $title = $title_and_place[0] . " " . $title_count[$title_index];
        $place = $title_and_place[1];
        $name = $names[$i];
        $teller = $tellers[array_rand($tellers)];
        $word = $words[array_rand($words)];
        $transitions = array_rand($names, random_int(2, 5));
        $tree_child_cnt = 2;
        for ($j = 0; $j < $tree_child_cnt; ++$j) {
            $node = $i * $tree_child_cnt + $j;
            if ($node < $PAGE_COUNT && !in_array($node, $transitions)) {
                $transitions[] = $node;
            }
        }
        $transitions = implode(', ', $transitions);

        $PAGES[] = "\n\t\t{\n" .
			"\t\t\tName:        \"$name\",\n" .
			"\t\t\tTransitions: []int{{$transitions}},\n" .
			"\t\t\tTitle:       \"$title\",\n" .
			"\t\t\tPlace:       \"$place\",\n" .
			"\t\t\tTeller:      \"$teller\",\n" .
			"\t\t\tWords:       \"$word\",\n" .
        "\t\t},";
    }

    $KEY_PAGES = array_rand($PAGES, 3);
    $KEY_PAGE_ITEMS = array_rand($items, count($KEY_PAGES));
    $KEY_PAGE_CHECK = [];
    $KEY_PAGE_SYMBOL = array_slice(
        ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'], 
        0, count($KEY_PAGES));
    for ($i = 0; $i < count($KEY_PAGES); $i++) {
        $KEY_PAGE_ITEMS[$i] = $items[$KEY_PAGE_ITEMS[$i]];
        $KEY_PAGE_CHECK[] = hash('md5', random_bytes(100));
    }

	$FLAG_PAGE = hash('md5', random_bytes(10));
	$FLAG_PAGE_GATE = array_rand($PAGES);
	$ACT_PAGE = hash('md5', random_bytes(10));
    
    $PAGES = implode("", $PAGES);
    $KEY_PAGES = implode(", ", $KEY_PAGES);
    $KEY_PAGE_ITEMS = '"' . implode("\", \"", $KEY_PAGE_ITEMS) . '"';
    $KEY_PAGE_CHECK = '"' . implode("\", \"", $KEY_PAGE_CHECK) . '"';
    $KEY_PAGE_SYMBOL = '"' . implode("\", \"", $KEY_PAGE_SYMBOL) . '"';
?>

package utils

type PageNode struct {
	Name        string
	Transitions []int
	Title       string
	Place       string
	Teller      string
	Words       string
}

type Constants struct {
	COOKIE_SECRET   string
	FLAG            string
	FLAG_CHECK      string
	PAGE_COUNT      int
	PAGES           []PageNode
	PAGE_INDEXES    map[string]int
	KEY_PAGES       []int
	KEY_PAGE_ITEMS  []string
	KEY_PAGE_SYMBOL []string
	KEY_PAGE_CHECK  []string
	FLAG_PAGE       string
	FLAG_PAGE_GATE  int
	ACT_PAGE        string
}

var Const Constants = Constants{}

func init() {
	Const.COOKIE_SECRET = "<?= $COOKIE_SECRET ?>"
	Const.FLAG = "<?= $FLAG ?>"
	Const.FLAG_CHECK = "<?= $FLAG_CHECK ?>"
	Const.PAGES = []PageNode{
<?= $PAGES ?>
	}
	Const.KEY_PAGES = []int{<?= $KEY_PAGES ?>}
	Const.KEY_PAGE_ITEMS = []string{<?= $KEY_PAGE_ITEMS ?>}
	Const.KEY_PAGE_CHECK = []string{<?= $KEY_PAGE_CHECK ?>}
	Const.KEY_PAGE_SYMBOL = []string{<?= $KEY_PAGE_SYMBOL ?>}
	Const.FLAG_PAGE = "<?= $FLAG_PAGE ?>"
	Const.FLAG_PAGE_GATE = <?= $FLAG_PAGE_GATE ?> 
	Const.ACT_PAGE = "<?= $ACT_PAGE ?>"
	Const.PAGE_COUNT = len(Const.PAGES)
	Const.PAGE_INDEXES = make(map[string]int)
	for i := 0; i < Const.PAGE_COUNT; i++ {
		Const.PAGE_INDEXES[Const.PAGES[i].Name] = i
	}

}
