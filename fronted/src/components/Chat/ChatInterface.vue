<script setup lang="ts">
import Default from "@/components/Chat/Default.vue";
import SwitchTheme from "@/components/SwitchTheme.vue";
import { useChatState } from "@/stores/useChatState";
import {
	File,
	Image,
	Monitor,
	Paperclip,
	Phone,
	Send,
	Smile,
	User,
	Video,
} from "lucide-vue-next";

const chatState = useChatState();

const notice = [
	{
		id: 1,
		author: "1号人物",
		content: "你猜说了什么",
	},
	{
		id: 2,
		author: "2号人物",
		content: "你猜说了什么",
	},
];
</script>

<template>
	<div class="w-full flex bg-gray-50 dark:bg-gray-900">
		<template v-if="!chatState.isChatting">
			<Default />
		</template>

		<template v-else>
			<main class="flex-1 flex flex-col h-full">
				<div
					class="flex items-center justify-between p-4 border-b border-gray-200 dark:border-gray-700"
				>
					<div class="flex items-center space-x-4">
						<div class="flex items-center">
							<div
								class="w-10 h-10 bg-blue-100 dark:bg-gray-700 rounded-full flex items-center justify-center"
							>
								<User
									class="text-blue-500 dark:text-blue-400"
									:size="20"
								/>
							</div>
							<div class="ml-3">
								<h2
									class="font-semibold text-gray-800 dark:text-gray-200"
								>
									{{
										chatState
										.currentSessionDetail
										?.name
									}}
								</h2>
								<p
									v-if='
										chatState
										.currentSessionDetail
										?.type ===
										"group"
									'
									class="text-sm text-gray-500 dark:text-gray-400"
								>
									1人
								</p>
							</div>
						</div>
					</div>
					<div class="flex items-center space-x-4">
						<Phone
							class="text-gray-600 dark:text-gray-300 hover:text-blue-500 cursor-pointer"
						/>
						<Video
							class="text-gray-600 dark:text-gray-300 hover:text-blue-500 cursor-pointer"
						/>
						<Monitor
							class="text-gray-600 dark:text-gray-300 hover:text-blue-500 cursor-pointer"
						/>
					</div>
				</div>

				<div class="flex-1 overflow-y-auto p-4 space-y-4"></div>

				<div class="p-4 border-t border-gray-200 dark:border-gray-700">
					<div class="flex items-center space-x-2">
						<div class="relative group">
							<button
								class="p-2 text-gray-500 hover:text-gray-700 dark:hover:text-gray-300"
							>
								<Paperclip :size="20" />
							</button>
							<div
								class="absolute bottom-full left-0 hidden group-hover:block bg-white dark:bg-gray-800 shadow-lg rounded-lg p-2 min-w-[120px]"
							>
								<button
									class="cursor-pointer w-full flex items-center space-x-2 p-2 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md dark:text-white"
								>
									<Image :size="16" />
									<span>图片</span>
								</button>
								<button
									class="cursor-pointer w-full flex items-center space-x-2 p-2 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md dark:text-white"
								>
									<File :size="16" />
									<span>文档</span>
								</button>
							</div>
						</div>
						<button
							class="p-2 text-gray-500 hover:text-gray-700 dark:hover:text-gray-300"
						>
							<Smile :size="20" />
						</button>
						<input
							type="text"
							placeholder="说些什么..."
							class="flex-1 p-3 rounded-xl bg-white focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-800 border border-gray-200 dark:border-gray-700 focus:outline-none focus:ring-1 dark:text-white"
						/>
						<button
							class="bg-blue-500 hover:bg-blue-600 cursor-pointer p-3 text-white rounded-xl transition-colors"
						>
							<Send :size="20" />
						</button>
					</div>
				</div>
			</main>

			<aside
				class="w-64 bg-white dark:bg-gray-800 border-l border-gray-200 dark:border-gray-700 h-full"
				v-if='chatState.currentSessionDetail?.type == "group"'
			>
				<div class="h-1/3">
					<div
						class="p-4 border-b border-gray-200 dark:border-gray-700 flex justify-between items-center"
					>
						<h2
							class="text-lg font-semibold text-gray-800 dark:text-gray-200"
						>
							公告
						</h2>
					</div>
					<div class="p-2 overflow-y-auto space-y-4">
						<div
							v-for="post in notice"
							:key="post.id"
							class="p-3 bg-gray-50 dark:bg-gray-900 rounded-lg"
						>
							<div class="flex items-center mb-2">
								<div
									class="w-8 h-8 bg-blue-100 dark:bg-gray-700 rounded-full flex items-center justify-center"
								>
									<User
										class="text-blue-500 dark:text-blue-400"
										:size="16"
									/>
								</div>
								<span
									class="ml-2 font-medium text-gray-700 dark:text-gray-300"
								>
									{{
										post
										.author
									}}
								</span>
							</div>
							<p class="text-sm text-gray-600 dark:text-gray-400">
								{{
									post
									.content
								}}
							</p>
						</div>
					</div>
				</div>
				<div>
					<div
						class="p-4 border-b border-gray-200 dark:border-gray-700"
					>
						<h2
							class="text-lg font-semibold text-gray-800 dark:text-gray-200"
						>
							成员
						</h2>
					</div>
				</div>
			</aside>
		</template>
	</div>
</template>
