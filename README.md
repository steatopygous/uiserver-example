## TL;DR
This example app implements a todo list that stores its data in a JSON file.
It uses [uiserver](https://github.com/steatopygous/uiserver) to provide
a REST API over that data, along with a UI that is included in the app
using an embedded file system, thanks to the
[embed package](https://golang.org/pkg/embed/) in Go 1.16.

The user interface implementation is based on the
[animation example](https://svelte.dev/tutorial/animate)
from the Svelte tutorial.

To build the UI
```shell
cd ui
npm i
npm run build
```

Run the app, from the root folder, using
```shell
go run .
```

There are a few command-line options available, which can be displayed using
(and are documented below)
```shell
go run . -help
```

To build a standalone executable that embeds both the server and UI
```shell
go build .
```
which will create a **uiserver-example** executable (or **uiserver-example.exe** on Windows).

## The Server

On startup, we create a **ToDoList** containing the data from the JSON file.
If none exists, one is created with a few example items. 

REST calls to the server return the ToDoList's content and manipulate it,
saving back to the JSON file whenever it is changed.  This is relatively expensive,
but fine for an example app, to avoid having a dependency on some kind of database.

Note that __uiserver__ is intended to provide a single-user app over local data.
While there is nothing stopping us from opening multiple browser tabs displaying
the UI, changes made in one will not be reflected in others unless their browser
tabs are refreshed.

## The Client

The Svelte app uses a store as an abstraction that keeps a local copy of the
todo list.  On startup, this is initialised by a GET request to the API.

Subsequent updates (adding new items, marking them as done, deleting individual items,
and purging all completed items) are done through methods on the store, which
only update local data once the corresponding API request has been successfully
completed.

## Command-Line Flags
By default, the server runs on port 1234.  This can be changed using **-port=PORT**.

The todo list data is stored in **todos.json** in the current folder, but
can be changed using **-json=PATH**.

The UI will be displayed in a new tab/window of the user's default browser thanks
to the [browser package](github.com/pkg/browser). It can be opened as a Chrome app
using **-display=chrome** (only tested on MacOS).  In that case, the window dimensions
are saved across sessions.

## Future Directions

The current UI is based on the standard
[Svelte 3 template](https://github.com/sveltejs/template).
Once [SvelteKit](https://kit.svelte.dev/docs) comes out of beta, it would be
nice to update to use it.

The UI catches API error responses, but currently provides no indication to the
user.  Errors are unlikely to occur, unless multiple browser tabs are accessing
the server, or the server has been closed down.

