# nebulariders-server

## Branching strategy

We have several branches:

1. main

- The production-ready state. This branch contains only the released code.

2. dev

- The integration branch for features. All new development should happen here.

3. feature branches

- Branch off from develop and merge back into develop when complete. Named like feature/feature-name.

4. release branches

- Once develop has enough features for a release, create a release branch from develop (e.g., release/1.0.0). Perform final testing and bug fixing here. When ready, merge into both main and develop.

5. hotfix branches

- For urgent fixes to production code. Branch off from main, apply the fix, then merge back into both main and develop.

Managing multiple versions is simple, when we want to change version lets say 2.0.0 we make a branch out of release/2.0.0 called release/2.0.1, then multiple feature branches that will be merged into release/2.0.1. Finally we dont merge 2.0.1 into dev because we might be on version 3.0.0 on dev or even more ahead.

## Merging or rebasing

- In order to merge your branch, use REBASE as it keeps the history cleaner and liniar

```
git checkout feature/feature-name
git rebase dev
```
