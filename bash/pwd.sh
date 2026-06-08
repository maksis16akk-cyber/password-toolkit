#!/bin/bash

check_strength() {
    local pwd="$1"
    local score=0
    [ ${#pwd} -ge 8 ] && ((score++))
    [[ "$pwd" =~ [0-9] ]] && ((score++))
    [[ "$pwd" =~ [A-Z] ]] && [[ "$pwd" =~ [a-z] ]] && ((score++))
    [[ "$pwd" =~ [^a-zA-Z0-9] ]] && ((score++))
    case $score in
        1) echo "Weak" ;;
        2) echo "Medium" ;;
        3) echo "Strong" ;;
        4) echo "Very Strong" ;;
        *) echo "Weak" ;;
    esac
}

generate() {
    local len=$1
    tr -dc 'A-Za-z0-9!@#$%^&*()_+' </dev/urandom | head -c "$len"
    echo
}

case "$1" in
    -g|--generate)
        LEN=${2:-12}
        PWD=$(generate "$LEN")
        echo "Generated: $PWD"
        echo "Strength: $(check_strength "$PWD")"
        ;;
    -c|--check)
        if [ -z "$2" ]; then echo "Usage: $0 -c <password>"; exit 1; fi
        echo "Strength: $(check_strength "$2")"
        ;;
    *)
        echo "Usage: $0 -g [length] or -c <password>"
        ;;
esac
