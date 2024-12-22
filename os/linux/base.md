# 结构

- /
  - var : 存放可变数据文件(动态)
    - lib: 存放与特定应用程序相关的状态信息和数据文件
    - log : 存放系统和应用程序的日志文件
  - etc : 存放系统的配置文件(静态)
    - os-release : 存储操作系统的识别信息(key-value)
    - shadow : 存储用户密码和其他与用户帐户相关的安全信息的重要文件
    - fstab : `file system table`包含系统启动时或执行 mount -a 命令时需要挂载的文件系统的静态信息 
  - opt : 存放可选的第三方软件包
  - tmp : 存储临时文件和数据(temporary)
  - dev : 表示设备文件
    - /null : 任何写入 `/dev/null` 的数据都会被丢弃

# 命令

## 目录&文件操作

### journalctl

```bash
journalctl [OPTIONS...] [MATCHES...]
# 访问和查看由 systemd 管理的系统日志，包含系统启动信息、服务日志、内核消息、应用程序日志等
```

- options : 

  - -S --since=DATE : 显示date之后的日志

  - -U --until=DATE : 显示date之前的日志

  - -u --unit=UNIT : 查看特定服务的日志

  - -f : 实时查看日志

  - xe: 查看当前登录会话的日志

  - -p : 按日志的优先级（日志级别）进行过滤

    ```bash
    0 - emergency（紧急）
    1 - alert（警报）
    2 - critical（严重）
    3 - error（错误）
    4 - warning（警告）
    5 - notice（通知）
    6 - info（信息
    7 - debug（调试）
    ```

    

### awk

```bash

```

```bash
awk [options...][program][file ...]
 
options:
  -F <fs>						fs可以是单个字符、字符串或正则表达式，默认为空格
  -v <var=value>		awk程序定义变量并赋给指定的值，可以在程序开始之前设置多个变量
  -f <path>					指定执行脚本

program: ‘
//{command1; command2} END{}’
  BEGIN			初始化代码块，在对每一行进行处理之前，初始化代码，主要是引用全局变量，设置FS分隔符
  //				匹配代码块，可以是字符串或正则表达式
  {}				命令代码块，包含一条或多条命令
  ;					多条命令使用分号分隔
  END				结尾代码块，在对每一行进行处理之后再执行的代码块，主要是进行最终计算或输出结尾摘要信息

特殊符
  -F '[:#/]'   			定义三个分隔符
  $0           			表示整个当前行
  $1/$n        			每行第一个字段/第n个字段
  NF					 			字段数量变量
  NR         	 			每行的记录号，多文件记录递增
  FNR        	 			与NR类似，不过多文件记录不递增，每个文件都从1开始
  \t           			制表符
  \n           			换行符
  FS           			BEGIN时定义分隔符
  RS       		 			输入的记录分隔符， 默认为换行符(即文本是按一行一行输入)
  ~            			匹配，与==相比不是精确比较
  !~           			不匹配，不精确比较
  ==         	 			等于，必须全部相等，精确比较
  !=           			不等于，精确比较
  &&							  逻辑与
  ||     			      逻辑或
  + 		            匹配时表示1个或1个以上
  /[0-9][0-9]+/   	两个或两个以上数字
  /[0-9][0-9]*/    	一个或一个以上数字
  FILENAME 					文件名
  OFS      					输出字段分隔符， 默认也是空格
  ORS       	 			输出的记录分隔符，默认为换行符

commands:
	func: 
		print 			打印
		system()		执行命令行命令

rule：
	1.当两个字符串或字符串与变量之间没有运算符时，默认会进行字符串拼接操作
```

- 选项: 可选参数，如 -F 指定字段分隔符
- 模式: 指定何时执行动作的条件
- 动作: 当模式匹配时要执行的操作
  - print: 用于打印整行或指定字段
  - $[n]: 字段变量
- 文件名: 需要处理的文件名

### find

```bash
find [path...] [expression]
-print	将当前找到的文件路径打印到标准输出。
-quit	立即终止 find 命令的执行。
```

- expression

  - -name : 指定名字文件，可以使用正则

  - -type : 指定类型

    ```bash
    d : 目录
    f :	文件
    l : 符号连接
    ```

  - size : 按大小寻找

    ```bash
    +100M : 大于100M
    -100M : 小于100M
    ```

  - -exec : 对找到的文件执行某个命令

### grep

```bash
grep [options...] "pattern" filename
```

- options :
  - -i : 忽略大小写
  - -n : 显示行号
  - -v : 反向匹配（不显示匹配的行）
  - -w : 匹配完整单词
  - -c : 显示匹配行的数量
  - -l : 显示包含匹配模式的文件名
  - -r  : 递归搜索目录
  - -E : 使用扩展正则表达式
  - -P : 使用 Perl 正则表达式。
  - -C：显示匹配行前后行数

### awk

```bash
Usage: mawk [Options] [Program] [file ...]

Program:
    The -f option value is the name of a file containing program text.
    If no -f option is given, a "--" ends option processing; the following
    parameters are the program text.

Options:
    -f program-file  Program  text is read from file instead of from the
                     command-line.  Multiple -f options are accepted.
    -F value         sets the field separator, FS, to value.
    -v var=value     assigns value to program variable var.
    --               unambiguous end of options.

    Implementation-specific options are prefixed with "-W".  They can be
    abbreviated:

    -W version       show version information and exit.
    -W compat        pre-POSIX 2001.
    -W dump          show assembler-like listing of program and exit.
    -W help          show this message and exit.
    -W interactive   set unbuffered output, line-buffered input.
    -W exec file     use file as program as well as last option.
    -W posix         stricter POSIX checking.
    -W random=number set initial random seed.
    -W sprintf=number adjust size of sprintf buffer.
    -W usage         show this message and exit.
```



### 基本正则

- .（点）: 匹配任何单个字符（除了换行符），a.c 可以匹配 "abc"、"axc" 
- `*`（星号）: 匹配前面的字符或子表达式零次或多次,，a*b 匹配 "b"、"ab"、"aab"、"aaab" 等
- ^（脱字符）: 匹配行的开头，^abc 匹配以 "abc" 开头的行
- $（美元符号）: 匹配行的结尾，abc$ 匹配以 "abc" 结尾的行
- []（方括号）: 匹配方括号内的任何一个字符，可以使用范围， [a-z] 匹配任何小写字母，[abc] 匹配 "a", "b" 或 "c"
- [^]（否定方括号）: 匹配不在方括号内的任何字符，`[^0-9] `匹配任何非数字字符
- \（反斜杠）: 用于转义特殊字符，使其失去特殊含义，`\.` 匹配一个实际的句点，而不是“任意字符”。

### cp

```
cp [OPTIONS...] source destination
```

- options :
  - -r : 递归
  - -f : 强制
  - -v : 输出过程
  - -u : 更新

### wc

```bash
wc [options] [flie...]
# word count 用于统计文件的字数、行数、字符数和字节数
```

- options
  - `-l` :  统计行数（Line count）
  - `-w : 统计单词数（Word count）
  - `-c : 统计字符数（Character count）
  - `-m` : 统计字符数（按字符计）
  - `-L` : 统计文件中最长一行的长度

### dirname

```bash
dirname [option] name...
-z, --zero		输出不换行
--help				帮助
--version			版本
```

```bash
Usage: dirname [OPTION] NAME...
Output each NAME with its last non-slash component and trailing slashes
removed; if NAME contains no /'s, output '.' (meaning the current directory).

  -z, --zero     end each output line with NUL, not newline
      --help        display this help and exit
      --version     output version information and exit

Examples:
  dirname /usr/bin/          -> "/usr"
  dirname dir1/str dir2/str  -> "dir1" followed by "dir2"
  dirname stdio.h            -> "."
  
GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/dirname>
or available locally via: info '(coreutils) dirname invocation'
```

### cut

```bash
cut [option] [file]
```

```bash
Usage: cut OPTION... [FILE]...
Print selected parts of lines from each FILE to standard output.

With no FILE, or when FILE is -, read standard input.

Mandatory arguments to long options are mandatory for short options too.
  -b, --bytes=LIST        select only these bytes
  -c, --characters=LIST   select only these characters
  -d, --delimiter=DELIM   use DELIM instead of TAB for field delimiter
  -f, --fields=LIST       select only these fields;  also print any line
                            that contains no delimiter character, unless
                            the -s option is specified
  -n                      (ignored)
      --complement        complement the set of selected bytes, characters
                            or fields
  -s, --only-delimited    do not print lines not containing delimiters
      --output-delimiter=STRING  use STRING as the output delimiter
                            the default is to use the input delimiter
  -z, --zero-terminated   line delimiter is NUL, not newline
      --help        display this help and exit
      --version     output version information and exit

Use one, and only one of -b, -c or -f.  Each LIST is made up of one
range, or many ranges separated by commas.  Selected input is written
in the same order that it is read, and is written exactly once.
Each range is one of:

  N     N'th byte, character or field, counted from 1
  N-    from N'th byte, character or field, to end of line
  N-M   from N'th to M'th (included) byte, character or field
  -M    from first to M'th (included) byte, character or field

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation <https://www.gnu.org/software/coreutils/cut>
or available locally via: info '(coreutils) cut invocation'
```



## 系统操作

### systemctl

```bash
systemctl [OPTIONS...] COMMAND ...
# 管理和控制 systemd 系统和服务管理器的命令行工具
```

- command

  - status <server-name> : 查看特定服务的当前状态
  - start <server-name> : 启动服务
  - restart <server-name> : 重启服务
  - reload <server-name> : 重新加载服务配置
  - stop <server-name> : 停止服务
  - enable <server-name> : 系统启动时自动启动
  - disable <server-name> : 禁用服务
  - daemon-reload: 重新加载 systemd 系统和服务管理器的配置文件

- options : 

  - --type : 指定要列出的单位类型（unit type）。常见的单位类型包括 service, socket, target, mount

  - --state : 过滤和显示具有特定状态的服务单元

    ```bash
    active: 单元正在运行或已完成
    inactive: 单元未运行且未激活
    reloading: 单元正在重新加载其配置
    failed: 单元失败了，可能是因为启动过程中出现错误
    activating: 单元正在启动中
    deactivating: 单元正在停止中
    ```

  - --force : 强制执行

### lsblk

```bash
lsblk
# list block devices，用于列出系统中所有可用的块设备信
```

- 

## install

```bash
install [选项] 源文件 目标文件
install [选项] 源文件... 目标目录
install -d 目标目录...
# 安装或复制文件及目录，并设置适当的权限和所有权
```

- 选项

  - -d 或 --directory: 创建目录。如果目录已存在，则不会报错

  - -m 或 --mode: 设置文件或目录的权限模式。模式可以是八进制数字或符号表示法

## curl

```bash
curl [选项] [URL]
#用于从服务器传输数据或向服务器传输数据的命令行工具
```

- 选项
  - -f: 禁止返回 HTTP 错误页面，如果请求失败，`curl` 会返回一个错误而不是显示错误页面
  - -s: 静默模式，不显示进度条或其他输出
  - -S: 即使使用 `-s` 选项，仍然在发生错误时显示错误信息
  - -L: 跟随重定向。如果服务器返回一个重定向响应，`curl` 会自动跟随重定向
  - -o: 指定下载的文件保存的路径和文件名

## chomd

```bash
chmod [选项] 权限模式 文件或目录
# 更改文件或目录权限
```

- 选项
  - -R 或 --recursive: 递归地更改目录及其内容的权限
  - -v 或 --verbose: 显示详细的输出信息
- 权限模式
  - 符号表示法
  - 八进制表示法

## tee

```bash
tee [选项] 文件...
# 将命令的输出同时显示在终端上并保存到文件中
```

- 选项
  - -a: 追加到文件末尾而不是覆盖文件

## sudo

```bash
# sudo - execute a command as another user
sudo [optins]

Options:
  -i, --login                   run login shell as the target user; a command may also be specified
```

## source

​	在当前shell环境中执行一个指定的脚本文件（和 . 同效）。

## journalctl

​	journalctl 是 systemd 系统和服务管理器的一部分，用于查询和显示系统日志。systemd 日志系统（也称为 journald）收集来自内核、系统服务和用户进程的日志信息，并将其存储在一个二进制日志文件中。journalctl 命令提供了多种选项来过滤和格式化这些日志信息，使其更易于阅读和分析。

```bash
journalctl [OPTIONS...] [MATCHES...]

Query the journal.

Flags:
     --system              Show the system journal
     --user                Show the user journal for the current user
  -M --machine=CONTAINER   Operate on local container
  -S --since=DATE          Show entries not older than the specified date
  -U --until=DATE          Show entries not newer than the specified date
  -c --cursor=CURSOR       Show entries starting at the specified cursor
     --after-cursor=CURSOR Show entries after the specified cursor
     --show-cursor         Print the cursor after all the entries
  -b --boot[=ID]           Show current boot or the specified boot
     --list-boots          Show terse information about recorded boots
  -k --dmesg               Show kernel message log from the current boot
  -u --unit=UNIT           Show logs from the specified unit
  -t --identifier=STRING   Show entries with the specified syslog identifier
  -p --priority=RANGE      Show entries with the specified priority
  -e --pager-end           Immediately jump to the end in the pager
  -f --follow              Follow the journal
  -n --lines[=INTEGER]     Number of journal entries to show
     --no-tail             Show all lines, even in follow mode
  -r --reverse             Show the newest entries first
  -o --output=STRING       Change journal output mode (short, short-iso,
                                   short-precise, short-monotonic, verbose,
                                   export, json, json-pretty, json-sse, cat)
     --utc                 Express time in Coordinated Universal Time (UTC)
  -x --catalog             Add message explanations where available
     --no-full             Ellipsize fields
  -a --all                 Show all fields, including long and unprintable
  -q --quiet               Do not show privilege warning
     --no-pager            Do not pipe output into a pager
  -m --merge               Show entries from all available journals
  -D --directory=PATH      Show journal files from directory
     --file=PATH           Show journal file
     --root=ROOT           Operate on catalog files underneath the root ROOT
     --interval=TIME       Time interval for changing the FSS sealing key
     --verify-key=KEY      Specify FSS verification key
     --force               Override of the FSS key pair with --setup-keys

Commands:
  -h --help                Show this help text
     --version             Show package version
  -F --field=FIELD         List all values that a specified field takes
     --new-id128           Generate a new 128-bit ID
     --disk-usage          Show total disk usage of all journal files
     --vacuum-size=BYTES   Reduce disk usage below specified size
     --vacuum-time=TIME    Remove journal files older than specified date
     --flush               Flush all journal data from /run into /var
     --header              Show journal header information

```

## tree

```bash
tree --help
usage: tree [-acdfghilnpqrstuvxACDFJQNSUX] [-L level [-R]] [-H  baseHREF]
        [-T title] [-o filename] [-P pattern] [-I pattern] [--gitignore]
        [--gitfile[=]file] [--matchdirs] [--metafirst] [--ignore-case]
        [--nolinks] [--hintro[=]file] [--houtro[=]file] [--inodes] [--device]
        [--sort[=]<name>] [--dirsfirst] [--filesfirst] [--filelimit #] [--si]
        [--du] [--prune] [--charset[=]X] [--timefmt[=]format] [--fromfile]
        [--fromtabfile] [--fflinks] [--info] [--infofile[=]file] [--noreport]
        [--version] [--help] [--] [directory ...]
  ------- Listing options -------
  -a            All files are listed.
  -d            List directories only.
  -l            Follow symbolic links like directories.
  -f            Print the full path prefix for each file.
  -x            Stay on current filesystem only.
  -L level      Descend only level directories deep.
  -R            Rerun tree when max dir level reached.
  -P pattern    List only those files that match the pattern given.
  -I pattern    Do not list files that match the given pattern.
  --gitignore   Filter by using .gitignore files.
  --gitfile X   Explicitly read gitignore file.
  --ignore-case Ignore case when pattern matching.
  --matchdirs   Include directory names in -P pattern matching.
  --metafirst   Print meta-data at the beginning of each line.
  --prune       Prune empty directories from the output.
  --info        Print information about files found in .info files.
  --infofile X  Explicitly read info file.
  --noreport    Turn off file/directory count at end of tree listing.
  --charset X   Use charset X for terminal/HTML and indentation line output.
  --filelimit # Do not descend dirs with more than # files in them.
  -o filename   Output to file instead of stdout.
  ------- File options -------
  -q            Print non-printable characters as '?'.
  -N            Print non-printable characters as is.
  -Q            Quote filenames with double quotes.
  -p            Print the protections for each file.
  -u            Displays file owner or UID number.
  -g            Displays file group owner or GID number.
  -s            Print the size in bytes of each file.
  -h            Print the size in a more human readable way.
  --si          Like -h, but use in SI units (powers of 1000).
  --du          Compute size of directories by their contents.
  -D            Print the date of last modification or (-c) status change.
  --timefmt <f> Print and format time according to the format <f>.
  -F            Appends '/', '=', '*', '@', '|' or '>' as per ls -F.
  --inodes      Print inode number of each file.
  --device      Print device ID number to which each file belongs.
  ------- Sorting options -------
  -v            Sort files alphanumerically by version.
  -t            Sort files by last modification time.
  -c            Sort files by last status change time.
  -U            Leave files unsorted.
  -r            Reverse the order of the sort.
  --dirsfirst   List directories before files (-U disables).
  --filesfirst  List files before directories (-U disables).
  --sort X      Select sort: name,version,size,mtime,ctime.
  ------- Graphics options -------
  -i            Don't print indentation lines.
  -A            Print ANSI lines graphic indentation lines.
  -S            Print with CP437 (console) graphics indentation lines.
  -n            Turn colorization off always (-C overrides).
  -C            Turn colorization on always.
  ------- XML/HTML/JSON options -------
  -X            Prints out an XML representation of the tree.
  -J            Prints out an JSON representation of the tree.
  -H baseHREF   Prints out HTML format with baseHREF as top directory.
  -T string     Replace the default HTML title and H1 header with string.
  --nolinks     Turn off hyperlinks in HTML output.
  --hintro X    Use file X as the HTML intro.
  --houtro X    Use file X as the HTML outro.
  ------- Input options -------
  --fromfile    Reads paths from files (.=stdin)
  --fromtabfile Reads trees from tab indented files (.=stdin)
  --fflinks     Process link information when using --fromfile.
  ------- Miscellaneous options -------
  --version     Print version and exit.
  --help        Print usage and this help message and exit.
  --            Options processing terminator.
```

```bash
用法: tree [-acdfghilnpqrstuvxACDFJQNSUX] [-L 层级 [-R]] [-H  baseHREF]
        [-T 标题] [-o 文件名] [-P 模式] [-I 模式] [--gitignore]
        [--gitfile[=]文件] [--matchdirs] [--metafirst] [--ignore-case]
        [--nolinks] [--hintro[=]文件] [--houtro[=]文件] [--inodes] [--device]
        [--sort[=]<名称>] [--dirsfirst] [--filesfirst] [--filelimit #] [--si]
        [--du] [--prune] [--charset[=]X] [--timefmt[=]格式] [--fromfile]
        [--fromtabfile] [--fflinks] [--info] [--infofile[=]文件] [--noreport]
        [--version] [--help] [--] [目录 ...]
  ------- 列表选项 -------
  -a            列出所有文件。
  -d            仅列出目录。
  -l            将符号链接当作目录来跟随。
  -f            打印每个文件的完整路径前缀。
  -x            仅停留在当前文件系统上。
  -L 层级       仅递归到指定层级的目录深度。
  -R            当达到最大目录层级时重新运行 tree。
  -P 模式       仅列出符合给定模式的文件。
  -I 模式       不列出符合给定模式的文件。
  --gitignore   使用 .gitignore 文件进行过滤。
  --gitfile X   显式读取 gitignore 文件。
  --ignore-case 忽略模式匹配时的大小写。
  --matchdirs   在 -P 模式匹配中包含目录名称。
  --metafirst   在每行的开头打印元数据。
  --prune       从输出中剪除空目录。
  --info        打印在 .info 文件中找到的信息。
  --infofile X  显式读取 info 文件。
  --noreport    关闭树状列表末尾的文件/目录计数。
  --charset X   使用字符集 X 进行终端/HTML 和缩进线输出。
  --filelimit # 如果目录中的文件超过 #，则不深入其中。
  -o 文件名     输出到文件而不是标准输出。
  ------- 文件选项 -------
  -q            将不可打印的字符打印为 '?'。
  -N            将不可打印的字符按原样打印。
  -Q            用双引号引用文件名。
  -p            打印每个文件的权限。
  -u            显示文件所有者或用户ID编号。
  -g            显示文件组所有者或组ID编号。
  -s            打印每个文件的大小（以字节为单位）。
  -h            以更易读的方式打印文件大小。
  --si          类似于 -h，但使用 SI 单位（1000 的幂）。
  --du          通过其内容计算目录的大小。
  -D            打印最后修改时间或 (-c) 状态更改时间。
  --timefmt <f> 按照格式 <f> 打印和格式化时间。
  -F            根据 ls -F 添加 '/', '=', '*', '@', '|' 或 '>'。
  --inodes      打印每个文件的 inode 编号。
  --device      打印每个文件所属的设备 ID 编号。
  ------- 排序选项 -------
  -v            按版本对文件进行字母数字排序。
  -t            按最后修改时间对文件进行排序。
  -c            按最后状态更改时间对文件进行排序。
  -U            不对文件进行排序。
  -r            反转排序顺序。
  --dirsfirst   在文件之前列出目录 (-U 禁用此功能)。
  --filesfirst  在目录之前列出文件 (-U 禁用此功能)。
  --sort X      选择排序方式：名称、版本、大小、mtime、ctime。
  ------- 图形选项 -------
  -i            不打印缩进线。
  -A            使用 ANSI 线条图形进行缩进线打印。
  -S            使用 CP437（控制台）图形进行缩进线打印。
  -n            总是关闭颜色化（-C 覆盖此选项）。
  -C            总是开启颜色化。
  ------- XML/HTML/JSON 选项 -------
  -X            以 XML 形式打印树状结构。
  -J            以 JSON 形式打印树状结构。
  -H baseHREF   以 HTML 格式打印，并将 baseHREF 作为顶级目录。
  -T 字符串     用字符串替换默认的 HTML 标题和 H1 头部。
  --nolinks     在 HTML 输出中关闭超链接。
  --hintro X    使用文件 X 作为 HTML 引入部分。
  --houtro X    使用文件 X 作为 HTML 结束部分。
  ------- 输入选项 -------
  --fromfile    从文件中读取路径 (.=标准输入)
  --fromtabfile 从制表符缩进的文件中读取树 (.=标准输入)
  --fflinks     在使用 --fromfile 时处理链接信息。
  ------- 杂项选项 -------
  --version     打印版本并退出。
  --help        打印用法和本帮助消息并退出。
  --            选项处理终止符。
```

## fallocate

```bash
root@janus-a:~# fallocate --help

Usage:
 fallocate [options] <filename>

Preallocate space to, or deallocate space from a file.

Options:
 -c, --collapse-range remove a range from the file
 -d, --dig-holes      detect zeroes and replace with holes
 -l, --length <num>   length for range operations, in bytes
 -n, --keep-size      maintain the apparent size of the file
 -o, --offset <num>   offset for range operations, in bytes
 -p, --punch-hole     replace a range with a hole (implies -n)
 -z, --zero-range     zero and ensure allocation of a range
 -v, --verbose        verbose mode

 -h, --help     display this help and exit
 -V, --version  output version information and exit

For more details see fallocate(1).
root@janus-a:~# 
```



# 权限模式

##  四位八进制表示法

​	八进制表示法使用三位八进制数字来表示文件或目录的权限。每一位数字代表一组用户的权限，分别是文件所有者（user）、文件所属组（group）和其他用户（others）。

- 位数作用

  - 第一位: 特殊权限位

  - 第二位: 文件所有者的权限

  - 第三位: 文件所属组的权限

  - 第四位: 其他用户的权限

- 权限位的含义

  - 4: 读权限（read, r）

  - *: 写权限（write, w）

  - 1: 执行权限（execute, x）

- 权限组合

  - 7: 读、写、执行（rwx）

  - 6: 读、写（rw-）

  - 5: 读、执行（r-x）

  - 4: 读（r--）

  - 3: 写、执行（-wx）

  - 2: 写（-w-）

  - 1: 执行（--x）

  - 0: 无权限（---）

##  符号表示法

符号表示法使用字符来表示用户类别和权限。

- 用户类别: 
  - u: 文件所有者（user）
  - g: 文件所属组（group）
  - o: 其他用户（others）
  - a: 所有用户（all，等同于 `ugo`）
- 操作符: 
  - +: 添加权限
  - -: 移除权限
  - =: 设置权限
- 权限: 
  - r: 读权限（read）
  - w: 写权限（write）
  - x: 执行权限（execute）

# shell 特殊符号

| 符号 | 作用                                         |
| ---- | -------------------------------------------- |
| \|   | 将前一个命令的输出作为下一个命令的输入       |
| >    | 将命令的输出重定向到文件，覆盖文件内容       |
| >>   | 将命令的输出重定向到文件，追加到文件末尾     |
| &    | 将命令放入后台执行                           |
| &&   | 只有前一个命令执行成功时，才会执行下一个命令 |
| *    | 匹配任意数量的字符（包括零个字符）           |
| ？   | 匹配一个字符                                 |
