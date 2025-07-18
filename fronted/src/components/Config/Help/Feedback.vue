<script setup lang="ts">
import { ref } from "vue";
import {
	AlertCircle,
	AlertTriangle,
	Check,
	FileText,
	Mail,
	Send,
} from "lucide-vue-next";

const feedbackType = ref("suggestion");
const feedbackContent = ref("");
const contactInfo = ref("");
const attachments = ref<File[]>([]);
const isSubmitting = ref(false);
const submitSuccess = ref(false);
const submitError = ref(false);
const appName = import.meta.env.VITE_APP_NAME;

const feedbackTypes = [
	{ id: "suggestion", label: "功能建议", icon: FileText },
	{ id: "bug", label: "错误报告", icon: AlertTriangle },
	{ id: "other", label: "其他反馈", icon: Mail },
];

const handleFileUpload = (event: Event) => {
	const input = event.target as HTMLInputElement;
	if (input.files && input.files.length > 0) {
		for (let i = 0; i < input.files.length; i++) {
			attachments.value.push(input.files[i]);
		}
	}
};

const removeAttachment = (index: number) => {
	attachments.value.splice(index, 1);
};

const submitFeedback = () => {
	if (!feedbackContent.value.trim()) {
		submitError.value = true;
		return;
	}

	isSubmitting.value = true;

	setTimeout(() => {
		isSubmitting.value = false;
		submitSuccess.value = true;
		feedbackContent.value = "";
		contactInfo.value = "";
		attachments.value = [];

		setTimeout(() => {
			submitSuccess.value = false;
		}, 5000);
	}, 1500);
};
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6">
		<div class="max-w-3xl mx-auto">
			<div class="flex items-center mb-8">
				<div
					class="p-3 bg-indigo-100 dark:bg-indigo-900/30 rounded-lg mr-4"
				>
					<Send class="w-6 h-6 text-indigo-500" />
				</div>
				<div>
					<h1
						class="text-2xl font-bold text-gray-900 dark:text-white"
					>
						意见反馈
					</h1>
					<p class="text-gray-500 dark:text-gray-400 mt-1">
						您的反馈对我们非常重要，帮助我们改进{{ appName }}
					</p>
				</div>
			</div>

			<div
				v-if="submitSuccess"
				class="mb-6 p-4 bg-green-100 dark:bg-green-900/30 text-green-800 dark:text-green-300 rounded-lg flex items-center"
			>
				<Check class="w-5 h-5 mr-2" />
				<span>反馈提交成功！感谢您的宝贵意见</span>
			</div>

			<div
				v-if="submitError"
				class="mb-6 p-4 bg-red-100 dark:bg-red-900/30 text-red-800 dark:text-red-300 rounded-lg flex items-center"
			>
				<AlertCircle class="w-5 h-5 mr-2" />
				<span>请填写反馈内容后再提交</span>
			</div>

			<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
				<div class="space-y-6">
					<div>
						<h2
							class="text-lg font-medium text-gray-900 dark:text-white mb-3"
						>
							反馈类型
						</h2>
						<div class="grid grid-cols-1 md:grid-cols-3 gap-3">
							<button
								v-for="type in feedbackTypes"
								:key="type.id"
								@click="
									feedbackType =
									type.id
								"
								class="p-4 border rounded-lg transition-all duration-200"
								:class='
									feedbackType ===
										type.id
									? "border-blue-500 bg-blue-50 dark:bg-blue-900/20"
									: "border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700"
								'
							>
								<span class="flex flex-col items-center">
									<component
										:is="
											type
											.icon
										"
										class="w-6 h-6 mb-2"
										:class='
											feedbackType ===
												type.id
											? "text-blue-500"
											: "text-gray-700 dark:text-gray-300"
										'
									/>
									<span
										class="font-medium"
										:class='
											feedbackType ===
												type.id
											? "text-blue-500"
											: "text-gray-700 dark:text-gray-300"
										'
									>
										{{
											type
											.label
										}}
									</span>
								</span>
							</button>
						</div>
					</div>

					<div>
						<h2
							class="text-lg font-medium text-gray-900 dark:text-white mb-3"
						>
							反馈内容
						</h2>
						<textarea
							v-model="feedbackContent"
							placeholder="请详细描述您的反馈、建议或遇到的问题..."
							class="w-full h-40 p-4 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
						></textarea>
					</div>

					<div>
						<h2
							class="text-lg font-medium text-gray-900 dark:text-white mb-3"
						>
							联系方式（可选）
						</h2>
						<p
							class="text-sm text-gray-500 dark:text-gray-400 mb-2"
						>
							如果需要我们回复您的反馈，请留下您的联系方式
						</p>
						<input
							v-model="contactInfo"
							type="text"
							placeholder="邮箱或手机号码"
							class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
						/>
					</div>

					<div>
						<h2
							class="text-lg font-medium text-gray-900 dark:text-white mb-3"
						>
							添加附件
						</h2>
						<div class="flex items-center space-x-3">
							<label class="cursor-pointer">
								<input
									type="file"
									class="hidden"
									multiple
									@change="handleFileUpload"
								/>
								<span
									class="px-4 py-2 bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 rounded-lg transition-colors"
								>
									选择文件
								</span>
							</label>
							<span
								class="text-sm text-gray-500 dark:text-gray-400"
							>
								最多可上传5个文件（图片、日志等）
							</span>
						</div>

						<div
							v-if="attachments.length > 0"
							class="mt-4 space-y-2"
						>
							<div
								v-for="(file, index) in attachments"
								:key="index"
								class="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-700 rounded-lg"
							>
								<div class="flex items-center">
									<FileText
										class="w-5 h-5 text-gray-500 mr-2"
									/>
									<span class="text-sm truncate max-w-xs">{{
										file
										.name
									}}</span>
								</div>
								<button
									@click="
										removeAttachment(
											index,
										)
									"
									class="text-red-500 hover:text-red-700 transition-colors"
								>
									<X class="w-5 h-5" />
								</button>
							</div>
						</div>
					</div>

					<div class="pt-4">
						<button
							@click="submitFeedback"
							:disabled="isSubmitting"
							class="w-full py-3 bg-blue-500 hover:bg-blue-600 text-white rounded-lg transition-colors flex items-center justify-center"
							:class='
								{
									"opacity-70 cursor-not-allowed":
										isSubmitting,
								}
							'
						>
							<template v-if="isSubmitting">
								<svg
									class="animate-spin -ml-1 mr-2 h-5 w-5 text-white"
									xmlns="http://www.w3.org/2000/svg"
									fill="none"
									viewBox="0 0 24 24"
								>
									<circle
										class="opacity-25"
										cx="12"
										cy="12"
										r="10"
										stroke="currentColor"
										stroke-width="4"
									>
									</circle>
									<path
										class="opacity-75"
										fill="currentColor"
										d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
									>
									</path>
								</svg>
								提交中...
							</template>
							<template v-else>
								<Send class="w-5 h-5 mr-2" />
								提交反馈
							</template>
						</button>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
