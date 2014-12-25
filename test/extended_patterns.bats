#!/usr/bin/env bats

load isolation
load helpers

@test "first argument interprets regex metacharacters" {
  given_file test1 people
  given_file test2 peoples
  git add .
  git-substitute peoples? persons
  assert_file_contains test1 persons
  assert_file_contains test2 persons
}

@test "first argument can accept escaped regex metacharactrs" {
  given_file pricetag '$50.00'
  git add .
  git-substitute '\$50\.00' 49.99
  assert_file_contains pricetag 49.99 
}

@test "first argument treats quotes literally" {
  given_file quoteme "'"
  git add .
  git-substitute \' bar
  assert_file_contains quoteme bar
}

@test "second argument can have backreferences" {
  given_file test1 "User.find_by_name('charles')"
  git add .
  git-substitute 'User\.find_by_name\((.*)\)' 'User.where(name: \1).first'
  assert_file_contains test1 "User.where(name: 'charles').first"
}
