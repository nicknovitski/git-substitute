#!/usr/bin/env bats

load isolation
load helpers

@test "without a path argument, substitutes in all tracked files" {
  given_file test1 foo
  given_file test2 foo
  git add .
  git-substitute foo bar
  assert_file_contains test1 bar
  assert_file_contains test2 bar
}

@test "subsitutes strings in files in path" {
  mkdir tmp
  given_file tmp/added foo
  git add tmp
  git-substitute foo bar tmp
  assert_file_contains tmp/added bar
}

@test "accepts multiple paths" {
  given_file path1 foo
  given_file path2 foo
  git add path1
  git add path2
  git-substitute foo bar path1 path2
  assert_file_contains path1 bar
  assert_file_contains path2 bar
}

@test "ignores files not in repository" {
  given_file unadded foo
  given_file added foo
  git add added
  git-substitute foo bar
  assert_file_contains unadded foo
}

@test "ignores files not in path" {
  mkdir searched
  mkdir unsearched
  given_file searched/test foo
  given_file unsearched/test foo
  git add searched unsearched
  git-substitute foo bar searched
  assert_file_contains unsearched/test foo
}
