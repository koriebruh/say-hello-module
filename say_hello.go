package module_say_hello

func SayHello(from string) string {
	return from + " : hello geiss"
}

func SayHelloTo(from string, target string) string {
	return from + ": Hello bg " + target
}
