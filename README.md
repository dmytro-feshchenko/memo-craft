# MemoCraft

## About

A lightweight yet powerful desktop app for note-taking fused with a flashcards app. Developed with love for personal use,
but feel free to use it to your heart's content.

Tech stack: Golang (latest), Vue 3, TypeScript, Vite

Built with [Wails v2](https://wails.io/) as a lightweight alternative to Electron written in Go

Supports [CommonMark Spec](https://spec.commonmark.org/0.29/)

## The idea!

The idea was born after years of trying different apps to help manage my PKM (personal knowledge management) and assist with self-education. But, alas, after trying dozens of options, I concluded that the app that suits my needs doesn't exist.

What features do I want to see in the PKM tool:
- Open-source
- Has a built-in markdown editor with an elegant inline preview
- Flashcards can be created directly in the editor and "related" to the note yet be used for dedicated spaced repetition from a central hub.
- Store every piece of data as plain files (like md), which makes it future-proof and able to sync files using any preferable file syncing method (e.g. Dropbox/Google Drive/iCloud, or even GitHub files).
- Can add attachments (at least images) and embed videos (e.g. YouTube).
- Has aesthetically pleasing UI which won't force me to live in 90s =)

Obsidian was the closest one, but, unfortunately, it's not open-source, and managing flashcards is tricky and can be done only using 3rd party plugins and integrations with other apps (such as Anki).
So, here we are.

## Supported build platforms

- Windows 10/11 AMD64/ARM64
- MacOS 10.13+ AMD64
- MacOS 11.0+ ARM64 - *my love!*
- Linux AMD64/ARM64

## How to build?

### Dev Dependencies

- Go 1.20+
- NPM (Node 15+)

For Windows only: [WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/)

For MacOS only: `xcode-select --install`

For Linux only: Linux requires the standard gcc build tools plus libgtk3 and libwebkit

### Optional Dependencies

- UPX for compressing your applications.
- NSIS for generating Windows installers.

### Todo:
- [ ] Flashcards
    - [ ] Card preview (in progress)
    - [ ] Card management
    - [ ] Central hub
      - [ ] List of decks
      - [ ] List of previews
      - [ ] Queue runner
      - [ ] [Future] Stats
      - [ ] [Future] Configure repetition algorithm
- [ ] Notes
  - [ ] Basic note engine for text (in progress)
  - [ ] Advanced note engine
    - [ ] Code blocks (in progress)
    - [ ] Images
    - [ ] [Future] Embedded video
- [ ] Settings
  - [ ] Multi-language (in progress)
  - [ ] Themes (in progress)
  - [ ] Configuration (in progress)
- [ ] Navigation
  - [ ] File tree
  - [ ] Tabs
    - [ ] [Future] Store opened tabs
  - [ ] Full-text search

### How does it work?

#### Editor

I'm building inline markdown editor based on [Marked](https://marked.js.org/), which is a lightweight low-level markdown compiler.
Also, I'm going to use [marked-highlight](https://www.npmjs.com/package/marked-highlight) package to highlight code blocks.

### Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
