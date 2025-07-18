<script setup lang="ts">
import { ref } from "vue";
import { ChevronDown, HelpCircle } from "lucide-vue-next";

const appName = import.meta.env.VITE_APP_NAME;
const faqItems = ref([
	{
		id: 1,
		question: "如何添加好友？",
		answer:
			'在应用主界面点击右上角的"添加好友"按钮，输入对方的用户名或ID发送请求，对方接受后即可成为好友。',
		open: false,
	},
	{
		id: 2,
		question: "如何创建群聊？",
		answer:
			'在聊天界面点击右上角的"+"按钮，选择"新建群组"，添加成员并设置群名称后即可创建群聊。',
		open: false,
	},
	{
		id: 3,
		question: "消息发送失败怎么办？",
		answer:
			"首先检查网络连接是否正常。如果网络正常，尝试重启应用或检查应用是否为最新版本。问题仍未解决请联系管理员。",
		open: false,
	},
	{
		id: 4,
		question: "如何更改个人资料？",
		answer:
			"进入设置 > 账号设置 > 个人资料，您可以在此修改头像、昵称、个性签名等信息。",
		open: false,
	},
	{
		id: 5,
		question: "如何开启暗黑模式？",
		answer:
			'在设置 > 外观设置 > 主题设置中，您可以选择"深色模式"来启用暗黑主题。',
		open: false,
	},
]);

const toggleFAQ = (id: number) => {
	const item = faqItems.value.find((item) => item.id === id);
	if (item) {
		item.open = !item.open;
	}
};
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6">
		<div class="max-w-3xl mx-auto">
			<div class="flex items-center mb-8">
				<div
					class="p-3 bg-blue-100 dark:bg-blue-900/30 rounded-lg mr-4"
				>
					<HelpCircle class="w-6 h-6 text-blue-500" />
				</div>
				<div>
					<h1
						class="text-2xl font-bold text-gray-900 dark:text-white"
					>
						常见问题
					</h1>
					<p class="text-gray-500 dark:text-gray-400 mt-1">
						在这里您可以找到关于使用{{ appName }}的常见问题解答
					</p>
				</div>
			</div>

			<div class="space-y-4">
				<div
					v-for="item in faqItems"
					:key="item.id"
					class="bg-white dark:bg-gray-800 rounded-xl shadow-sm overflow-hidden"
				>
					<button
						@click="toggleFAQ(item.id)"
						class="w-full flex items-center justify-between p-5 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
					>
						<div class="flex items-center">
							<div
								class="w-6 h-6 bg-blue-100 dark:bg-blue-900/30 rounded-full flex items-center justify-center mr-3"
							>
								<span class="text-blue-500 font-medium">Q</span>
							</div>
							<span
								class="font-medium text-left text-gray-900 dark:text-white"
							>{{ item.question }}</span>
						</div>
						<ChevronDown
							class="w-5 h-5 text-gray-400 transition-transform duration-300"
							:class='
								{
									"rotate-180":
										item.open,
								}
							'
						/>
					</button>

					<div
						v-if="item.open"
						class="px-5 pb-5 pt-2 border-t border-gray-200 dark:border-gray-700 animate-fade-in"
					>
						<div class="flex items-start">
							<div
								class="w-6 h-6 bg-green-100 dark:bg-green-900/30 rounded-full flex items-center justify-center mr-3 mt-1 flex-shrink-0"
							>
								<span class="text-green-500 font-medium"
								>A</span>
							</div>
							<p class="text-gray-600 dark:text-gray-300">
								{{
									item
									.answer
								}}
							</p>
						</div>
					</div>
				</div>
			</div>

			<div
				class="mt-10 pt-6 border-t border-gray-200 dark:border-gray-700"
			>
				<h2
					class="text-lg font-semibold text-gray-900 dark:text-white mb-4"
				>
					未找到您的问题？
				</h2>
				<p class="text-gray-600 dark:text-gray-300 mb-4">
					如果您的问题未在列表中，请通过意见反馈或联系我们页面寻求帮助。
				</p>
				<div class="flex space-x-3">
					<button
						class="px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded-lg transition-colors"
					>
						意见反馈
					</button>
					<button
						class="px-4 py-2 bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 text-gray-800 dark:text-gray-200 rounded-lg transition-colors"
					>
						联系我们
					</button>
				</div>
			</div>
		</div>
	</div>
</template>
