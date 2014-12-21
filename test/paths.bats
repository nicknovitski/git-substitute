#!/usr/bin/env bats

load isolation

@test "subsitutes strings in files in path" {
  mkdir tmp
  echo foo > tmp/added
  git add tmp
  git-substitute foo bar tmp
  [ `cat tmp/added` = "bar" ]
}

@test "accepts multiple paths" {
  echo foo > path1
  git add path1
  echo foo > path2
  git add path2
  git-substitute foo bar path1 path2
  [ `cat path1` = "bar" ]
  [ `cat path2` = "bar" ]
}

@test "ignores files not in repository" {
  echo foo > unadded
  echo foo > added
  git add added
  git-substitute foo bar
  [ `cat unadded` = "foo" ]
}

@test "ignores files not in path" {
  mkdir searched
  mkdir unsearched
  echo foo > searched/test
  echo foo > unsearched/test
  git add searched unsearched
  git-substitute foo bar searched
  [ `cat unsearched/test` = "foo" ]
}
