#!/usr/bin/env bats

load isolation
load helpers

assert_recognizes_slash_d() {
  OPTION=$1
  given_file twos 222
  given_file dees ddd
  git add .
  git-substitute $OPTION '\d+' foo
  assert_file_contains twos foo
  assert_file_contains dees ddd
}

@test "with -P, perl character classes become available" {
  assert_recognizes_slash_d -P
}

@test "with --perl-regexp, perl character classes become available" {
  assert_recognizes_slash_d --perl-regexp
}
