<script setup lang="ts">
import { useChatState } from "@/stores/useChatState.ts";
import type { PrivateChat } from "@/stores/useChatState.ts";
import { UserPlus } from "lucide-vue-next";

const chatState = useChatState();

const friends = [
	{ id: 1, name: "1号人物", avatar: "A", online: true },
	{ id: 2, name: "2号人物", avatar: "B", online: true },
	{ id: 3, name: "3号人物", avatar: "C", online: false },
];

const openPrivateChat = (friend: (typeof friends)[number]) => {
	const session: PrivateChat = {
		id: `private_${friend.id}`,
		type: "private",
		unread: 0,
		friendId: friend.id.toString(),
		friendName: friend.name,
		friendNickName: friend.name,
		friendAvatar: friend.avatar,
		lastMessage: undefined,
		lastTime: undefined,
	};

	chatState.openSession(session);
};
</script>

<template>
	<aside
		class="w-64 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 flex flex-col flex-none"
	>
		<div class="p-4 border-b border-gray-200 dark:border-gray-700">
			<button
				class="mt-2 w-full flex items-center space-x-2 p-2 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg"
			>
				<UserPlus
					:size="18"
					class="text-blue-500"
				/>
				<span class="text-gray-700 dark:text-gray-300"> 加个好友 </span>
			</button>
		</div>

		<div class="flex-1 overflow-y-auto p-2">
			<div
				v-for="friend in friends"
				:key="friend.id"
				@click="openPrivateChat(friend)"
				class="flex items-center p-3 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg cursor-pointer"
			>
				<div class="relative">
					<div
						class="w-10 h-10 bg-blue-100 dark:bg-gray-700 rounded-full flex items-center justify-center"
					>
						<span class="text-gray-600 dark:text-gray-300">
							{{ friend.avatar }}
						</span>
					</div>
					<div
						v-if="friend.online"
						class="absolute bottom-0 right-0 w-3 h-3 bg-green-500 rounded-full border-2 border-white dark:border-gray-800"
					/>
				</div>
				<div class="ml-3">
					<p class="font-medium text-gray-700 dark:text-gray-200">
						{{ friend.name }}
					</p>
					<p class="text-sm text-gray-500 dark:text-gray-400">
						{{
							friend.online
							? "在线"
							: "离线"
						}}
					</p>
				</div>
			</div>
		</div>
	</aside>
</template>
