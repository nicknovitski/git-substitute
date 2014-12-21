#!/usr/bin/env bats

load isolation

@test "with -G, first argument treats POSIX extended metacharacters as literals" {
  echo ? > question
  echo + > addition
  echo '|' > nestpas
  git add .
  git-substitute -G ? !
  git-substitute -G + -
  git-substitute -G '|' 'un pipe'
  [ `cat question` = ! ]
  [ `cat addition` = - ]
  [ "`cat nestpas`" = 'un pipe' ]
}

@test "with --basic-regexp, first argument treats POSIX extended metacharacters as literals" {
  echo ? > question
  echo + > addition
  echo '|' > nestpas
  git add .
  git-substitute --basic-regexp ? !
  git-substitute --basic-regexp + -
  git-substitute --basic-regexp '|' 'un pipe'
  [ `cat question` = ! ]
  [ `cat addition` = - ]
  [ "`cat nestpas`" = 'un pipe' ]
}

@test "with -G, round brackets are literal, and escaped round brackets make a grouped expression" {
  echo '(parenthetical) explicit' > parens
  git add .
  git-substitute -G '(\(.*\)) \(.*\)' '\2 (\1)'
  [ "`cat parens`" = 'explicit (parenthetical)' ]
}

@test "with --basic-regexp, round brackets are literal, and escaped round brackets make a grouped expression" {
  echo '(parenthetical) explicit' > parens
  git add .
  git-substitute --basic-regexp '(\(.*\)) \(.*\)' '\2 (\1)'
  [ "`cat parens`" = 'explicit (parenthetical)' ]
}

@test "with -G, angle brackets are literal, and escaped angle brackets make a count expression" {
  echo {{{ > brackets
  git add .
  git-substitute -G '{\{2,3\}' '{{'
  [ "`cat brackets`" = '{{' ]
}

@test "with --basic-regexp, angle brackets are literal, and escaped angle brackets make a count expression" {
  echo {{{ > brackets
  git add .
  git-substitute --basic-regexp '{\{2,3\}' '{{'
  [ "`cat brackets`" = '{{' ]
}
