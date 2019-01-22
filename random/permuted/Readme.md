# permuted

This is a simple program which randomly renames files or directories in a given
directory. Optionally it saves the used permutation together with the list of
files/directories on which it acted so that the action can be reverted.

While it can seem pretty useless, it was actually used once for shuffling
results of some biological experiments, so that the observer can't tell what is
the number of an experiment by the directory name.

*Example.* Use the program to shuffle directories in the directory `test`:
```
  $ ./permuted --target test --save saved.perm --dirs
```
Then reverse the action (the permutation was saved in the `test/saved.perm`):
```
  $ ./permuted --target test --perm saved.perm --dirs --inv
```
