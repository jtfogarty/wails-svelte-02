<script>
	import File from './File.svelte';

	export let expanded = false;
	export let name;
	export let files;

	function toggle() {
		expanded = !expanded;
	}
</script>

<button class:expanded on:click={toggle}>{name}</button>

{#if expanded}
	<ul>
		{#each files as file}
			<li>
				{#if file.files}
					<svelte:self {...file}/>
				{:else}
					<File {...file}/>
				{/if}
			</li>
		{/each}
	</ul>
{/if}

<style>
	button {
		padding: 0 0 0 1.5em;
		background: url(/src/assets/images/folder.svg) 0 0.1em no-repeat;
		background-size: 1em 1em;
		font-weight: bold;
        color:#eee;
		cursor: pointer;
		border: none;
		margin: 0;
	}
	button:hover {
		color: #007ba7
	}

	.expanded {
		background-image: url(/src/assets/images/folder-open.svg);
	}
	.expanded:hover {
		color: #007ba7;
	}
	ul {
		padding: 0.2em 0 0 0.5em;
		margin: 0 0 0 0.5em;
		list-style: none;
		border-left: 1px solid #eee;
        margin-left: 0;
	}

	li {
		padding: 0.2em 0;
        margin-left: 0;
	}


</style>