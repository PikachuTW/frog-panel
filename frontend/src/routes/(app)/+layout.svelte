<script lang="ts">
	import '../../app.css';
	import { ModeWatcher } from "mode-watcher";
	import favicon from '$lib/assets/favicon.svg';

	import * as Sidebar from '$lib/components/ui/sidebar';
	import AppSidebar from '$lib/components/main/sidebar.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<Sidebar.Provider>
	<AppSidebar />
	<main>
		<ModeWatcher />
		{@render children?.()}
	</main>
</Sidebar.Provider>
