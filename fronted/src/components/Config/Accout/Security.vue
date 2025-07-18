<script setup lang="ts">
import { ref } from "vue";
import { Check, Key, Lock, Shield, Smartphone } from "lucide-vue-next";

const securitySettings = ref({
	passwordSet: true,
	twoFactorEnabled: false,
	recoveryEmail: "zhangsan@example.com",
	recoveryPhone: "13800138000",
	sessionActive: true,
	devices: [
		{
			id: 1,
			name: "PC",
			os: "Linux",
			lastActive: "今天 10:30",
			location: "宇宙",
		},
		{
			id: 2,
			name: "PC",
			os: "Linux",
			lastActive: "昨天 14:20",
			location: "宇宙",
		},
	],
});

const passwordChange = ref({
	current: "",
	new: "",
	confirm: "",
});

const twoFactorMethods = ref([
	{ id: "authenticator", name: "身份验证器应用", icon: Key, enabled: true },
	{ id: "sms", name: "短信验证", icon: Smartphone, enabled: false },
]);

const changePassword = () => {
	console.log("更改密码");
};

const toggleTwoFactor = (method: string) => {
	const methodIndex = twoFactorMethods.value.findIndex((m) =>
		m.id === method
	);
	if (methodIndex !== -1) {
		twoFactorMethods.value[methodIndex].enabled = !twoFactorMethods
			.value[methodIndex].enabled;
	}
};

const removeDevice = (id: number) => {
	securitySettings.value.devices = securitySettings.value.devices.filter(
		(device) => device.id !== id,
	);
};
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6">
		<div class="max-w-3xl mx-auto">
			<div class="mb-8">
				<h1 class="text-2xl font-bold text-gray-900 dark:text-white">
					安全设置
				</h1>
				<p class="text-gray-500 dark:text-gray-400 mt-2">
					管理您的账户安全设置和登录设备
				</p>
			</div>

			<div
				class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6 mb-6"
			>
				<div class="flex items-start justify-between mb-4">
					<div class="flex items-start">
						<Lock class="w-6 h-6 text-blue-500 mt-1 mr-3" />
						<div>
							<h2
								class="text-lg font-semibold text-gray-900 dark:text-white"
							>
								登录密码
							</h2>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								定期更改密码以保护您的账户安全
							</p>
						</div>
					</div>
					<button
						v-if="securitySettings.passwordSet"
						@click="changePassword"
						class="px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded-lg transition-colors"
					>
						更改密码
					</button>
					<button
						v-else
						@click="changePassword"
						class="px-4 py-2 bg-red-500 hover:bg-red-600 text-white rounded-lg transition-colors"
					>
						设置密码
					</button>
				</div>

				<div
					v-if="passwordChange.current"
					class="mt-4 pt-4 border-t border-gray-200 dark:border-gray-700"
				>
					<div class="space-y-4">
						<div>
							<label
								class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
							>当前密码</label>
							<input
								v-model="
									passwordChange
									.current
								"
								type="password"
								class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							/>
						</div>

						<div>
							<label
								class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
							>新密码</label>
							<input
								v-model="
									passwordChange
									.new
								"
								type="password"
								class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							/>
						</div>

						<div>
							<label
								class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
							>确认新密码</label>
							<input
								v-model="
									passwordChange
									.confirm
								"
								type="password"
								class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							/>
						</div>

						<div class="flex justify-end">
							<button
								class="px-4 py-2 bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 rounded-lg mr-2"
							>
								取消
							</button>
							<button
								class="px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded-lg flex items-center space-x-1"
							>
								<Check class="w-4 h-4" />
								<span>保存更改</span>
							</button>
						</div>
					</div>
				</div>
			</div>

			<div
				class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6 mb-6"
			>
				<div class="flex items-start">
					<Shield class="w-6 h-6 text-blue-500 mt-1 mr-3" />
					<div class="flex-1">
						<div class="flex justify-between items-start">
							<div>
								<h2
									class="text-lg font-semibold text-gray-900 dark:text-white"
								>
									双重验证
								</h2>
								<p
									class="text-sm text-gray-500 dark:text-gray-400 mt-1"
								>
									为您的账户添加额外的安全层
								</p>
							</div>
							<div class="flex items-center">
								<span
									class="mr-2 text-sm"
									:class='
										securitySettings
											.twoFactorEnabled
										? "text-green-500"
										: "text-gray-500"
									'
								>
									{{
										securitySettings
											.twoFactorEnabled
										? "已启用"
										: "已禁用"
									}}
								</span>
								<button
									@click="
										securitySettings
										.twoFactorEnabled =
											!securitySettings
												.twoFactorEnabled
									"
									class="relative inline-flex items-center h-6 rounded-full w-11 bg-gray-200 dark:bg-gray-700 transition-colors"
									:class='
										securitySettings
											.twoFactorEnabled
										? "bg-green-500"
										: "bg-gray-300"
									'
								>
									<span
										class="inline-block w-4 h-4 transform bg-white rounded-full transition-transform"
										:class='
											securitySettings
												.twoFactorEnabled
											? "translate-x-6"
											: "translate-x-1"
										'
									/>
								</button>
							</div>
						</div>
						<div
							v-if="
								securitySettings
								.twoFactorEnabled
							"
							class="mt-6 pt-4 border-t border-gray-200 dark:border-gray-700"
						>
							<h3
								class="font-medium text-gray-900 dark:text-white mb-4"
							>
								验证方法
							</h3>
							<div class="space-y-3">
								<div
									v-for="method in twoFactorMethods"
									:key="
										method
										.id
									"
									class="p-4 bg-gray-50 dark:bg-gray-700 rounded-lg flex items-center justify-between"
								>
									<div class="flex items-center">
										<component
											:is="
												method
												.icon
											"
											class="w-5 h-5 text-blue-500 mr-3"
										/>
										<span
											class="font-medium text-gray-700 dark:text-gray-300"
										>{{
											method
											.name
										}}</span>
									</div>
									<button
										@click="
											toggleTwoFactor(
												method
													.id,
											)
										"
										class="relative inline-flex items-center h-5 rounded-full w-9 bg-gray-200 dark:bg-gray-600 transition-colors"
										:class='
											method
												.enabled
											? "bg-green-500"
											: "bg-gray-300"
										'
									>
										<span
											class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
											:class='
												method
													.enabled
												? "translate-x-5"
												: "translate-x-1"
											'
										/>
									</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>

			<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
				<div class="flex items-start">
					<Smartphone class="w-6 h-6 text-blue-500 mt-1 mr-3" />
					<div class="flex-1">
						<h2
							class="text-lg font-semibold text-gray-900 dark:text-white"
						>
							登录设备
						</h2>
						<p
							class="text-sm text-gray-500 dark:text-gray-400 mt-1"
						>
							管理已登录此账户的设备
						</p>
					</div>
				</div>

				<div class="mt-6 space-y-4">
					<div
						v-for="device in securitySettings.devices"
						:key="device.id"
						class="p-4 bg-gray-50 dark:bg-gray-700 rounded-lg flex items-center"
					>
						<div class="mr-4">
							<div
								class="w-10 h-10 bg-blue-100 dark:bg-gray-600 rounded-lg flex items-center justify-center"
							>
								<Smartphone class="w-5 h-5 text-blue-500" />
							</div>
						</div>

						<div class="flex-1 min-w-0">
							<div class="flex items-center justify-between">
								<h3
									class="font-medium text-gray-900 dark:text-white"
								>
									{{
										device
										.name
									}}
								</h3>
								<span
									v-if="
										device
										.id ===
										1
									"
									class="px-2 py-1 bg-green-100 dark:bg-green-900/30 text-green-800 dark:text-green-400 text-xs rounded-full"
								>
									当前设备
								</span>
							</div>
							<div
								class="flex flex-wrap text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								<span class="mr-3">{{
									device
									.os
								}}</span>
								<span class="mr-3">最后活跃: {{
										device
										.lastActive
									}}</span>
								<span>{{
									device
									.location
								}}</span>
							</div>
						</div>

						<button
							v-if="device.id !== 1"
							@click="removeDevice(device.id)"
							class="ml-4 px-3 py-1.5 text-sm text-red-500 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
						>
							移除
						</button>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
