<script setup lang="ts">
import { ref } from "vue";
import {
	Check,
	Laptop,
	Monitor,
	Moon,
	Palette,
	Send,
	Sun,
} from "lucide-vue-next";
import useTheme from "@/stores/useTheme";

const theme = useTheme();

const colorThemes = ref([
	{ id: "blue", name: "蓝色", class: "bg-blue-500" },
	{ id: "green", name: "绿色", class: "bg-green-500" },
	{ id: "purple", name: "紫色", class: "bg-purple-500" },
	{ id: "red", name: "红色", class: "bg-red-500" },
	{ id: "orange", name: "橙色", class: "bg-orange-500" },
	{ id: "pink", name: "粉色", class: "bg-pink-500" },
]);

const themeOptions = ref([
	{ id: "light", name: "浅色模式", icon: Sun },
	{ id: "dark", name: "深色模式", icon: Moon },
	{ id: "system", name: "跟随系统", icon: Laptop },
]);
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6">
		<div class="max-w-3xl mx-auto">
			<div class="mb-8">
				<h1 class="text-2xl font-bold text-gray-900 dark:text-white">
					主题设置
				</h1>
				<p class="text-gray-500 dark:text-gray-400 mt-2">
					自定义应用的外观和主题
				</p>
			</div>

			<div
				class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6 mb-6"
			>
				<div class="flex items-start mb-6">
					<Palette class="w-6 h-6 mt-1 mr-3 text-blue-500" />
					<div class="flex-1">
						<h2
							class="text-lg font-semibold text-gray-900 dark:text-white"
						>
							主题模式
						</h2>
						<p
							class="text-sm text-gray-500 dark:text-gray-400 mt-1"
						>
							选择您喜欢的主题模式
						</p>
					</div>
				</div>

				<div class="grid grid-cols-1 md:grid-cols-3 gap-4">
					<button
						v-for="option in themeOptions"
						:key="option.id"
						@click="theme.setTheme(option.id as any)"
						class="p-4 border rounded-xl transition-all"
						:class='
							theme.getTheme === option.id
							? "border-blue-500 bg-blue-50 dark:bg-blue-900/20"
							: "border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700"
						'
					>
						<span class="flex flex-col items-center">
							<component
								:is="option.icon"
								class="w-8 h-8 mb-3"
								:class='
									theme
										.getTheme ===
										option.id
									? "text-blue-500"
									: "text-gray-700 dark:text-gray-300"
								'
							/>
							<span
								class="font-medium"
								:class='
									theme
										.getTheme ===
										option.id
									? "text-blue-500"
									: "text-gray-700 dark:text-gray-300"
								'
							>
								{{
									option
									.name
								}}
							</span>
						</span>
					</button>
				</div>
			</div>

			<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
				<div class="flex items-start mb-6">
					<Monitor class="w-6 h-6 mt-1 mr-3 text-blue-500" />
					<div class="flex-1">
						<h2
							class="text-lg font-semibold text-gray-900 dark:text-white"
						>
							主题预览
						</h2>
						<p
							class="text-sm text-gray-500 dark:text-gray-400 mt-1"
						>
							查看您的主题设置效果
						</p>
					</div>
				</div>

				<div class="flex flex-col md:flex-row gap-6">
					<div class="flex-1">
						<div
							class="bg-gray-50 dark:bg-gray-900 rounded-lg overflow-hidden border border-gray-200 dark:border-gray-700"
						>
							<div
								class="bg-white dark:bg-gray-800 p-3 border-b border-gray-200 dark:border-gray-700"
							>
								<div class="flex items-center">
									<div
										class="w-3 h-3 bg-red-500 rounded-full mr-2"
									>
									</div>
									<div
										class="w-3 h-3 bg-yellow-500 rounded-full mr-2"
									>
									</div>
									<div
										class="w-3 h-3 bg-green-500 rounded-full"
									>
									</div>
								</div>
							</div>
							<div class="p-4">
								<div class="flex items-center mb-4">
									<div class="relative">
										<div
											class="w-10 h-10 bg-blue-100 dark:bg-gray-700 rounded-full flex items-center justify-center"
										>
											<span
												class="text-gray-600 dark:text-gray-300"
											>
												IA
											</span>
										</div>
										<div
											class="absolute bottom-0 right-0 w-3 h-3 bg-green-500 rounded-full border-2 border-white dark:border-gray-800"
										/>
									</div>
									<div class="ml-3">
										<p
											class="font-medium text-gray-700 dark:text-gray-200"
										>
											ImAur
										</p>
										<p
											class="text-sm text-gray-500 dark:text-gray-400"
										>
											在线
										</p>
									</div>
								</div>

								<div class="space-y-3">
									<div class="flex justify-end">
										<div
											class="text-white rounded-xl rounded-br-none px-4 py-2 max-w-xs bg-blue-500"
										>
											你好，最近怎么样？
										</div>
									</div>

									<div class="flex">
										<div
											class="bg-gray-200 dark:bg-gray-700 rounded-xl rounded-bl-none px-4 py-2 max-w-xs dark:text-white"
										>
											我很好，谢谢！
										</div>
									</div>
								</div>

								<div
									class="mt-6 flex items-center border-t border-gray-200 dark:border-gray-700 pt-4"
								>
									<input
										type="text"
										placeholder="输入消息..."
										class="flex-1 p-2 bg-gray-100 dark:bg-gray-700 rounded-lg focus:outline-none dark:text-gray-300"
									/>
									<button
										class="ml-2 p-2 text-white rounded-lg bg-blue-500"
									>
										<Send class="w-5 h-5" />
									</button>
								</div>
							</div>
						</div>
						<p
							class="text-center text-sm text-gray-500 dark:text-gray-400 mt-2"
						>
							桌面端预览
						</p>
					</div>

					<div class="flex-1 max-w-xs">
						<div
							class="bg-gray-50 dark:bg-gray-900 rounded-lg overflow-hidden border border-gray-200 dark:border-gray-700"
						>
							<div
								class="bg-white dark:bg-gray-800 p-3 border-b border-gray-200 dark:border-gray-700 flex justify-between items-center"
							>
								<div
									class="text-sm font-medium text-gray-900 dark:text-white"
								>
									聊天
								</div>
								<div
									class="w-6 h-6 bg-gray-200 dark:bg-gray-700 rounded-full"
								>
								</div>
							</div>
							<div class="p-3">
								<div class="space-y-2">
									<div
										class="flex items-center p-2 rounded-lg bg-gray-100 dark:bg-gray-700"
									>
										<div
											class="w-8 h-8 bg-gray-300 dark:bg-gray-600 rounded-full flex items-center justify-center mr-2"
										>
											<span
												class="text-xs text-gray-600 dark:text-gray-300"
											>UN</span>
										</div>
										<div>
											<div
												class="text-sm font-medium text-gray-900 dark:text-white"
											>
												UNK
											</div>
											<div
												class="text-xs text-gray-500 dark:text-gray-400"
											>
												你好！
											</div>
										</div>
									</div>

									<div
										class="flex items-center p-2 rounded-lg bg-blue-100 dark:bg-blue-900/30"
									>
										<div
											class="w-8 h-8 rounded-full flex items-center justify-center mr-2 bg-blue-300 dark:bg-blue-700"
										>
											<span class="text-xs text-white"
											>IA</span>
										</div>
										<div>
											<div
												class="text-sm font-medium text-gray-900 dark:text-white"
											>
												ImAur
											</div>
											<div
												class="text-xs text-gray-500 dark:text-gray-400"
											>
												我不好，不谢！
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
						<p
							class="text-center text-sm text-gray-500 dark:text-gray-400 mt-2"
						>
							移动端预览
						</p>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
