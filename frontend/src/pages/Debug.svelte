<script>
	import { h } from "gridjs";
	import Grid from "gridjs-svelte";
	import { SvelteWrapper } from "gridjs-svelte/plugins";
	import { RowSelection } from "gridjs/plugins/selection";

	// https://gridjs.io/docs/plugins/basics

	
	// https://gridjs.io/docs/config/className
	const className = {
		error: "error",
	};

	// https://gridjs.io/docs/config/style
	const style = {
		table: {
			width: "100%",
		},
		header: {
			display: "flex",
			alignItems: "center",
			flexDirection: "row-reverse",
		},
	};

	// https://gridjs.io/docs/config/language
	const language = {
		search: {
			placeholder: "Find people",
		},
	};

	// https://gridjs.io/docs/config/columns
	const columns = [
		
		"Name",{
		        id: 'myCheckbox',
        name: 'Select',
        plugin: {
          // install the RowSelection plugin
          component: RowSelection,
        }},"ddd",
		{
			name: "Gnder",
			width: "10%",
			formatter: cell => {
				if (cell === "n/a") {
					return "???";
				}

				return cell
					.toString()
					.replace(/\w\S*/g, txt => {
						return [
							txt.charAt(0).toUpperCase(),
							txt.substring(1).toLowerCase()
						].join("")
					});
			},
		},
		{
			name: "Eye Color",
			width: "5%",
			plugin: {
				component: SvelteWrapper,
				props: {
				},
			},
		},
		{
			name: "Action",
			width: "10%",
			formatter: (cell, row) => {
				console.log("Detail", cell)
				console.log(row)
				// https://preactjs.com/guide/v10/api-reference#h--createelement
				return h(
					"button",
					{
						className: "btn",
						onClick: () => alert(`See the detail on ${cell}`),
					},
					"Detail",
				);
			},
		},
	];

	// https://gridjs.io/docs/config/server
	const server = {
		url: "https://swapi.dev/api/people",
		total: (data) => data.count,
		then: (data) => data.results.map(people => ([
				people.name,
				people.gender,
				people.eye_color,
				people.url
			])
		),
	};

	// https://gridjs.io/docs/config/pagination
	const pagination = {
		enabled: true,
		limit: 10,
		server: {
			// https://github.com/grid-js/gridjs/issues/84
			url: (prev, page) => {
				return `${prev}${prev.includes("?")
					? "&"
					: "?"}page=${page + 1}`
			},
		},
	};

	// https://gridjs.io/docs/config/search
	const search = {
		enabled: true,
		server: {
			url: (prev, keyword) => `${prev}?search=${keyword}`,
		},
	};
</script>

<Grid
	{search}
	{columns}
	{server}
	{pagination}
	{language}
	{className}
	{style}
/>

<style global>
	@import "https://cdn.jsdelivr.net/npm/gridjs/dist/theme/mermaid.min.css";

	:root {
		--primary-color: #ffe300;
		--error-color: red;
	}

	:global(.btn) {
		background-color: var(--primary-color);
		padding: 0.5rem 1rem;
		border-width: 0;
		border-radius: 0.5rem;
		cursor: pointer;
	}
	
	:global(.error) {
		color: var(--error-color);
	}
</style>