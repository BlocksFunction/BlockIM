<script setup lang="ts">
import { ref } from "vue";
import { Camera, Check, Edit, User } from "lucide-vue-next";

const userProfile = ref({
	name: "ImAur",
	username: "ImAur",
	bio: "IDK",
	email: "IDK",
	phone: "IDK",
	gender: "IDK",
	birthday: "IDK",
	region: "IDK",
});

const editMode = ref(false);
const tempProfile = ref({ ...userProfile.value });

const saveProfile = () => {
	userProfile.value = { ...tempProfile.value };
	editMode.value = false;
};

const cancelEdit = () => {
	tempProfile.value = { ...userProfile.value };
	editMode.value = false;
};
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6 text-gray-900 dark:text-white">
		<div class="max-w-3xl mx-auto">
			<div class="flex justify-between items-center mb-8">
				<h1 class="text-2xl font-bold">个人资料</h1>
				<div v-if="!editMode">
					<button
						@click="editMode = true"
						class="flex items-center space-x-1 px-4 py-2 bg-white dark:bg-gray-800 border border-gray-300 dark:border-gray-600 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
					>
						<Edit class="w-4 h-4" />
						<span>编辑资料</span>
					</button>
				</div>
				<div v-else class="flex space-x-2">
					<button
						@click="cancelEdit"
						class="px-4 py-2 bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600 rounded-lg transition-colors"
					>
						取消
					</button>
					<button
						@click="saveProfile"
						class="px-4 py-2 bg-blue-500 hover:bg-blue-600 text-white rounded-lg transition-colors flex items-center space-x-1"
					>
						<Check class="w-4 h-4" />
						<span>保存更改</span>
					</button>
				</div>
			</div>

			<div class="flex flex-col items-center mb-8">
				<div class="relative mb-4">
					<div
						class="w-32 h-32 bg-gray-200 dark:bg-gray-700 rounded-full overflow-hidden"
					>
						<div
							class="w-full h-full flex items-center justify-center text-gray-400"
						>
							<User class="w-16 h-16" />
						</div>
					</div>
					<button
						v-if="editMode"
						class="absolute bottom-0 right-0 p-2 bg-white dark:bg-gray-800 rounded-full border border-gray-300 dark:border-gray-600 shadow-sm hover:bg-gray-50 dark:hover:bg-gray-700"
					>
						<Camera
							class="w-5 h-5 text-gray-600 dark:text-gray-300"
						/>
					</button>
				</div>
				<h2 class="text-xl font-semibold text-gray-900 dark:text-white">
					{{ userProfile.name }}
				</h2>
				<p class="text-gray-500 dark:text-gray-400">
					@{{ userProfile.username }}
				</p>
			</div>

			<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
				<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
					<div>
						<label
							class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
						>姓名</label>
						<input
							v-model="tempProfile.name"
							:readonly="!editMode"
							type="text"
							class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							:class='
								editMode
								? "cursor-text"
								: "cursor-default"
							'
						/>
					</div>

					<div>
						<label
							class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
						>用户名</label>
						<input
							v-model="tempProfile.username"
							:readonly="!editMode"
							type="text"
							class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							:class='
								editMode
								? "cursor-text"
								: "cursor-default"
							'
						/>
					</div>

					<div class="md:col-span-2">
						<label
							class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
						>个人简介</label>
						<textarea
							v-model="tempProfile.bio"
							:readonly="!editMode"
							rows="3"
							class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							:class='
								editMode
								? "cursor-text"
								: "cursor-default"
							'
						></textarea>
					</div>

					<div>
						<label
							class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
						>电子邮件</label>
						<input
							v-model="tempProfile.email"
							:readonly="!editMode"
							type="email"
							class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							:class='
								editMode
								? "cursor-text"
								: "cursor-default"
							'
						/>
					</div>

					<div>
						<label
							class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
						>手机号码</label>
						<input
							v-model="tempProfile.phone"
							:readonly="!editMode"
							type="tel"
							class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							:class='
								editMode
								? "cursor-text"
								: "cursor-default"
							'
						/>
					</div>

					<div>
						<label
							class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
						>性别</label>
						<select
							v-model="tempProfile.gender"
							:disabled="!editMode"
							class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							:class='
								editMode
								? "cursor-pointer"
								: "cursor-default"
							'
						>
							<option value="男">男</option>
							<option value="女">女</option>
							<option value="其他">其他</option>
						</select>
					</div>

					<div>
						<label
							class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
						>出生日期</label>
						<input
							v-model="tempProfile.birthday"
							:readonly="!editMode"
							type="date"
							class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							:class='
								editMode
								? "cursor-text"
								: "cursor-default"
							'
						/>
					</div>

					<div>
						<label
							class="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-1"
						>地区</label>
						<input
							v-model="tempProfile.region"
							:readonly="!editMode"
							type="text"
							class="w-full p-3 bg-gray-50 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
							:class='
								editMode
								? "cursor-text"
								: "cursor-default"
							'
						/>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
