#!/bin/bash

# バージョン情報
VERSION="0.0.0"

# ヘルプメッセージの表示関数
show_help() {
    echo "  -v   Display version information"
    echo "  -h   Show usage and option"
}

# バージョン情報の表示関数
show_version() {
    echo "Monkey version: $VERSION"
}

# 引数がない場合はインタプリタを起動
if [ $# -eq 0 ]; then
    go run main.go
else
    # Monkey scriptが指定されている場合はファイルを探して実行する
    file_name=$1
    if [ -e "$file_name" ]; then
        go run main.go $file_name
        exit 0
    else 
        # オプションの処理
        while getopts ":vh" opt; do
            case $opt in
                v)
                    show_version
                    exit 0
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
        # ファイルが存在しない場合はエラー
        echo "Error: $file_name is not found."
        exit 1
    fi

    exit 1
fi

exit 0