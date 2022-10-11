package sequence

//Hold static ansi special character sequences
const (
	//Bright eight colors set
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Megenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	//Bright eight color set
	BBlack   = "\033[30;1m"
	BRed     = "\033[31;1m"
	BGreen   = "\033[32;1m"
	BYellow  = "\033[33;1m"
	BBlue    = "\033[34;1m"
	BMegenta = "\033[35;1m"
	BCyan    = "\033[36;1m"
	BWhite   = "\033[37;1m"
	Reset    = "\033[0m"

	//Background colors
	BkgBlack   = "\033[40m"
	BkgRed     = "\033[41m"
	BkgGreen   = "\033[42m"
	BkgYellow  = "\033[43m"
	BkgBlue    = "\033[44m"
	BkgMegenta = "\033[45m"
	BkgCyan    = "\033[46m"
	BkgWhite   = "\033[47m"

	//Bright background colors
	BBkgBlack   = "\033[40;1m"
	BBkgRed     = "\033[41;1m"
	BBkgGreen   = "\033[42;1m"
	BBkgYellow  = "\033[43;1m"
	BBkgBlue    = "\033[44;1m"
	BBkgMegenta = "\033[45;1m"
	BBkgCyan    = "\033[46;1m"
	BBkgWhite   = "\033[47;1m"

	//Text decorations
	Bold    = "\033[1m"
	ULine   = "\033[4m"
	Reverse = "\033[7m"

	//Clear codes
	CLEntire = "\033[2K"
	CLStart  = "\033[1K"
	CLEnd    = "\033[0K"
)
