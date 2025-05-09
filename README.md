# Panache

> A small web app developed in Go.

## Installation

### Development environment

```bash
git clone https://github.com/alexisbcz/panache.git # clone the Git repository
cd panache # change directory to Git repository

make setup # install deps
make db-up # spin up postgres container
make migration-up # migrate the database
make seed # seed the database
```

To get started, open a new terminal launching the development server:

```bash
make css
```

And in another one:

```bash
make dev
```

### Production environment

_WIP_
