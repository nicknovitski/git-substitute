setup() {
  PATH=$BATS_TEST_DIRNAME/../bin/:$PATH
  cd $BATS_TMPDIR
  rm -rf $BATS_TEST_NAME
  mkdir $BATS_TEST_NAME
  cd $BATS_TEST_NAME
  git init
}
