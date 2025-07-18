<script setup lang="ts">
import { ref } from "vue";
import {
	Bell,
	BellOff,
	Phone,
	PhoneCall,
	PhoneOff,
	Vibrate,
} from "lucide-vue-next";

const settings = ref({
	incomingEnabled: true,
	missedEnabled: true,
	sound: "ringtone",
	vibrate: true,
	silentMode: false,
	customRingtones: [
		{ id: "ringtone", name: "默认铃声" },
		{ id: "classic", name: "经典铃声" },
		{ id: "digital", name: "数字铃声" },
		{ id: "melody", name: "旋律铃声" },
	],
});

const toggleSetting = (setting: keyof typeof settings.value) => {
	if (
		setting === "incomingEnabled" || setting === "missedEnabled" ||
		setting === "vibrate" || setting === "silentMode"
	) {
		settings.value[setting] = !settings.value[setting];
	}
};
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
							通话通知
						</h1>
						<p class="text-gray-500 dark:text-gray-400 mt-1">
							管理您接收通话通知的方式
						</p>
					</div>
				</div>
			</div>

			<div class="space-y-6">
				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-center justify-between">
						<div class="flex items-center">
							<Phone class="w-6 h-6 mr-3 text-blue-500" />
							<div>
								<h2
									class="text-lg font-semibold text-gray-900 dark:text-white"
								>
									来电通知
								</h2>
								<p
									class="text-sm text-gray-500 dark:text-gray-400 mt-1"
								>
									接收来电通知提醒
								</p>
							</div>
						</div>
						<button
							@click='
								toggleSetting(
									"incomingEnabled",
								)
							'
							class="relative inline-flex items-center h-6 rounded-full w-11 transition-colors"
							:class='
								settings.incomingEnabled
								? "bg-green-500"
								: "bg-gray-300 dark:bg-gray-600"
							'
						>
							<span
								class="inline-block w-4 h-4 transform bg-white rounded-full transition-transform"
								:class='
									settings
										.incomingEnabled
									? "translate-x-6"
									: "translate-x-1"
								'
							/>
						</button>
					</div>
				</div>

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-center justify-between">
						<div class="flex items-center">
							<PhoneOff class="w-6 h-6 mr-3 text-blue-500" />
							<div>
								<h2
									class="text-lg font-semibold text-gray-900 dark:text-white"
								>
									未接来电通知
								</h2>
								<p
									class="text-sm text-gray-500 dark:text-gray-400 mt-1"
								>
									接收未接来电通知提醒
								</p>
							</div>
						</div>
						<button
							@click='
								toggleSetting(
									"missedEnabled",
								)
							'
							class="relative inline-flex items-center h-6 rounded-full w-11 transition-colors"
							:class='
								settings.missedEnabled
								? "bg-green-500"
								: "bg-gray-300 dark:bg-gray-600"
							'
						>
							<span
								class="inline-block w-4 h-4 transform bg-white rounded-full transition-transform"
								:class='
									settings
										.missedEnabled
									? "translate-x-6"
									: "translate-x-1"
								'
							/>
						</button>
					</div>
				</div>

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-start">
						<Bell class="w-6 h-6 mt-1 mr-3 text-blue-500" />
						<div class="flex-1">
							<h2
								class="text-lg font-semibold text-gray-900 dark:text-white"
							>
								铃声设置
							</h2>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								选择来电铃声
							</p>

							<div class="mt-4 grid grid-cols-2 gap-3">
								<button
									v-for="
										ringtone in settings
										.customRingtones
									"
									:key="
										ringtone
										.id
									"
									class="p-3 rounded-lg border transition-colors"
									:class='
										settings
											.sound ===
											ringtone
												.id
										? "border-blue-500 bg-blue-50 dark:bg-blue-900/20"
										: "border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700"
									'
									@click="
										settings
										.sound =
											ringtone
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
											ringtone
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
										来电时设备震动
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

						<div class="flex items-center justify-between">
							<div class="flex items-center">
								<BellOff class="w-6 h-6 mr-3 text-blue-500" />
								<div>
									<h2
										class="font-medium text-gray-900 dark:text-white"
									>
										静音模式
									</h2>
									<p
										class="text-sm text-gray-500 dark:text-gray-400 mt-1"
									>
										来电时不发出声音
									</p>
								</div>
							</div>
							<button
								@click='
									toggleSetting(
										"silentMode",
									)
								'
								class="relative inline-flex items-center h-5 rounded-full w-9 transition-colors"
								:class='
									settings
										.silentMode
									? "bg-green-500"
									: "bg-gray-300 dark:bg-gray-600"
								'
							>
								<span
									class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
									:class='
										settings
											.silentMode
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
