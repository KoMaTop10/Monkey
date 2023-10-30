#!/bin/bash

# バージョン情報
VERSION="0.0.0"

# ヘルプメッセージの表示関数
show_help() {
    echo "  -v   Display version information"
    echo "  -t   This is test option."
    echo "  -h   Show usage and option"
}

# バージョン情報の表示関数
show_version() {
    echo "Monkey version: $VERSION"
}

# オプションの処理
option_cmd() { 
    while getopts ":vth" opt; do
        case $opt in
            v)
                show_version
                exit 0
                ;;
            t) 
                go run main.go
                ;;
            h)
                echo "Hello! This is programming language Monkey!"
                echo "Show usage and option"
                show_help
                exit 0
                ;;
            \?)
                echo "Invalid option: -$OPTARG"
                show_help
                exit 1
                ;;
        esac
    done
}
# 引数がない場合の条件分岐
if [ $# -eq 0 ]; then
    go run main.go
else
    file_name=$1
    if [ -e "$file_name" ]; then
        cat $file_name
        exit 0
    else 
        echo "Error: $file_name is not found."
        exit 1
    fi

    option_cmd
    echo "Invalid argument. Use -h for help."
    exit 1
fi



exit 0






