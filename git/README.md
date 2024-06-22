### Quick Summary

1. **Create Feature Branch**: Always start with a new branch.

```sh
git checkout -b feature-branch
```

2. **Commit Changes Frequently**: Make regular, small commits with clear messages.

```sh
git add .
git commit -m "your commit message"
```

3. **Pull from Main/Dev Regularly**: Keep your branch updated with the latest changes.

```sh
git pull origin main
git pull origin dev
```

4. **Push Branch**: Push your branch to the remote repository.

```sh
git push origin feature-branch
```

5. **Create Pull Request**: Submit your branch for review and merging.
6. **Resolve Merge Conflicts**: Address any conflicts that arise during the pull request process.
7. **Merge Branch**: Once approved, merge your branch into the main branch.

### Updating Your Branch to Keep Up-to-Date with dev

assume you are working on a `your-feature-branch` and you want to update your branch with the latest changes from the `dev`.

1. Switch to the `dev` branch:

```sh
git checkout dev
```

2. Pull the latest changes from the remote `dev` branch:

```sh
git pull origin dev
```

3. Switch back to your feature branch:

```sh
git checkout <your-feature-branch>
```

#### Avoid extra merge commits

Instead of merging the `dev` branch into your feature branch, you can rebase your feature branch on top of the `dev` branch.

````sh
we need to pull the latest changes from `dev` branch

1. Switch to the `dev` branch:

```sh
git checkout dev
````

2. Pull the latest changes from the remote `dev` branch:

```sh
git pull origin dev
```

3. Switch back to your feature branch:

```sh
git checkout <your-feature-branch>
```

4. Rebase your feature branch on top of the `dev` branch:

```sh
git rebase dev
```

> Rebasing will put `dev` branch commits under your `feature-branch` commits.

5. If there are any conflicts, resolve them and continue the rebase:

```sh
git rebase --continue
```

6. Once the rebase is complete, Your might need to use `-f` to force push

```sh
git push origin -f <your-feature-branch>
```

### Creating a Pull Request (PR)

1. Push your branch to the remote repository

```sh
git push origin <your-feature-branch>
```

2. Create a Pull Request on GitHub

- Go to your repository on the web platform.
- Click on the "New Pull Request" or "Create Merge Request" button.
- Select your branch (source branch) and the branch you want to merge into (target branch, such as dev or main).
- Write descriptions
- Submit the pull request.
- Merge the Pull Request **once approved!**.
