# README.md

## How to run using poetry

Reference links:
- https://medium.com/staqu-dev-logs/keeping-python-code-clean-with-pre-commit-hooks-black-flake8-and-isort-cac8b01e0ea1
- [Poetry](https://python-poetry.org/docs/)
- [Commitizen](https://github.com/commitizen-tools/commitizen)
- https://elegant.oncrashreboot.com/use-pre-commit-git-hooks
- https://elegant.oncrashreboot.com/use-commitizen-for-git-commits

### 1. Install poetry

```bash
curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/get-poetry.py | python -
```

### 2. Install packages and dependencies

```zsh
poetry install
```

### 3. Development

```bash


# make a repo and get cracking
mkdir src/pXXX
cd src/pXXX
touch main.py
touch test_main.py

# go into shell
poetry shell
poetry env info

# Use iPython to code and debug solutions.
ipython

# once finished editing
git add .
# it is probably a good practice to run pre-commit before we start commitizen...
pre-commit run
cz commit

# OR

cd /projectEuler/python3
poetry shell
pre-commit install --hook-type commit-msg
git add .
cz c
# run through the prompt from commitizen
# then it will lint check with pre-commit

husky > pre-commit (node v14.16.1)
ℹ No staged files match any configured task.
[WARNING] Unstaged files detected.
[INFO] Stashing unstaged files to /Users/user-xyz/.cache/pre-commit/patch1621857346.
commitizen check.....................................(no files to check)Skipped
black................................................(no files to check)Skipped
flake8...............................................(no files to check)Skipped
isort................................................(no files to check)Skipped
Debug Statements (Python)............................(no files to check)Skipped
[INFO] Restored changes from /Users/user-xyz/.cache/pre-commit/patch1621857346.

[master 850b35a] docs(readme.md): updating commitizen docs
 1 file changed, 4 insertions(+), 1 deletion(-)


# once you're finished with passing checks, you can push
git push
# You're done!
```

### Run hooks manually without commiting

```bash
pre-commit run
```

Reference: https://elegant.oncrashreboot.com/use-pre-commit-git-hooks

### Running tests

```bash
cd /projectEuler/python3
pytest /src/pXX/test_main.py

# OR run on all with verbose turned ON

pytest -v
```



### Using pre-commit

https://pre-commit.com/#plugins

You can update your hooks to the latest version automatically by running `pre-commit autoupdate`. By default, this will bring the hooks to the latest tag on the default branch.


pre-commit run --all-files
