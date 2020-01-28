# Curator

Curator is a Daemon which watches a directory and converts any books in the directory to `.mobi` files and places the converted books to another directory.

Also it can optionally send the converted books as an email. This can enable usecases where the books can be sent to Kindle email automatically.

## Prerequisites

This uses `ebook-convertor` tool from `calibre`. So `calibre` is expected to be installed.

You can see instructions to install [here](https://calibre-ebook.com/download).

## Installation

Precompiled binary can be downloaded from [here](https://github.com/iamd3vil/curator/releases).

`calibre` needs to be installed

## Configuration

`curator` will check for a `curator.toml` file in the same directory. A sample config file `curator.toml.sample` can be found in the repo.
