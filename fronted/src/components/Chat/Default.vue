<script setup lang="ts">
import { MessageCircle, UserPlus, Users } from "lucide-vue-next";
</script>

<template>
	<div
		class="flex-1 flex flex-col items-center justify-center p-8 text-center"
	>
		<div
			class="w-32 h-32 bg-gray-100 dark:bg-gray-800 rounded-full flex items-center justify-center mb-6"
		>
			<MessageCircle
				class="text-gray-400 dark:text-gray-500"
				:size="48"
			/>
		</div>
		<h2
			class="text-2xl font-semibold text-gray-700 dark:text-gray-300 mb-2"
		>
			欢迎使用BlockIm
		</h2>
		<p class="text-gray-500 dark:text-gray-400 max-w-md">
			从左侧选择好友开始聊天，或者创建新的群组与多人交流~
		</p>
		<div class="mt-6 flex gap-4">
			<button
				class="cursor-pointer px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors flex items-center"
			>
				<UserPlus class="mr-2" :size="18" />
				添加好友
			</button>
			<button
				class="cursor-pointer px-6 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 transition-colors flex items-center"
			>
				<Users class="mr-2" :size="18" />
				新建群聊
			</button>
		</div>
	</div>
</template>
