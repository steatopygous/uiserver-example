<script>
    import {todos} from "./stores/todos";

    export let heading;
    export let items;

    function toggle(id) {
        todos.toggle(id);
    }
</script>

<div class="list">
    <h3>{heading}</h3>
    <ul>
        {#each items as id}
            <li on:click={() => toggle(id)}>{$todos[id].text}
        {/each}
    </ul>
</div>

<style>
    .list {
        display: flex;
        flex-direction: column;
        justify-items: start;

        padding-left: 10px;
        width: 200px;
    }

    ul {
        padding-left: 0;
    }

    li {
        list-style-type: none;
        padding-left: 0;
    }
</style>
