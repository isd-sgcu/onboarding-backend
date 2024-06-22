# Git

## How to Collaborate with Git

### 1. Fork the repository

- Go to the repository you want to contribute to
- Click on the "Fork" button at the top right corner
- This will create a copy of the repository in your GitHub account
- You can now clone this repository to your local machine
- Click on the "Code" button and copy the URL
- Open your terminal and run the following command:

```bash
git clone <URL>
```

### 2. Add a remote upstream

- Change into the directory of the repository

```bash
cd <repository>
```

- Add a remote upstream to the original repository

```bash
git remote add upstream <URL>
```

- Verify the remote has been added

```bash
git remote -v
```

### 3. Create a new branch

- Before making any changes, create a new branch

```bash
git checkout -b <branch-name>
```

- Make sure you are on the new branch

```bash
git branch
```

### 4. Make changes

- Make changes to the files in the repository
- Add the changes to the staging area

```bash
git add .
```

- Commit the changes

```bash
git commit -m "Your commit message"
```

### 5. Push changes to your fork

- Push the changes to your fork on GitHub

```bash
git push origin <branch-name>
```

### 6. Create a pull request

- Go to your fork on GitHub
- Click on the "Compare & pull request" button
- Add a title and description to your pull request
- Click on the "Create pull request" button
- Wait for the repository owner to review and merge your changes
- If there are any conflicts, resolve them and push the changes to your fork
- Your pull request will be automatically updated
- Once your pull request is merged, you can delete the branch
- Change back to the main branch

```bash
git checkout main
```

- Delete the branch

```bash
git branch -d <branch-name>
```

### 7. Sync your fork

- Fetch the changes from the original repository

```bash
git fetch upstream
```

- Check out the main branch

```bash
git checkout main
```

- Merge the changes from the original repository

```bash
git merge upstream/main
```

- Push the changes to your fork on GitHub

```bash
git push origin main
```

### 8. Repeat

- Repeat the process for any new contributions you want to make
- Make sure to always sync your fork before creating a new branch
- Keep your fork up to date with the original repository
- Happy contributing to ISD projects!

### 9. Resources

- [GitHub Guides](https://guides.github.com/)
- [Pro Git Book](https://git-scm.com/book/en/v2)
- [Atlassian Git Tutorials](https://www.atlassian.com/git/tutorials)
- [Git Documentation](https://git-scm.com/doc)
- [Git Cheat Sheet](https://education.github.com/git-cheat-sheet-education.pdf)
- [Git Explorer](https://gitexplorer.com/)
- [Learn Git Branching](https://learngitbranching.js.org/)
- [Git Immersion](http://gitimmersion.com/)
