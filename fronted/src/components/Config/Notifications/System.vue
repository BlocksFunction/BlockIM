<script setup lang="ts">
import { ref } from "vue";
import {
	Bell,
	Info,
	RefreshCw,
	Shield,
	UserCheck,
	Vibrate,
} from "lucide-vue-next";

const settings = ref({
	updatesEnabled: true,
	securityEnabled: true,
	accountEnabled: true,
	sound: "chime",
	vibrate: false,
	customSounds: [
		{ id: "chime", name: "提示音" },
		{ id: "ding", name: "叮咚声" },
		{ id: "beep", name: "蜂鸣声" },
	],
});

const toggleSetting = (setting: keyof typeof settings.value) => {
	if (
		setting === "updatesEnabled" || setting === "securityEnabled" ||
		setting === "accountEnabled" || setting === "vibrate"
	) {
		settings.value[setting] = !settings.value[setting];
	}
};

const appName = import.meta.env.VITE_APP_NAME;
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6">
		<div class="max-w-3xl mx-auto">
			<div class="mb-8">
				<div class="flex items-center mb-4">
					<div>
						<h1
							class="text-2xl font-bold text-gray-900 dark:text-white"
						>
							系统通知
						</h1>
						<p class="text-gray-500 dark:text-gray-400 mt-1">
							管理应用系统通知和更新提醒
						</p>
					</div>
				</div>
			</div>

			<div class="space-y-6">
				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-start">
						<Bell class="w-6 h-6 mt-1 mr-3 text-blue-500" />
						<div class="flex-1">
							<h2
								class="text-lg font-semibold text-gray-900 dark:text-white"
							>
								通知类型
							</h2>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								选择接收的系统通知类型
							</p>

							<div class="mt-4 space-y-4">
								<div class="flex items-center justify-between">
									<div class="flex items-center">
										<RefreshCw
											class="w-5 h-5 mr-3 text-blue-500"
										/>
										<div>
											<h2
												class="font-medium text-gray-900 dark:text-white"
											>
												应用更新
											</h2>
											<p
												class="text-sm text-gray-500 dark:text-gray-400"
											>
												新版本可用时通知
											</p>
										</div>
									</div>
									<button
										@click='
											toggleSetting(
												"updatesEnabled",
											)
										'
										class="relative inline-flex items-center h-5 rounded-full w-9 transition-colors"
										:class='
											settings
												.updatesEnabled
											? "bg-green-500"
											: "bg-gray-300 dark:bg-gray-600"
										'
									>
										<span
											class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
											:class='
												settings
													.updatesEnabled
												? "translate-x-5"
												: "translate-x-1"
											'
										/>
									</button>
								</div>

								<div class="flex items-center justify-between">
									<div class="flex items-center">
										<Shield
											class="w-5 h-5 mr-3 text-blue-500"
										/>
										<div>
											<h2
												class="font-medium text-gray-900 dark:text-white"
											>
												安全通知
											</h2>
											<p
												class="text-sm text-gray-500 dark:text-gray-400"
											>
												账户安全相关通知
											</p>
										</div>
									</div>
									<button
										@click='
											toggleSetting(
												"securityEnabled",
											)
										'
										class="relative inline-flex items-center h-5 rounded-full w-9 transition-colors"
										:class='
											settings
												.securityEnabled
											? "bg-green-500"
											: "bg-gray-300 dark:bg-gray-600"
										'
									>
										<span
											class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
											:class='
												settings
													.securityEnabled
												? "translate-x-5"
												: "translate-x-1"
											'
										/>
									</button>
								</div>

								<div class="flex items-center justify-between">
									<div class="flex items-center">
										<UserCheck
											class="w-5 h-5 mr-3 text-blue-500"
										/>
										<div>
											<h2
												class="font-medium text-gray-900 dark:text-white"
											>
												账户活动
											</h2>
											<p
												class="text-sm text-gray-500 dark:text-gray-400"
											>
												登录活动和其他账户通知
											</p>
										</div>
									</div>
									<button
										@click='
											toggleSetting(
												"accountEnabled",
											)
										'
										class="relative inline-flex items-center h-5 rounded-full w-9 transition-colors"
										:class='
											settings
												.accountEnabled
											? "bg-green-500"
											: "bg-gray-300 dark:bg-gray-600"
										'
									>
										<span
											class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
											:class='
												settings
													.accountEnabled
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

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-start">
						<Bell class="w-6 h-6 mt-1 mr-3 text-blue-500" />
						<div class="flex-1">
							<h2
								class="text-lg font-semibold text-gray-900 dark:text-white"
							>
								通知声音
							</h2>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								选择系统通知的提示音
							</p>

							<div class="mt-4 grid grid-cols-2 gap-3">
								<button
									v-for="
										sound in settings
										.customSounds
									"
									:key="sound.id"
									class="p-3 rounded-lg border transition-colors"
									:class='
										[
											settings
													.sound ===
													sound
														.id
												? "border-blue-500 bg-blue-50 dark:bg-blue-900/20"
												: "border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700",
										]
									'
									@click="
										settings
										.sound =
											sound
												.id
									"
								>
									<span class="flex items-center">
										<span
											class="w-8 h-8 bg-blue-100 dark:bg-blue-900/30 rounded-full flex items-center justify-center mr-3"
										>
											<Bell
												class="w-4 h-4 text-blue-500"
											/>
										</span>
										<span
											class="text-sm font-medium text-gray-700 dark:text-gray-300"
										>{{
											sound
											.name
										}}</span>
									</span>
								</button>
							</div>
						</div>
					</div>
				</div>

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="space-y-4">
						<div class="flex items-center justify-between">
							<div class="flex items-center">
								<Vibrate class="w-6 h-6 mr-3 text-blue-500" />
								<div>
									<h2
										class="font-medium text-gray-900 dark:text-white"
									>
										震动提醒
									</h2>
									<p
										class="text-sm text-gray-500 dark:text-gray-400 mt-1"
									>
										系统通知时设备震动
									</p>
								</div>
							</div>
							<button
								@click='
									toggleSetting(
										"vibrate",
									)
								'
								class="relative inline-flex items-center h-5 rounded-full w-9 transition-colors"
								:class='
									settings.vibrate
									? "bg-green-500"
									: "bg-gray-300 dark:bg-gray-600"
								'
							>
								<span
									class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
									:class='
										settings
											.vibrate
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
</template>
