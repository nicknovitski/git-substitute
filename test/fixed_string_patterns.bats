#!/usr/bin/env bats

load isolation
load helpers

@test "with -F, escaped and unescaped round brackets are treated as literals" {
  given_file parens '()'
  given_file slashed '\('
  git add .
  git-substitute -F '()' unescaped
  git-substitute -F '\(' escaped
  assert_file_contains parens unescaped
  assert_file_contains slashed escaped
}

@test "with --fixed-strings, escaped and unescaped round brackets are treated as literals" {
  given_file parens '()'
  given_file slashed '\('
  git add .
  git-substitute --fixed-strings '()' unescaped
  git-substitute --fixed-strings '\(' escaped
  assert_file_contains parens unescaped
  assert_file_contains slashed escaped
}

@test "with fixed-strings options, basic regex metacharacters are always treated as literals" {
  skip "I suspect these already work but I can't be bothered to design the tests at the moment"
}
