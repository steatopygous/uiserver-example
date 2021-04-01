<script>
    import { todos } from './stores.js';

    let ids = Object.keys($todos);

    let text = '';

    $: done = ids.filter(id => $todos[id].done);
    $: pending = ids.filter(id => !$todos[id].done);

    function addItem() {
        todos.add(text);

        text = '';
    }
</script>

<main>
    <h1>Todos</h1>

    <div>
        <input type="text" bind:value={text} />
        <button on:click={addItem}>Add</button>
    </div>

    <div class="todos">
        <div class="list">
            <ul>
            {#each done as item}
                <li>{item.text}
            {/each}
            </ul>
        </div>

        <div class="list">
            <ul>
            {#each pending as item}
                <li>{item.text}
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
	    align-items: center;
    }

    .list {
        width: 500px;
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
