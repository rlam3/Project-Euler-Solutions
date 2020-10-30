# README.md

## How to run using poetry

Reference links: [Poetry](https://python-poetry.org/docs/)

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
mkdir src/pXXX

# go into shell
poetry shell
poetry env info

# Use iPython to code and debug solutions.
ipython

# once finished editing
git add . 
cz commit

# run through the prompt

git push 

# You're done
```

### Running tests

```bash
cd /projectEuler/python3
pytest
```
