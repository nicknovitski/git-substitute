#!/usr/bin/env bats

setup() {
  mkdir tmp
  cd tmp
  git init
}

teardown() {
  cd ..
  rm -rf tmp
}

@test "substitutes strings across the repo" {
  echo foo > test1
  echo foo > test2
  git add .
  ../bin/git-substitute foo bar
  [ `cat test1` = bar ]
  [ `cat test2` = bar ]
}

@test "first argument interprets regex metacharacters" {
  echo people > test1
  echo peoples > test2
  git add .
  ../bin/git-substitute peoples? persons
  [ `cat test1` = persons ]
  [ `cat test2` = persons ]
}

@test "second argument treats regex metacharacters literally" {
  echo foo > test1
  git add .
  ../bin/git-substitute foo ^.?*+[]{}
  [ `cat test1` = "^.?*+[]{}" ]
}

@test "first argument can accept escaped regex metacharactrs" {
  echo '$50.00' > test1
  git add .
  ../bin/git-substitute '\$50\.00' 49.99
  [ `cat test1` = 49.99 ]
}

@test "first argument treats quotes literally" {
  echo "'" > test1
  git add .
  ../bin/git-substitute \' bar
  [ `cat test1` = bar ]
}

@test "second argument can have backreferences" {
  echo "User.find_by_name('charles')" > test1
  git add .
  ../bin/git-substitute 'User\.find_by_name\((.*)\)' 'User.where(name: \1).first'
  [ "`cat test1`" = "User.where(name: 'charles').first" ]
}
