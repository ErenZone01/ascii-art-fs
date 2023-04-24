go run . "" | cat -e
go run . "\n" | cat -e
go run . "hello" standard | cat -e
go run . "HELLO" thinkertoy | cat -e
go run . "HeLlo HuMaN" shadow| cat -e
go run . "1Hello 2There" thinkertoy | cat -e
go run . "Hello\nThere" shadow | cat -e
go run . "Hello\n\nThere!!!!" thinkertoy | cat -e
go run . "{Hello & There #}" shadow | cat -e
go run . 'hello There 1 to 2!' | cat -e
go run . "MaD3IrA&LiSboN" | cat -e
go run . "1a\"#FdwHywR&/()=" thinkertoy | cat -e
go run . "{|}~" shadow | cat -e
go run . "[\]^_ 'a" | cat -e
go run . "RGB" thinkertoy | cat -e
go run . ":;<=>?@" | cat -e
go run . '\!" #$%&'"'"'()*+,-./' | cat -e
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" thinkertoy | cat -e
go run . "abcdefghijklmnopqrstuvwxyz" shadow | cat -e
go run . "abcdeWXC" | cat -e
go run . "abcde 12" shadow | cat -e
go run . "A@&!" | cat -e
go run . "ab  2!%AZE" thinkertoy | cat -e