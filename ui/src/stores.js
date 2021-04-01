import { writable } from 'svelte/store'

export const todos = createTodos();

function createTodos() {
    const { subscribe, update } = writable({});

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
                update(items => {
                    items[id].done = !items[id].done;

                    return items;
                });
            }
        },

        async delete(id) {
            const body = { id };

            const response = await fetch(`/api/todos/${id}`, {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(body)
            });

            if (response.ok) {
                update(items => {
                    delete items[id];

                    return items;
                });
            }
        }
    };

    async function loadItems() {
        console.log('stores.js loadItems()');

        const response = await fetch('/api/todos');

        console.log('stores.js loadItems() - response.ok =', response.ok);

        if (response.ok) {
            const items = await response.json();

            console.log('stores.js loadItems() - items =', items);

            update(_ => items);
        }
    }
}

