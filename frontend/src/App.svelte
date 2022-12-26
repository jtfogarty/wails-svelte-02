<script>
  import { HSplitPane } from 'svelte-split-pane';
	//import Cat from "./Cat.svelte";

	import Timeline from "./Timeline.svelte";
	
	// export let menu = 0;
	import Folder from './Folder.svelte';
	let root = [
		{
			name: 'Cat Facts',
			files: []
		},
		{
			name: 'Timeline',
			files: []
		}
	]
/* 	let root = [
		{
			name: 'Important work stuff',
			files: [
				{ name: 'quarterly-results.xlsx' }
			]
		},
		{
			name: 'Animal GIFs',
			files: [
				{
					name: 'Dogs',
					files: [
						{ name: 'treadmill.gif' },
						{ name: 'rope-jumping.gif' }
					]
				},
				{
					name: 'Goats',
					files: [
						{ name: 'parkour.gif' },
						{ name: 'rampage.gif' }
					]
				},
				{ name: 'cat-roomba.gif' },
				{ name: 'duck-shuffle.gif' },
				{ name: 'monkey-on-a-pig.gif' }
			]
		},
		{ name: 'TODO.md' }
	]; */

</script>

<HSplitPane rightPaneSize="75%" leftPaneSize="25%" minRightPaneSize="50px" minLeftPaneSize="50px" updateCallback={() => {
  console.log('HSplitPane Updated!');
}}>
  <left slot="left" align="left">
    <div style="padding-left: 1.5em; padding-top: 1.5em;">
      <Folder name="Home" files={root} expanded/>
    </div>
  </left>
  <right slot="right">
    <Timeline />
  </right>
</HSplitPane>

