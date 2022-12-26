<script>
	export let name;
	$: type = name.slice(name.lastIndexOf('.') + 1);

</script>

<!-- <span style="background-image: url(/src/assets/images/{type}.svg)" onmouseover="this.style.backgroundColor='#007ba7'" onmouseout="this.style.backgroundColor=''">{name}</span>

<style>
	span {
		padding: 0 0 0 1.5em;
		background: 0 0.1em no-repeat;
		background-size: 1em 1em;
	}
</style> -->

<button style="background-image: url(/src/assets/images/{type}.svg)" on:click={() => alert(name)}>{name}</button>

<style>
	button {
		padding: 0 0 0 1.5em;
		background: 0 0.1em no-repeat;
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
</style>