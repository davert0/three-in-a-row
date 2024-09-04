package main

import (
	"three_in_row/internal/field"
	"three_in_row/internal/game"
	"three_in_row/internal/renderer"
	"three_in_row/internal/rules"
)

//var rootCmd = &cobra.Command{
//	Use:   "game",
//	Short: "A simple CLI game",
//	Run:   runGame,
//}

func main() {
	// я запускаю игру
	// производится инициализация ресурсов
	// выводится приветственное сообщение
	// ожидается ввод пользователя

	rules_ := rules.NewRules()
	engine_ := game.NewGameEngine(rules_)
	renderer_ := renderer.NewConsoleRenderer()
	field_ := field.NewField(8, 8)
	newGame := game.NewGame(engine_, field_, renderer_)

	err := newGame.Run()
	if err != nil {
		panic(err)
	}

	//if err := rootCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
}

//func runGame(cmd *cobra.Command, args []string) {
//	fmt.Println("Добро пожаловать в игру!")
//	fmt.Println("Введите 'start' для начала игры или 'exit' для выхода.")
//
//	reader := bufio.NewReader(os.Stdin)
//
//	// Обработка Ctrl+C
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
//	go func() {
//		<-c
//		fmt.Println("\nИгра завершена.")
//		os.Exit(0)
//	}()
//
//	for {
//		fmt.Print("> ")
//		input, _ := reader.ReadString('\n')
//		input = strings.TrimSpace(input)
//
//		switch input {
//		case "start":
//			playGame(reader)
//		case "exit":
//			fmt.Println("До свидания!")
//			return
//		default:
//			fmt.Println("Неизвестная команда. Попробуйте 'start' или 'exit'.")
//		}
//	}
//}
//
//func playGame(reader *bufio.Reader) {
//	fmt.Println("Игра началась!")
//	fmt.Println("Введите координаты в формате 'x y' или 'exit' для завершения игры.")
//
//	for {
//		fmt.Print("Координаты > ")
//		input, _ := reader.ReadString('\n')
//		input = strings.TrimSpace(input)
//
//		if input == "exit" {
//			fmt.Println("Игра завершена.")
//			return
//		}
//
//		// Здесь можно добавить логику обработки координат
//		fmt.Printf("Получены координаты: %s\n", input)
//	}
//}
