<script setup lang="ts">
import {
	Code,
	File,
	Github,
	Globe,
	Heart,
	Laptop,
	Lock,
	MessageCircle,
	Moon,
	Send,
	Server,
	Smartphone,
} from "lucide-vue-next";

const features = [
	{ id: 1, text: "端到端加密", icon: Lock },
	{ id: 2, text: "跨平台支持", icon: Smartphone },
	{ id: 3, text: "开源免费", icon: Code },
	{ id: 4, text: "暗黑模式", icon: Moon },
	{ id: 5, text: "实时通讯", icon: Send },
	{ id: 6, text: "文件传输", icon: File },
];

const teamMembers = [
	{ id: 1, name: "ImAur", role: "一位Dev", avatar: "IA" },
];

const appName = import.meta.env.VITE_APP_NAME;
const appDescription = import.meta.env.VITE_APP_DESCRIPTIO;
const appVersion = import.meta.env.VITE_APP_VERSION;
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6">
		<div class="max-w-4xl mx-auto">
			<div class="text-center mb-12">
				<div
					class="mx-auto w-24 h-24 bg-gradient-to-br from-blue-500 to-indigo-600 rounded-2xl flex items-center justify-center mb-6"
				>
					<MessageCircle class="w-12 h-12 text-white" />
				</div>
				<h1
					class="text-3xl font-bold text-gray-900 dark:text-white mb-2"
				>
					{{ appName }}
				</h1>
				<p class="text-lg text-gray-600 dark:text-gray-300 mb-4">
					{{ appDescription }}
				</p>
				<div class="text-gray-500 dark:text-gray-400">
					版本 {{ appVersion }}
				</div>
			</div>

			<div
				class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6 mb-8"
			>
				<h2
					class="text-xl font-semibold text-gray-900 dark:text-white mb-6"
				>
					应用特色
				</h2>

				<div class="grid grid-cols-2 md:grid-cols-3 gap-4">
					<div
						v-for="feature in features"
						:key="feature.id"
						class="p-4 bg-gray-50 dark:bg-gray-700 rounded-lg flex items-center"
					>
						<div
							class="mr-3 p-2 bg-blue-100 dark:bg-blue-900/20 rounded-lg"
						>
							<component
								:is="feature.icon"
								class="w-5 h-5 text-blue-500"
							/>
						</div>
						<span
							class="font-medium text-gray-700 dark:text-gray-300"
						>{{ feature.text }}</span>
					</div>
				</div>
			</div>

			<div class="grid grid-cols-1 md:grid-cols-2 gap-8 mb-8">
				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<h2
						class="text-xl font-semibold text-gray-900 dark:text-white mb-6"
					>
						技术栈
					</h2>

					<div class="space-y-4">
						<div class="flex items-center">
							<div
								class="w-10 h-10 bg-blue-100 dark:bg-blue-900/20 rounded-lg flex items-center justify-center mr-3"
							>
								<Globe class="w-5 h-5 text-blue-500" />
							</div>
							<div>
								<h3
									class="font-medium text-gray-900 dark:text-white"
								>
									前端
								</h3>
								<p
									class="text-sm text-gray-500 dark:text-gray-400"
								>
									Vue, Vite, TypeScript, Tailwind CSS
								</p>
							</div>
						</div>

						<div class="flex items-center">
							<div
								class="w-10 h-10 bg-green-100 dark:bg-green-900/20 rounded-lg flex items-center justify-center mr-3"
							>
								<Server class="w-5 h-5 text-green-500" />
							</div>
							<div>
								<h3
									class="font-medium text-gray-900 dark:text-white"
								>
									后端
								</h3>
								<p
									class="text-sm text-gray-500 dark:text-gray-400"
								>
									Spring Boot, PostgreSQL, 各种NoSQL
								</p>
							</div>
						</div>

						<div class="flex items-center">
							<div
								class="w-10 h-10 bg-purple-100 dark:bg-purple-900/20 rounded-lg flex items-center justify-center mr-3"
							>
								<Laptop class="w-5 h-5 text-yellow-500" />
							</div>
							<div>
								<h3
									class="font-medium text-gray-900 dark:text-white"
								>
									电脑端
								</h3>
								<p
									class="text-sm text-gray-500 dark:text-gray-400"
								>
									Tauri, Vue
								</p>
							</div>
						</div>

						<div class="flex items-center">
							<div
								class="w-10 h-10 bg-purple-100 dark:bg-purple-900/20 rounded-lg flex items-center justify-center mr-3"
							>
								<Smartphone class="w-5 h-5 text-purple-500" />
							</div>
							<div>
								<h3
									class="font-medium text-gray-900 dark:text-white"
								>
									移动端
								</h3>
								<p
									class="text-sm text-gray-500 dark:text-gray-400"
								>
									Expo, React
								</p>
							</div>
						</div>
					</div>
				</div>

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<h2
						class="text-xl font-semibold text-gray-900 dark:text-white mb-6"
					>
						开源项目
					</h2>

					<div class="flex items-center mb-4">
						<Github
							class="w-6 h-6 text-gray-900 dark:text-white mr-3"
						/>
						<h3 class="font-medium text-gray-900 dark:text-white">
							GitHub 仓库
						</h3>
					</div>

					<p class="text-gray-600 dark:text-gray-300 mb-4">
						BlockIM 是完全开源的，欢迎开发者贡献代码或提出改进建议。
					</p>

					<a
						href="https://github.com/BlocksFunction/BlockIM-Web"
						target="_blank"
						class="inline-flex items-center px-4 py-2 bg-gray-800 hover:bg-gray-900 text-white rounded-lg transition-colors"
					>
						<Github class="w-4 h-4 mr-2" />
						访问 GitHub
					</a>

					<div
						class="mt-6 pt-4 border-t border-gray-200 dark:border-gray-700"
					>
						<div class="flex items-center">
							<Heart class="w-5 h-5 text-red-500 mr-2" />
							<span class="text-gray-600 dark:text-gray-300"
							>感谢所有贡献者！</span>
						</div>
					</div>
				</div>
			</div>

			<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
				<h2
					class="text-xl font-semibold text-gray-900 dark:text-white mb-6"
				>
					团队介绍
				</h2>

				<div
					class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4"
				>
					<div
						v-for="member in teamMembers"
						:key="member.id"
						class="text-center p-4 border border-gray-200 dark:border-gray-700 rounded-lg"
					>
						<div
							class="w-16 h-16 bg-gradient-to-br from-blue-400 to-indigo-500 rounded-full flex items-center justify-center mx-auto mb-4"
						>
							<span class="text-white font-medium text-xl">{{
								member.avatar
							}}</span>
						</div>
						<h3 class="font-semibold text-gray-900 dark:text-white">
							{{ member.name }}
						</h3>
						<p class="text-sm text-gray-500 dark:text-gray-400">
							{{ member.role }}
						</p>
					</div>
				</div>
			</div>

			<div
				class="mt-8 text-center text-sm text-gray-500 dark:text-gray-400"
			>
				<p>© 2025 BlockIM. 保留所有权利。</p>
				<p class="mt-2">遵循 AGPL 开源协议</p>
			</div>
		</div>
	</div>
</template>
