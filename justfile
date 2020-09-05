#!/usr/bin/env just --justfile

# ==================simple one======================

version := "0.2.7"
tardir  := "awesomesauce-" + version

alias alias_t := recipe-name

# just会在运行真正命令之前，将每个命令打印到标准错误(stderr)，这就是为什么echo 'This is a recipe!'会被打印。

# just version=1 recipe-name
recipe-name:
    echo 'This is a recipe! with var {{tardir}}'

# 这是一个注释 当然你可以使用@作为行开头，这样会抑制打印。
another-recipe:
    @echo 'This is another recipe.'

# 转义
braces:
	echo '{{'I {{LOVE}} curly braces!'}}'

system-info:
	@echo "This is an {{arch()}} {{os()}} machine".

# env_var(key) - 用名称key检索环境变量，如果不存在会中止.
# env_var_or_default(key, default) - 用名称key检索环境变量，如果它不存在则返回default值.

# 环境变量 .env
env:
    echo 'tt env $ENV1 '

dir:
    cd {{invocation_directory()}}; echo tt

# 接收参数
build target:
    @echo 'Building {{target}}...'

default := 'all'
# 接收参数
test target tests=default:
    @echo 'Testing {{target}}:{{tests}}...'

# search with "kw"
search QUERY:
    lynx https://www.google.com/?q={{QUERY}}

# 条件
conditional:
    #!/usr/bin/env sh
    if true; then
        echo 'True!'
    fi

# 循环
for:
    #!/usr/bin/env sh
    for file in `ls .`; do
        echo $file
    done

_silent-script:
    echo "silent one"

# ==================complex one======================





