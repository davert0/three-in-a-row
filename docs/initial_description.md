> Составьте предварительное, самое общее, словесное описание системы, которую вы хотите сделать.
Придерживаясь принципов модульности, определите в ней 5-7 наиболее общих сущностей (потенциальных АТД),
пока кратко и неформально их опишите.

### Игра
- Центральная сущность, инициализирующая остальные, внутри происходит главный цикл

### Поле
-  Игрвое поле, включающее в себя набор клеток, имеет  размеры

### Ячейка
- Клетка на игровом поле, может содержать в себе игровой элемент

### Игровой элемент
- Простое содержимое ячейки, имеет тип (цвет, буквенное обозначение, etc)

### Игрок
- Актор, также имеет лог своих действий

### Правила
- Объект инкапсулирующий в себе правила игры (условия при котором собирается ряд, условия начисление бонусов)

### Исполнитель
- Объект, отвечающий за исполнение игры - анализ ситуации на поле и применение действий согласно правилам   


