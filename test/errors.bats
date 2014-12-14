#!/usr/bin/env bats

@test "exit with status 1 with no args" {
  run bin/git-substitute
  [ "$status" -eq 1 ]
}

@test "exit with status 1 with 1 arg" {
  run bin/git-substitute anything
  [ "$status" -eq 1 ]
}

@test "exit with status 1 if pattern not found" {
  run bin/git-substitute 'w{4,}' 'www'
  [ "$status" -eq 1 ]
}
