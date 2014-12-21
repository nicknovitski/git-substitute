#!/usr/bin/env bats

load isolation

@test "exit with status 1 with no args" {
  run git-substitute
  [ "$status" -eq 1 ]
}

@test "exit with status 1 with 1 arg" {
  run git-substitute anything
  [ "$status" -eq 1 ]
}

@test "exit with status 1 and no output if pattern not found" {
  run git-substitute anything something
  [ "$status" -eq 1 ]
  [ "$output" = "" ]
}
