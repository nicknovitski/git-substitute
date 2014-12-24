given_file() {
  echo "$2" > "$1"
}

assert_file_contains() {
  [ "`cat $1`" = "$2" ]
}
