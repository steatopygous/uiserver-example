<script>
    import { todos } from './stores/todos.js';
    import { preferences } from './stores/preferences.js';

    import ItemList from './ItemList.svelte';

    $: ids = Object.keys($todos);

    let text = '';

    $: done = ids.filter(id => $todos[id].done);
    $: pending = ids.filter(id => !$todos[id].done);

    let outerWidth;
    let outerHeight;

    $: setWindowSize(outerWidth, outerHeight);

    function addItem() {
        text = text.trim();

        if (text.length > 0) {
            todos.add(text);

            text = '';
        }
    }

    function purge() {
        todos.purge();
    }

    function setWindowSize(outerWidth, outerHeight) {
        if (outerWidth !== undefined) {
            console.log('setWindowSize() - outerWidth =', outerWidth, 'outerHeight =', outerHeight);

            preferences.setWindowSize(outerWidth, outerHeight);
        }
    }

    function keyUp(e) {
        if (e.code === 'Enter') {
            addItem();
        }
    }
</script>

<svelte:window
        bind:outerWidth={outerWidth} bind:outerHeight={outerHeight}
        on:keyup={keyUp}
/>

<main>
    <h1>Todos</h1>

    <div>
        <input type="text" bind:value={text} />
        <button on:click={addItem}>Add</button>
    </div>

    <div class="lists">
        <ItemList heading="Pending" items={pending} />
        <ItemList heading="Complete" items={done} />
    </div>

    <button on:click={purge}>Purge</button>
</main>

<style>
	main {
	    display: flex;
	    flex-direction: column;
	    align-items: center;
        width: 100%;
        padding-top: 0;
        margin-top: 0;
	}

    input {
        display: inline-block;
    }

    .lists {
	    display: flex;
	    flex-direction: row;
        justify-content: center;
        width: 100%;
    }

	h1 {
	    margin-top: 0;
	    padding-top: 0;
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 3em;
		font-weight: 100;
	}

    button {
        margin-top: 10px;
    }
</style>
