<script lang="ts">
	import { FilePlus, Info, Settings, UserRound } from '@lucide/svelte';
	import * as Sidebar from '$lib/components/ui/sidebar';
	import * as Avatar from '$lib/components/ui/avatar'
	import { m } from '$lib/paraglide/messages';

	const items = [
		{
			title: m.add_instance(),
			url: '/add-instance',
			icon: FilePlus
		},
		{
			title: m.settings(),
			url: '/settings',
			icon: Settings
		},
		{
			title: m.help(),
			url: '/help',
			icon: Info
		}
	];
</script>

<Sidebar.Root>
	<Sidebar.Header>
		<div class="flex gap-2 items-center p-1">
			<Avatar.Root>
				<Avatar.Image src="" alt="@shadcn" />
				<Avatar.Fallback><UserRound /></Avatar.Fallback>
			</Avatar.Root>
			Admin
		</div>
		<hr />
	</Sidebar.Header>
	<Sidebar.Content>
		<Sidebar.Group>
			<Sidebar.GroupContent>
				<Sidebar.Menu>
					{#each items as item (item.title)}
						<Sidebar.MenuItem>
							<Sidebar.MenuButton>
								{#snippet child({ props })}
									<a href={item.url} {...props}>
										<item.icon />
										<span>{item.title}</span>
									</a>
								{/snippet}
							</Sidebar.MenuButton>
						</Sidebar.MenuItem>
					{/each}
				</Sidebar.Menu>
			</Sidebar.GroupContent>
		</Sidebar.Group>
	</Sidebar.Content>
</Sidebar.Root>
