#!/usr/bin/env bats

load isolation

@test "substitutes strings across the repo" {
  echo foo > test1
  echo foo > test2
  git add .
  git-substitute foo bar
  [ `cat test1` = bar ]
  [ `cat test2` = bar ]
}

@test "does not make new untracked files" {
  echo foo > test1
  git add .
  git-substitute foo bar
  [ `git status --porcelain 2>/dev/null| grep "^??" | wc -l` = 0 ]
}

@test "first argument interprets regex metacharacters" {
  echo people > test1
  echo peoples > test2
  git add .
  git-substitute peoples? persons
  [ `cat test1` = persons ]
  [ `cat test2` = persons ]
}

@test "second argument treats regex metacharacters literally" {
  echo foo > test1
  git add .
  git-substitute foo ^.?*+[]{}
  [ `cat test1` = "^.?*+[]{}" ]
}

@test "first argument can accept escaped regex metacharactrs" {
  echo '$50.00' > test1
  git add .
  git-substitute '\$50\.00' 49.99
  [ `cat test1` = 49.99 ]
}

@test "first argument treats quotes literally" {
  echo "'" > test1
  git add .
  git-substitute \' bar
  [ `cat test1` = bar ]
}

@test "second argument can have backreferences" {
  echo "User.find_by_name('charles')" > test1
  git add .
  git-substitute 'User\.find_by_name\((.*)\)' 'User.where(name: \1).first'
  [ "`cat test1`" = "User.where(name: 'charles').first" ]
}
