#!/usr/bin/env bats

load isolation

@test "give standard git error when run outside of a git repository" {
  cd ..
  run git-substitute
  [ "$status" -eq 1 ]
  [ "$output" = "fatal: Not a git repository (or any of the parent directories): .git" ]
}

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
