<script>
    import { quintOut } from 'svelte/easing';
    import { crossfade } from 'svelte/transition';

    import { todos } from './stores/todos.js';
    import { preferences } from './stores/preferences.js';

    import ItemList from './ItemList.svelte'

    $: ids = Object.keys($todos);

    $: done = ids.filter(id => $todos[id].done);
    $: pending = ids.filter(id => !$todos[id].done);

    let outerWidth;
    let outerHeight;

    $: setWindowSize(outerWidth, outerHeight);

    function addItem(target) {
        const text = target.value.trim();

        if (text.length > 0) {
            todos.add(text);

            target.value = '';
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
</script>

<div class='board'>
    <input
        placeholder="What needs to be done?"
        on:keydown={e => e.key === 'Enter' && addItem(e.target)}
    >

    <div class='left'>
        <ItemList heading="Pending" items={pending}/>
    </div>

    <div class='right'>
        <ItemList heading="Completed" items={done}/>
    </div>
</div>

<style>
    .board {
        display: grid;
        grid-template-columns: 1fr 1fr;
        grid-gap: 1em;
        max-width: 36em;
        margin: 0 auto;
    }

    .board > input {
        font-size: 1.4em;
        grid-column: 1/3;
    }
</style>
