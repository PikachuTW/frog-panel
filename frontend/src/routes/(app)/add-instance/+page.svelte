<script>
	import { m } from '$lib/paraglide/messages';
	import { Label } from '$lib/components/ui/label/index.js';
	import * as RadioGroup from '$lib/components/ui/radio-group/index.js';
	let { data } = $props();

	const mcTypes = data.versionTypes;
	let modLoader = $state(mcTypes[0]);
	let selectedVersion = $state(null);

	// const mcVersions = $derived.by(async () => {
	// 	const res = await (await fetch(`/api/versions/${modLoader}`)).json();
	// 	return res;
	// });
</script>

<div class="px-8">
	<h1 class="py-6 text-3xl font-bold">{m.add_instance()}</h1>
	<div class="flex gap-4">
		<div>
			<h2 class="text-bold mb-4 text-xl capitalize">{m.mod_loader()}</h2>
			<RadioGroup.Root bind:value={modLoader}>
				{#each mcTypes as mcType}
					<div class="flex items-center space-x-2">
						<RadioGroup.Item value={mcType} id={mcType} />
						<Label for={mcType} class="capitalize">{mcType}</Label>
					</div>
				{/each}
			</RadioGroup.Root>
		</div>
	</div>
</div>
