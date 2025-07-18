<script setup lang="ts">
import { ref } from "vue";
import {
	Eye,
	EyeOff,
	Globe,
	Lock,
	Mail,
	Phone,
	Search,
	User,
	Users,
} from "lucide-vue-next";

const privacySettings = ref({
	profileVisibility: "public",
	lastSeen: "everyone",
	readReceipts: true,
	blockedUsers: [
		{ id: 1, name: "李四", username: "lisi" },
		{ id: 2, name: "王五", username: "wangwu" },
	],
	dataSharing: {
		analytics: true,
		personalizedAds: false,
		improveApp: true,
	},
	contactPermissions: {
		email: true,
		phone: false,
		findByPhone: true,
	},
});

const newBlockedUser = ref("");
const addBlockedUser = () => {
	if (newBlockedUser.value.trim()) {
		const newUser = {
			id: privacySettings.value.blockedUsers.length + 1,
			name: newBlockedUser.value,
			username: newBlockedUser.value.toLowerCase().replace(/\s+/g, ""),
		};
		privacySettings.value.blockedUsers.push(newUser);
		newBlockedUser.value = "";
	}
};

const removeBlockedUser = (id: number) => {
	privacySettings.value.blockedUsers = privacySettings.value.blockedUsers
		.filter((user) => user.id !== id);
};
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6">
		<div class="max-w-3xl mx-auto">
			<div class="mb-8">
				<h1 class="text-2xl font-bold text-gray-900 dark:text-white">
					隐私设置
				</h1>
				<p class="text-gray-500 dark:text-gray-400 mt-2">
					控制谁可以看到您的信息以及如何与他人互动
				</p>
			</div>

			<div class="space-y-6">
				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-start">
						<Globe class="w-6 h-6 text-blue-500 mt-1 mr-3" />
						<div class="flex-1">
							<h2
								class="text-lg font-semibold text-gray-900 dark:text-white"
							>
								个人资料可见性
							</h2>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								谁可以查看您的个人资料信息
							</p>
						</div>
					</div>

					<div class="mt-6 space-y-4">
						<div
							class="flex items-center justify-between p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg"
						>
							<div class="flex items-center">
								<User class="w-5 h-5 text-gray-500 mr-3" />
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>个人资料可见性</span>
							</div>
							<select
								v-model="
									privacySettings
									.profileVisibility
								"
								class="px-3 py-1.5 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							>
								<option value="public">所有人</option>
								<option value="contacts">仅联系人</option>
								<option value="nobody">仅自己</option>
							</select>
						</div>

						<div
							class="flex items-center justify-between p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg"
						>
							<div class="flex items-center">
								<Eye class="w-5 h-5 text-gray-500 mr-3" />
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>最后上线时间</span>
							</div>
							<select
								v-model="
									privacySettings
									.lastSeen
								"
								class="px-3 py-1.5 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							>
								<option value="everyone">所有人</option>
								<option value="contacts">仅联系人</option>
								<option value="nobody">无人</option>
							</select>
						</div>

						<div
							class="flex items-center justify-between p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg"
						>
							<div class="flex items-center">
								<Check class="w-5 h-5 text-gray-500 mr-3" />
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>已读回执</span>
							</div>
							<button
								@click="
									privacySettings
									.readReceipts =
										!privacySettings
											.readReceipts
								"
								class="relative inline-flex items-center h-5 rounded-full w-9 bg-gray-200 dark:bg-gray-700 transition-colors"
								:class='
									privacySettings
										.readReceipts
									? "bg-green-500"
									: "bg-gray-300"
								'
							>
								<span
									class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
									:class='
										privacySettings
											.readReceipts
										? "translate-x-5"
										: "translate-x-1"
									'
								/>
							</button>
						</div>
					</div>
				</div>

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-start">
						<Users class="w-6 h-6 text-blue-500 mt-1 mr-3" />
						<div class="flex-1">
							<h2
								class="text-lg font-semibold text-gray-900 dark:text-white"
							>
								联系人权限
							</h2>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								控制谁可以联系您以及如何找到您
							</p>
						</div>
					</div>

					<div class="mt-6 space-y-4">
						<div
							class="flex items-center justify-between p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg"
						>
							<div class="flex items-center">
								<Mail class="w-5 h-5 text-gray-500 mr-3" />
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>通过电子邮件联系我</span>
							</div>
							<button
								@click="
									privacySettings
									.contactPermissions
									.email =
										!privacySettings
											.contactPermissions
											.email
								"
								class="relative inline-flex items-center h-5 rounded-full w-9 bg-gray-200 dark:bg-gray-700 transition-colors"
								:class='
									privacySettings
										.contactPermissions
										.email
									? "bg-green-500"
									: "bg-gray-300"
								'
							>
								<span
									class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
									:class='
										privacySettings
											.contactPermissions
											.email
										? "translate-x-5"
										: "translate-x-1"
									'
								/>
							</button>
						</div>

						<div
							class="flex items-center justify-between p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg"
						>
							<div class="flex items-center">
								<Phone class="w-5 h-5 text-gray-500 mr-3" />
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>通过电话号码联系我</span>
							</div>
							<button
								@click="
									privacySettings
									.contactPermissions
									.phone =
										!privacySettings
											.contactPermissions
											.phone
								"
								class="relative inline-flex items-center h-5 rounded-full w-9 bg-gray-200 dark:bg-gray-700 transition-colors"
								:class='
									privacySettings
										.contactPermissions
										.phone
									? "bg-green-500"
									: "bg-gray-300"
								'
							>
								<span
									class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
									:class='
										privacySettings
											.contactPermissions
											.phone
										? "translate-x-5"
										: "translate-x-1"
									'
								/>
							</button>
						</div>

						<div
							class="flex items-center justify-between p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg"
						>
							<div class="flex items-center">
								<Search class="w-5 h-5 text-gray-500 mr-3" />
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>通过电话号码找到我</span>
							</div>
							<button
								@click="
									privacySettings
									.contactPermissions
									.findByPhone =
										!privacySettings
											.contactPermissions
											.findByPhone
								"
								class="relative inline-flex items-center h-5 rounded-full w-9 bg-gray-200 dark:bg-gray-700 transition-colors"
								:class='
									privacySettings
										.contactPermissions
										.findByPhone
									? "bg-green-500"
									: "bg-gray-300"
								'
							>
								<span
									class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
									:class='
										privacySettings
											.contactPermissions
											.findByPhone
										? "translate-x-5"
										: "translate-x-1"
									'
								/>
							</button>
						</div>
					</div>
				</div>

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-start">
						<Lock class="w-6 h-6 text-blue-500 mt-1 mr-3" />
						<div class="flex-1">
							<h2
								class="text-lg font-semibold text-gray-900 dark:text-white"
							>
								数据共享
							</h2>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								控制您与我们共享的数据
							</p>
						</div>
					</div>

					<div class="mt-6 space-y-4">
						<div
							class="flex items-center justify-between p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg"
						>
							<div class="flex items-center">
								<div
									class="w-5 h-5 bg-blue-100 dark:bg-blue-900/30 rounded-full flex items-center justify-center mr-3"
								>
									<span class="text-xs text-blue-500">A</span>
								</div>
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>共享分析数据</span>
							</div>
							<button
								@click="
									privacySettings
									.dataSharing
									.analytics =
										!privacySettings
											.dataSharing
											.analytics
								"
								class="relative inline-flex items-center h-5 rounded-full w-9 bg-gray-200 dark:bg-gray-700 transition-colors"
								:class='
									privacySettings
										.dataSharing
										.analytics
									? "bg-green-500"
									: "bg-gray-300"
								'
							>
								<span
									class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
									:class='
										privacySettings
											.dataSharing
											.analytics
										? "translate-x-5"
										: "translate-x-1"
									'
								/>
							</button>
						</div>

						<div
							class="flex items-center justify-between p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg"
						>
							<div class="flex items-center">
								<div
									class="w-5 h-5 bg-purple-100 dark:bg-purple-900/30 rounded-full flex items-center justify-center mr-3"
								>
									<span class="text-xs text-purple-500"
									>P</span>
								</div>
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>个性化广告</span>
							</div>
							<button
								@click="
									privacySettings
									.dataSharing
									.personalizedAds =
										!privacySettings
											.dataSharing
											.personalizedAds
								"
								class="relative inline-flex items-center h-5 rounded-full w-9 bg-gray-200 dark:bg-gray-700 transition-colors"
								:class='
									privacySettings
										.dataSharing
										.personalizedAds
									? "bg-green-500"
									: "bg-gray-300"
								'
							>
								<span
									class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
									:class='
										privacySettings
											.dataSharing
											.personalizedAds
										? "translate-x-5"
										: "translate-x-1"
									'
								/>
							</button>
						</div>

						<div
							class="flex items-center justify-between p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg"
						>
							<div class="flex items-center">
								<div
									class="w-5 h-5 bg-green-100 dark:bg-green-900/30 rounded-full flex items-center justify-center mr-3"
								>
									<span class="text-xs text-green-500"
									>I</span>
								</div>
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>帮助改进应用</span>
							</div>
							<button
								@click="
									privacySettings
									.dataSharing
									.improveApp =
										!privacySettings
											.dataSharing
											.improveApp
								"
								class="relative inline-flex items-center h-5 rounded-full w-9 bg-gray-200 dark:bg-gray-700 transition-colors"
								:class='
									privacySettings
										.dataSharing
										.improveApp
									? "bg-green-500"
									: "bg-gray-300"
								'
							>
								<span
									class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
									:class='
										privacySettings
											.dataSharing
											.improveApp
										? "translate-x-5"
										: "translate-x-1"
									'
								/>
							</button>
						</div>
					</div>
				</div>

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-start">
						<EyeOff class="w-6 h-6 text-blue-500 mt-1 mr-3" />
						<div class="flex-1">
							<h2
								class="text-lg font-semibold text-gray-900 dark:text-white"
							>
								屏蔽的用户
							</h2>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								被屏蔽的用户无法联系您或查看您的在线状态
							</p>
						</div>
					</div>

					<div class="mt-6">
						<div class="flex mb-4">
							<input
								v-model="newBlockedUser"
								type="text"
								placeholder="输入用户名或ID"
								class="flex-1 p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-l-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							/>
							<button
								@click="addBlockedUser"
								class="px-4 py-3 bg-blue-500 hover:bg-blue-600 text-white rounded-r-lg transition-colors"
							>
								添加
							</button>
						</div>

						<div
							v-if="
								privacySettings
								.blockedUsers.length >
								0
							"
							class="space-y-3"
						>
							<div
								v-for="
									user in privacySettings
									.blockedUsers
								"
								:key="user.id"
								class="p-3 bg-gray-50 dark:bg-gray-700 rounded-lg flex items-center justify-between"
							>
								<div class="flex items-center">
									<div
										class="w-10 h-10 bg-gray-200 dark:bg-gray-600 rounded-full flex items-center justify-center mr-3"
									>
										<span
											class="text-gray-600 dark:text-gray-300 font-medium"
										>{{
											user
											.name
											.charAt(
												0,
											)
										}}</span>
									</div>
									<div>
										<div
											class="font-medium text-gray-900 dark:text-white"
										>
											{{
												user
												.name
											}}
										</div>
										<div
											class="text-sm text-gray-500 dark:text-gray-400"
										>
											@{{
												user
												.username
											}}
										</div>
									</div>
								</div>
								<button
									@click="
										removeBlockedUser(
											user.id,
										)
									"
									class="px-3 py-1.5 text-sm text-red-500 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
								>
									取消屏蔽
								</button>
							</div>
						</div>
						<div
							v-else
							class="p-8 text-center bg-gray-50 dark:bg-gray-700 rounded-lg"
						>
							<p class="text-gray-500 dark:text-gray-400">
								没有屏蔽的用户
							</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
