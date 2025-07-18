<script setup lang="ts">
import ChatInterface from "@/components/Chat/ChatInterface.vue";
import FriendList from "@/components/Chat/FriendList.vue";
import ConfigInterface from "@/components/Config/ConfigInterface.vue";
import FriendInterface from "@/components/Friend/FriendInterface.vue";
import {
	MessageCircle,
	MessageCircleMore,
	Settings,
	Users,
} from "lucide-vue-next";
import { ref } from "vue";

enum Page {
	Chat = "chat",
	Friend = "friend",
	Config = "config",
}

const currentPage = ref<Page>(Page.Chat);
</script>

<template>
	<div
		class="min-h-screen bg-gray-50 dark:bg-dark-900 transition-colors duration-200"
	>
		<div class="flex h-screen">
			<aside
				class="w-20 bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 flex flex-col flex-none items-center py-4 space-y-4"
			>
				<button
					class="p-3 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-xl transition-colors"
					:class='
						currentPage === Page.Chat
						? "bg-gray-200 dark:bg-gray-600"
						: ""
					'
					@click="currentPage = Page.Chat"
				>
					<MessageCircle class="text-gray-600 dark:text-gray-300" />
				</button>
				<button
					class="p-3 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-xl transition-colors"
					:class='
						currentPage === Page.Friend
						? "bg-gray-200 dark:bg-gray-600"
						: ""
					'
					@click="currentPage = Page.Friend"
				>
					<Users
						class="text-gray-600 dark:text-gray-300"
						:size="24"
					/>
				</button>
				<button
					class="p-3 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-xl transition-colors"
				>
					<MessageCircleMore
						class="text-gray-600 dark:text-gray-300"
						:size="24"
					/>
				</button>
				<button
					class="p-3 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-xl transition-colors"
					:class='
						currentPage === Page.Config
						? "bg-gray-200 dark:bg-gray-600"
						: ""
					'
					@click="currentPage = Page.Config"
				>
					<Settings
						class="text-gray-600 dark:text-gray-300"
						:size="24"
					/>
				</button>
			</aside>
			<template v-if="currentPage === Page.Chat">
				<FriendList />
				<ChatInterface />
			</template>

			<template v-if="currentPage === Page.Friend">
				<FriendInterface />
			</template>

			<template v-if="currentPage === Page.Config">
				<ConfigInterface />
			</template>
		</div>
	</div>
</template>
