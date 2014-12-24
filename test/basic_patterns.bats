#!/usr/bin/env bats

load isolation
load helpers

assert_treats_metacharacters_as_literals() {
  OPTION=$1
  given_file question ?
  given_file addition +
  given_file nestpas '|'
  git add .
  git-substitute $OPTION ? !
  git-substitute $OPTION + -
  git-substitute $OPTION '|' 'un pipe'
  assert_file_contains question !
  assert_file_contains addition -
  assert_file_contains nestpas 'un pipe'
}

@test "with -G, first argument treats POSIX extended metacharacters as literals" {
  assert_treats_metacharacters_as_literals -G
}

@test "with --basic-regexp, first argument treats POSIX extended metacharacters as literals" {
  assert_treats_metacharacters_as_literals --basic-regexp
}

assert_reverses_round_brackets() {
  OPTION=$1
  given_file parens '(parenthetical) explicit'
  git add .
  git-substitute $OPTION '(\(.*\)) \(.*\)' '\2 (\1)'
  assert_file_contains parens 'explicit (parenthetical)'
}

@test "with -G, round brackets are literal, and escaped round brackets make a grouped expression" {
  assert_reverses_round_brackets -G
}

@test "with --basic-regexp, round brackets are literal, and escaped round brackets make a grouped expression" {
  assert_reverses_round_brackets --basic-regexp
}

assert_reverses_angle_brackets() {
  OPTION=$1
  given_file brackets {{{
  git add .
  git-substitute $OPTION '{\{2,3\}' '{{'
  assert_file_contains brackets {{
}

@test "with -G, angle brackets are literal, and escaped angle brackets make a count expression" {
  assert_reverses_angle_brackets -G
}

@test "with --basic-regexp, angle brackets are literal, and escaped angle brackets make a count expression" {
  assert_reverses_angle_brackets --basic-regexp
}
