<script setup lang="ts">
import { Globe, Mail, MessageCircle, Phone, Users } from "lucide-vue-next";

const contactMethods = [
	{
		id: "email",
		icon: Mail,
		title: "电子邮件",
		description: "给管理员发送邮件",
		details: "IDK",
		action: "IDK",
	},
	{
		id: "phone",
		icon: Phone,
		title: "管理员电话",
		description: "你猜什么时候工作",
		details: "你猜电话号码是什么",
		action: "你猜电话号码是什么",
	},
	{
		id: "forum",
		icon: MessageCircle,
		title: "社区论坛",
		description: "加入我们的用户社区",
		details: "你猜URL是什么",
		action: "你猜URL是什么",
	},
];

const socialMedia = [
	{
		id: "weibo",
		name: "你猜是什么平台",
		handle: "@BlockIM",
		icon: Globe,
		color: "bg-red-500",
		url: "你猜URL是什么",
	},
	{
		id: "wechat",
		name: "你猜是什么平台",
		handle: "BlockIM",
		icon: Users,
		color: "bg-green-500",
		url: "#",
	},
];
const appName = import.meta.env.VITE_APP_NAME;
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6">
		<div class="max-w-3xl mx-auto">
			<div class="flex items-center mb-8">
				<div
					class="p-3 bg-amber-100 dark:bg-amber-900/30 rounded-lg mr-4"
				>
					<Phone class="w-6 h-6 text-amber-500" />
				</div>
				<div>
					<h1
						class="text-2xl font-bold text-gray-900 dark:text-white"
					>
						联系我们
					</h1>
					<p class="text-gray-500 dark:text-gray-400 mt-1">
						获取{{ appName }}的支持和帮助
					</p>
				</div>
			</div>

			<div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
				<div
					class="bg-gradient-to-br from-blue-500 to-indigo-600 rounded-xl p-6 text-white"
				>
					<h2 class="text-xl font-bold mb-2">管理员咨询</h2>
					<p class="mb-4 opacity-90">我们随时为您提供帮助</p>
					<div class="flex items-center">
						<div class="mr-3 bg-white/20 p-2 rounded-full">
							<Phone class="w-5 h-5" />
						</div>
						<div>
							<p class="text-sm opacity-80">管理员电话</p>
							<p class="font-semibold text-lg">114514</p>
						</div>
					</div>
				</div>

				<div
					class="bg-gradient-to-br from-purple-500 to-pink-500 rounded-xl p-6 text-white"
				>
					<h2 class="text-xl font-bold mb-2">反馈邮箱</h2>
					<p class="mb-4 opacity-90">发送您的建议和反馈</p>
					<div class="flex items-center">
						<div class="mr-3 bg-white/20 p-2 rounded-full">
							<Mail class="w-5 h-5" />
						</div>
						<div>
							<p class="text-sm opacity-80">电子邮箱</p>
							<p class="font-semibold text-lg">
								IDK
							</p>
						</div>
					</div>
				</div>
			</div>

			<div
				class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6 mb-8"
			>
				<h2
					class="text-xl font-semibold text-gray-900 dark:text-white mb-6"
				>
					联系我们
				</h2>

				<div class="space-y-6">
					<div
						v-for="method in contactMethods"
						:key="method.id"
						class="flex items-start p-4 border border-gray-200 dark:border-gray-700 rounded-lg hover:border-blue-500 transition-colors"
					>
						<div
							class="p-3 bg-blue-100 dark:bg-blue-900/20 rounded-lg mr-4 flex-shrink-0"
						>
							<component
								:is="method.icon"
								class="w-5 h-5 text-blue-500"
							/>
						</div>
						<div class="flex-1 min-w-0">
							<h3
								class="font-semibold text-gray-900 dark:text-white"
							>
								{{
									method
									.title
								}}
							</h3>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								{{
									method
									.description
								}}
							</p>
							<p
								class="mt-2 font-medium text-gray-500 dark:text-gray-400"
							>
								{{
									method
									.details
								}}
							</p>
						</div>
						<a
							:href="method.action"
							class="ml-4 px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded-lg transition-colors whitespace-nowrap"
						>
							联系
						</a>
					</div>
				</div>
			</div>

			<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
				<h2
					class="text-xl font-semibold text-gray-900 dark:text-white mb-6"
				>
					社交媒体
				</h2>

				<div class="grid grid-cols-1 md:grid-cols-2 gap-4">
					<a
						v-for="platform in socialMedia"
						:key="platform.id"
						:href="platform.url"
						class="p-4 rounded-lg border border-gray-200 dark:border-gray-700 hover:shadow-md transition-all"
					>
						<div class="flex items-center">
							<div class="mr-4">
								<div
									:class='
										[
											platform
												.color,
											"w-12 h-12 rounded-lg flex items-center justify-center",
										]
									'
								>
									<component
										:is="
											platform
											.icon
										"
										class="w-6 h-6 text-white"
									/>
								</div>
							</div>
							<div>
								<h3
									class="font-semibold text-gray-900 dark:text-white"
								>
									{{
										platform
										.name
									}}
								</h3>
								<p
									class="text-sm text-gray-500 dark:text-gray-400"
								>
									{{
										platform
										.handle
									}}
								</p>
							</div>
						</div>
					</a>
				</div>
			</div>
		</div>
	</div>
</template>
