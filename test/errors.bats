#!/usr/bin/env bats

@test "exit with status 1 with no args" {
  run ./git-substitute
  [ "$status" -eq 1 ]
}

@test "exit with status 1 with 1 arg" {
  run ./git-substitute anything
  [ "$status" -eq 1 ]
}

@test "exit with status 1 if pattern not found" {
  run ./git-substitute 'w{4,}' 'www'
  [ "$status" -eq 1 ]
}
