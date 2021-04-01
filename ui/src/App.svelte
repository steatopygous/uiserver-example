<script>
    import { todos } from './stores/todos.js';
    import { preferences } from './stores/preferences.js';

    $: ids = Object.keys($todos);

    let text = '';

    $: done = ids.filter(id => $todos[id].done);
    $: pending = ids.filter(id => !$todos[id].done);

    let outerWidth;
    let outerHeight;

    $: setWindowSize(outerWidth, outerHeight);

    function addItem() {
        todos.add(text);

        text = '';
    }

    function toggle(id) {
        todos.toggle(id);
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
</script>

<svelte:window bind:outerWidth={outerWidth} bind:outerHeight={outerHeight} />

<main>
    <h1>Todos</h1>

    <button on:click={purge}>Purge</button>
    <div>
        <input type="text" bind:value={text} />
        <button on:click={addItem}>Add</button>
    </div>

    <div class="todos">
        <div class="list">
            <h3>Pending</h3>
            <ul>
            {#each pending as id}
                <li on:click={() => toggle(id)}>{$todos[id].text}
            {/each}
            </ul>
        </div>

        <div class="list">
            <h3>Done</h3>
            <ul>
            {#each done as id}
                <li on:click={() => toggle(id)}>{$todos[id].text}
            {/each}
            </ul>
        </div>
    </div>
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

    .todos {
	    display: flex;
	    flex-direction: row;
        justify-content: center;
        width: 100%;
    }

    .list {
        display: flex;
        justify-items: start;
        width: 300px;
    }

	h1 {
	    margin-top: 0;
	    padding-top: 0;
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
</style>
