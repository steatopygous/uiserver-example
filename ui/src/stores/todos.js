import { writable } from 'svelte/store'

export const todos = createTodos();

function createTodos() {
    let items = {};

    const { subscribe, set, update } = writable(items);

    loadItems();

    return {
        subscribe,

        async add(text) {
            const body = { text };

            const response = await fetch('/api/todos', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(body)
            });

            if (response.ok) {
                const { id } = await response.json();

                update(items => {
                    items[id] = { id, text, done: false };

                    return items;
                });
            }
        },

        async toggle(id) {
            const body = { done: !items[id].done };

            const response = await fetch(`/api/todos/${id}`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(body)
            });

            if (response.ok) {
                update(_ => {
                    items[id].done = !items[id].done;

                    return items;
                });
            }
        },

        async remove(id) {
            const body = { id };

            const response = await fetch(`/api/todos/${id}`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(body)
            });

            if (response.ok) {
                update(_ => {
                    delete items[id];

                    return items;
                });
            }
        },

        async purge() {
            const body = {};

            const response = await fetch(`/api/todos/purge`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(body)
            });

            if (response.ok) {
                update(_ => {
                    const done = Object.keys(items).filter(id => items[id].done);

                    for (let id of done) {
                        delete items[id];
                    }

                    return items;
                });
            }
        },
    };

    async function loadItems() {
        const response = await fetch('/api/todos');

        console.log('todos.js loadItems() - response.ok =', response.ok);

        if (response.ok) {
            items = await response.json();

            console.log('todos.js loadItems() - items =', items);

            set(items);
        }
    }
}

