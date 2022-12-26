
<script>
	import {
		Timeline,
		TimelineItem,
		TimelineSeparator,
		TimelineDot,
		TimelineConnector,
		TimelineContent,
		TimelineOppositeContent
	} from 'svelte-vertical-timeline';

	const options = [
		{ title: "Mobile Telephony", discription: 'The first successful mobile telephony service was offered to first class passengers on the Deutsche Reichsbahn on the route between Berlin and Hamburg', time: '1926' },
		{ title: "First radiotelephone call from car", discription: 'First calls were made on a car radiotelephone in Chicago. Due to the small number of radio frequencies available, the service quickly reached capacity.', time: '1946' },
		{ title: "Sweden mobile phone system", discription: 'First automated mobile phone system for private vehicles launched in Sweden. The device to install in the car used vacuum tube technology with rotary dial and weighed 40Kg', time: '1956' },
		{ title: "NMT Group", discription: 'The Nordic Mobile Telephone (NMT) Group was established', time: '1969' },
		{time: '1987', title: "GMS Standards", discription: "The Technical specifications for the GSM standard are approved"},
		{time: '1992', title: "SMS Message", discription: "The world’s first ever SMS message was sent in the UK, Vodafone"},
		{time: '1996/97', title: "UK Mobile", discription: "UK phone ownership stood at 16% of households."},
		{time: '1999', title: "BlackBerry", discription: "BlackBerry phone was unveiled"},
		{time: '2000', title: "Popular Phones", discription: "Nokia 3310; Sharp J-SH04 first camera phone"},
		{time: "2003", title: "G3", discription: ""},
		{time: '2006', title: "UK Mobile", discription: "UK phone ownership stood at 80% of households."},
		{time: "2007", title: "Apple iPhone", discription: ""},
		{time: "2008", title: "More Firsts", discription: "First Android; Apple’s App Store and Android Market"},
		{time: "2009", title: "G4", discription: ""},
		{time: "2010", title: "Galaxy S", discription: ""},
		{time: "2017", title: "Galaxy S8 / iPhone X",discription: ""},
		{time: "2018", title: "Galaxy S9 / iPhone XS, XS Max, XR", discription: ""},
		{time: "2019", title: "Galaxy S10 / iPhone 11", discription: ""},
		{time: "2020", title: "Galaxy S20 / iPhone 12", discription: ""},
		{time: "2021", title: "Galaxy S21 / iPhone 13", discription: ""},
		{time: "2022", title: "Galaxy S22 / iPhone XS, XS Max", discription: ""}

	];
</script>
<div style="text-align: left; padding-left: 0.5em; padding-top: 1em;">https://www.uswitch.com/mobiles/guides/history-of-mobile-phones/</div>
<div  style="display: flex; align-items: left; padding: 10px 40px; border: solid 5px black;">
	<Timeline >
		{#each options as option}
			<TimelineItem position="left">
				<TimelineOppositeContent slot="opposite-content">
					<p>{option.time}</p>
				</TimelineOppositeContent>
				<TimelineSeparator>
					<TimelineDot style={'width: 20px; height: 20px; background-color: #007ba7; display: flex; justify-content: center; border-color: transparent; '}/>
					<TimelineConnector />
				</TimelineSeparator>
				<TimelineContent>
					<h3>{option.title}</h3>
					{option.discription}
				</TimelineContent>
			</TimelineItem>
		{/each}
	</Timeline>
</div>