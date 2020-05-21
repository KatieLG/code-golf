package hole

import (
	"math/rand"
	"strings"
)

func brainfuck() ([]string, string) {
	args := []string{
		"+++++++++++++++[>++>+++>++++>+++++>++++++>+++++++>++++++++<<<<<<<-]+++++++++++++++>>>++++++.>>+++++++.>>-----.<-.<<<<<<-----.",
		"+++++++++++++++[>++>+++>++++>+++++>++++++>+++++++>++++++++<<<<<<<-]+++++++++++++++>>>>-.>+++++++.>>--.<<.<+++++++++.>++.>>----.<.>--.++++.<<<<<<<-----.",
		"+++++++++++++++++++++++++[>++>+++>++++>+++++<<<<-]+++++++++++++++++++++++++>>+.>>--------.<---.<<<---------------.",
		"+++++++++++++++++++++++++[>++>+++>++++>+++++<<<<-]+++++++++++++++++++++++++>>+++++.>+.>-----------.------.<<<<---------------.",
		"+++++++++++++++[>++>+++>++++>+++++>++++++>+++++++>++++++++<<<<<<<-]+++++++++++++++>>>>+++++.>>----.>------.------.<<<<<<++.>>------.<<<-----.",
		"+++++++++++++++++++++[>++>+++>++++>+++++>++++++<<<<<-]+++++++++++++++++++++>>>----.--------.++++++++.<<<-----------.",
		"+++++++++++++++++++++[>++>+++>++++>+++++>++++++<<<<<-]+++++++++++++++++++++>>>----.>>-----.-----.<-.>-----.-.<<<<<-----------.",
		"+++++++++++++++++++++[>++>+++>++++>+++++>++++++<<<<<-]+++++++++++++++++++++>>>--.>>---------.<-------.>++++.<<<<<-----------.",
		"++++++++++++++++++[>++>+++>++++>+++++>++++++>+++++++<<<<<<-]++++++++++++++++++>>>-----.>>+++.<++++++++++.+.<<<----.>>++++.>>.---.<+.<<<<--------.",
		">>>>>>>>>>>>>>>>>>>>>>>>>>>>++++++++++++++++++++++++++[-<<[+<]+[>]>][<<[[-]-----<]>[>]>]<<[++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++<]>[.>]++++++++++.",
		"+++++[>+++++[>++>++>+++>+++>++++>++++<<<<<<-]<-]+++++[>>[>]<[+.<<]>[++.>>>]<[+.<]>[-.>>]<[-.<<<]>[.>]<[+.<]<-]++++++++++.",
	}

	outs := []string{
		"Bash",
		"JavaScript",
		"Lua",
		"Perl",
		"Perl 6",
		"PHP",
		"Python",
		"Ruby",
		"Code Golf",
		"abcdefghijklmnopqrstuvwxyz",
		"eL34NfeOL454KdeJ44JOdefePK55gQ67ShfTL787KegJ77JTeghfUK88iV9:XjgYL:;:KfiJ::JYfijgZK;;k[<=]lh^L=>=KgkJ==J^gklh_K>>m`?@bnicL@A@KhmJ@@JchmnidKAA",
	}

	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
		outs[i], outs[j] = outs[j], outs[i]
	})

	return args, strings.Join(outs, "\n")
}
