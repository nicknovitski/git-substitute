#!/usr/bin/env bats

load isolation
load helpers

@test "does not make new untracked files" {
  given_file test1 foo
  git add .
  git-substitute foo bar
  [ `git status --porcelain 2>/dev/null| grep "^??" | wc -l` = 0 ]
}

@test "second argument treats regex metacharacters literally" {
  given_file test1 foo
  git add .
  git-substitute foo ^.?*+[]{}
  assert_file_contains test1 '^.?*+[]{}'
}
