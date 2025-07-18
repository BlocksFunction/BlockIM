<script setup lang="ts">
import { ref } from "vue";
import { Bell, BellOff, Mail, Vibrate } from "lucide-vue-next";

const settings = ref({
	enabled: true,
	sound: "default",
	vibrate: true,
	preview: "full",
	silentHours: {
		enabled: false,
		start: "22:00",
		end: "07:00",
	},
	customSounds: [
		{ id: "default", name: "默认提示音" },
		{ id: "chime", name: "清脆铃声" },
		{ id: "ding", name: "叮咚声" },
		{ id: "bubble", name: "气泡声" },
	],
});

const toggleSetting = (setting: keyof typeof settings.value) => {
	if (setting === "enabled" || setting === "vibrate") {
		settings.value[setting] = !settings.value[setting];
	}
};

const toggleSilentHours = () => {
	settings.value.silentHours.enabled = !settings.value.silentHours.enabled;
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
							消息通知
						</h1>
						<p class="text-gray-500 dark:text-gray-400 mt-1">
							管理您接收新消息的方式
						</p>
					</div>
				</div>
			</div>

			<div class="space-y-6">
				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-center justify-between">
						<div class="flex items-center">
							<Bell class="w-6 h-6 mr-3 text-blue-500" />
							<div>
								<h2
									class="text-lg font-semibold text-gray-900 dark:text-white"
								>
									消息通知
								</h2>
								<p
									class="text-sm text-gray-500 dark:text-gray-400 mt-1"
								>
									接收新消息通知提醒
								</p>
							</div>
						</div>
						<button
							@click='toggleSetting("enabled")'
							class="relative inline-flex items-center h-6 rounded-full w-11 transition-colors"
							:class='
								settings.enabled
								? "bg-green-500"
								: "bg-gray-300 dark:bg-gray-600"
							'
						>
							<span
								class="inline-block w-4 h-4 transform bg-white rounded-full transition-transform"
								:class='
									settings.enabled
									? "translate-x-6"
									: "translate-x-1"
								'
							/>
						</button>
					</div>
				</div>

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-start">
						<Mail class="w-6 h-6 mt-1 mr-3 text-blue-500" />
						<div class="flex-1">
							<h2
								class="text-lg font-semibold text-gray-900 dark:text-white"
							>
								通知声音
							</h2>
							<p
								class="text-sm text-gray-500 dark:text-gray-400 mt-1"
							>
								选择新消息的通知提示音
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
										settings
											.sound ===
											sound
												.id
										? "border-blue-500 bg-blue-50 dark:bg-blue-900/20"
										: "border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700"
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
										接收新消息时设备震动
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
								<Bell class="w-6 h-6 mr-3 text-blue-500" />
								<div>
									<h2
										class="font-medium text-gray-900 dark:text-white"
									>
										消息预览
									</h2>
									<p
										class="text-sm text-gray-500 dark:text-gray-400 mt-1"
									>
										在通知中显示消息内容
									</p>
								</div>
							</div>
							<select
								v-model="settings.preview"
								class="px-3 py-1.5 bg-gray-50 focus:border-blue-500 focus:ring-blue-500 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2"
							>
								<option value="full">显示完整内容</option>
								<option value="partial">仅显示发送者</option>
								<option value="none">不显示内容</option>
							</select>
						</div>
					</div>
				</div>

				<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
					<div class="flex items-start">
						<BellOff class="w-6 h-6 mt-1 mr-3 text-blue-500" />
						<div class="flex-1">
							<div class="flex justify-between items-start">
								<div>
									<h2
										class="text-lg font-semibold text-gray-900 dark:text-white"
									>
										静音时段
									</h2>
									<p
										class="text-sm text-gray-500 dark:text-gray-400 mt-1"
									>
										在指定时间段内静音通知
									</p>
								</div>
								<button
									@click="toggleSilentHours"
									class="relative inline-flex items-center h-5 rounded-full w-9 transition-colors"
									:class='
										settings
											.silentHours
											.enabled
										? "bg-green-500"
										: "bg-gray-300 dark:bg-gray-600"
									'
								>
									<span
										class="inline-block w-3 h-3 transform bg-white rounded-full transition-transform"
										:class='
											settings
												.silentHours
												.enabled
											? "translate-x-5"
											: "translate-x-1"
										'
									/>
								</button>
							</div>

							<div
								v-if="
									settings
									.silentHours
									.enabled
								"
								class="mt-4 pt-4 border-t border-gray-200 dark:border-gray-700"
							>
								<div class="grid grid-cols-2 gap-4">
									<div>
										<label
											class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
										>开始时间</label>
										<input
											v-model="
												settings
												.silentHours
												.start
											"
											type="time"
											class="w-full p-2.5 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2"
										/>
									</div>
									<div>
										<label
											class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
										>结束时间</label>
										<input
											v-model="
												settings
												.silentHours
												.end
											"
											type="time"
											class="w-full p-2.5 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2"
										/>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
