#!/usr/bin/env bats

load isolation
load helpers

assert_interprets_question_mark_metacharacter() {
  OPTION=$1
  given_file test1 people
  given_file test2 peoples
  git add .
  git-substitute $OPTION peoples? persons
  assert_file_contains test1 persons
  assert_file_contains test2 persons
}

@test "with -E, first argument interprets question marks as metacharacters" {
  assert_interprets_question_mark_metacharacter -E
}

@test "with --extended-regexp, first argument interprets question marks as metacharacters" {
  assert_interprets_question_mark_metacharacter --extended-regexp
}

assert_treats_dollar_sign_as_literal() {
  OPTION=$1
  given_file pricetag '$50.00'
  git add .
  git-substitute $OPTION '\$50\.00' 49.99
  assert_file_contains pricetag 49.99 
}

@test "with -E, first argument treats escaped dollar signs as literals" {
  assert_treats_dollar_sign_as_literal -E
}

@test "with --extended-regexp, first argument treats escaped dollar signs as literals" {
  assert_treats_dollar_sign_as_literal --extended-regexp
}

assert_standard_regex_round_brackets() {
  OPTION=$1
  given_file test1 "User.find_by_name('charles')"
  git add .
  git-substitute $OPTION 'User\.find_by_name\((.*)\)' 'User.where(name: \1).first'
  assert_file_contains test1 "User.where(name: 'charles').first"
}

@test "with -E, round brackets make a grouped expression, and escaped round brackets are literal" {
  assert_standard_regex_round_brackets -E
}

@test "with --extended-regexp, round brackets make a grouped expression, and escaped round brackets are literal" {
  assert_standard_regex_round_brackets --extended-regexp
}
